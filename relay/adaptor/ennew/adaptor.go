package ennew

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/songquanpeng/one-api/relay/adaptor"
	"github.com/songquanpeng/one-api/relay/meta"
	"github.com/songquanpeng/one-api/relay/model"
	"github.com/songquanpeng/one-api/relay/relaymode"
	"io"
	"net/http"
	"fmt"
	"time"
	"encoding/json"
	"net/url"
)

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
	formData := url.Values{}
	formData.Set("appKey", appKey)
	formData.Set("appSecret", appSecret)

	// 配置请求URL: https://middle-open-platform.ennew.com/admin/client/getToken
    resp, err := http.PostForm("https://middle-open-platform.ennew.com/admin/client/getToken", formData)
    if err != nil {
        return "", fmt.Errorf("failed to get API key: %v", err)
    }
    defer resp.Body.Close()

    var apiKeyResponse struct {
        Code    int         `json:"code"`
        Data    string      `json:"data"`
    }
    if err := json.NewDecoder(resp.Body).Decode(&apiKeyResponse); err != nil {
        return "", fmt.Errorf("failed to parse API key response: %v", err)
    }

    if apiKeyResponse.Code != 200 {
        return "", fmt.Errorf("failed to get API key, got non-200 status code: %d", apiKeyResponse.Code)
    }

    // 存储API Key和过期时间（过期时间为24小时后）
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
