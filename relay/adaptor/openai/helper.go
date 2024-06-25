package openai

import (
	"fmt"
	"github.com/ticoAg/one-api-new/common/logger"
	"github.com/ticoAg/one-api-new/relay/channeltype"
	"github.com/ticoAg/one-api-new/relay/model"
	"strconv"
	"strings"
)

func ResponseText2Usage(responseText string, modeName string, promptTokens int) *model.Usage {
	usage := &model.Usage{}
	usage.PromptTokens = promptTokens
	usage.CompletionTokens = CountTokenText(responseText, modeName)
	usage.TotalTokens = usage.PromptTokens + usage.CompletionTokens
	return usage
}

func GetFullRequestURL(baseURL string, requestURL string, channelType int) string {
	logger.SysLog("baseURL:" + baseURL + ", requestURL:" + requestURL + ", channelType:" + strconv.Itoa(channelType))
	fullRequestURL := fmt.Sprintf("%s%s", baseURL, requestURL)
	logger.SysLog("fullRequestURL:" + fullRequestURL)

	if strings.HasPrefix(baseURL, "https://gateway.ai.cloudflare.com") {
		switch channelType {
		case channeltype.OpenAI:
			fullRequestURL = fmt.Sprintf("%s%s", baseURL, strings.TrimPrefix(requestURL, "/v1"))
		case channeltype.Azure:
			fullRequestURL = fmt.Sprintf("%s%s", baseURL, strings.TrimPrefix(requestURL, "/openai/deployments"))
		}
	}
	return fullRequestURL
}
