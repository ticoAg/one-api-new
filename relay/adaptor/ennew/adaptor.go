package ennew

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/ticoAg/one-api-new/relay/adaptor"
	"github.com/ticoAg/one-api-new/relay/meta"
	"github.com/ticoAg/one-api-new/relay/model"
	"github.com/ticoAg/one-api-new/relay/relaymode"
	"io"
	// "bytes"
	"net/http"
	"fmt"
	"time"
	"encoding/json"
)
import "net/url"

type Adaptor struct {
	ChannelType int
}

func (a *Adaptor) Init(meta *meta.Meta) {
	a.ChannelType = meta.ChannelType
}

func (a *Adaptor) GetRequestURL(meta *meta.Meta) (string, error) {
	return GetFullRequestURL(meta.BaseURL, meta.RequestURLPath, meta.ChannelType), nil
}


var apiKeyCache = make(map[string]string)
var apiKeyExpireTime = make(map[string]time.Time)

func getAPIKey(appKey, appSecret string) (string, error) {
    if apiKey, ok := apiKeyCache[appKey]; ok && time.Now().Before(apiKeyExpireTime[appKey]) {
        return apiKey, nil
    }

	baseUrl := "https://middle-open-platform.ennew.com/admin/client/getToken"
	values := url.Values{}
	values.Add("appKey", appKey)
	values.Add("appSecret", appSecret)

	resp, err := http.Get(baseUrl + "?" + values.Encode())
	if err != nil {
		return "", fmt.Errorf("filed to get API key: %v", err)
	}
    defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read response body: %v", err)
	}
    fmt.Println("API Response Body:", string(bodyBytes))

    var apiKeyResponse struct {
        Code    int         `json:"code"`
        Data    string      `json:"data"`
    }
    // if err := json.NewDecoder(resp.Body).Decode(&apiKeyResponse); err != nil {
    //     return "", fmt.Errorf("failed to parse API key response: %v", err)
    // }

	if err := json.Unmarshal(bodyBytes, &apiKeyResponse); err != nil {
        return "", fmt.Errorf("failed to parse API key response: %v", err)
    }

    if apiKeyResponse.Code != 200 {
        return "", fmt.Errorf("failed to get API key, got non-200 status code: %d", apiKeyResponse.Code)
    }

    apiKeyCache[appKey] = apiKeyResponse.Data
    apiKeyExpireTime[appKey] = time.Now().Add(24 * time.Hour) // 过期时间为24小时后

    return apiKeyResponse.Data, nil
}
func (a *Adaptor) SetupRequestHeader(c *gin.Context, req *http.Request, meta *meta.Meta) error {
	apiKey, err := getAPIKey(meta.Config.AK, meta.Config.SK)
	
	if err != nil {
        return fmt.Errorf("failed to get API key: %v", err)
    }
	adaptor.SetupCommonRequestHeader(c, req, meta)
	req.Header.Set("X-GW-Authorization", apiKey)
	return nil
}

func (a *Adaptor) ConvertRequest(c *gin.Context, relayMode int, request *model.GeneralOpenAIRequest) (any, error) {
	if request == nil {
		return nil, errors.New("request is nil")
	}
	return request, nil
}

func (a *Adaptor) ConvertImageRequest(request *model.ImageRequest) (any, error) {
	if request == nil {
		return nil, errors.New("request is nil")
	}
	return request, nil
}

func (a *Adaptor) DoRequest(c *gin.Context, meta *meta.Meta, requestBody io.Reader) (*http.Response, error) {
	return adaptor.DoRequestHelper(a, c, meta, requestBody)
}

func (a *Adaptor) DoResponse(c *gin.Context, resp *http.Response, meta *meta.Meta) (usage *model.Usage, err *model.ErrorWithStatusCode) {
	if meta.IsStream {
		var responseText string
		err, responseText, usage = StreamHandler(c, resp, meta.Mode)
		if usage == nil || usage.TotalTokens == 0 {
			usage = ResponseText2Usage(responseText, meta.ActualModelName, meta.PromptTokens)
		}
		if usage.TotalTokens != 0 && usage.PromptTokens == 0 { // some channels don't return prompt tokens & completion tokens
			usage.PromptTokens = meta.PromptTokens
			usage.CompletionTokens = usage.TotalTokens - meta.PromptTokens
		}
	} else {
		switch meta.Mode {
		case relaymode.ImagesGenerations:
			err, _ = ImageHandler(c, resp)
		default:
			err, usage = Handler(c, resp, meta.PromptTokens, meta.ActualModelName)
		}
	}
	return
}

func (a *Adaptor) GetModelList() []string {
	_, modelList := GetCompatibleChannelMeta(a.ChannelType)
	return modelList
}

func (a *Adaptor) GetChannelName() string {
	channelName, _ := GetCompatibleChannelMeta(a.ChannelType)
	return channelName
}
