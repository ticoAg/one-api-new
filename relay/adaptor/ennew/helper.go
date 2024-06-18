package ennew

import (
	"fmt"
	"github.com/songquanpeng/one-api/common/logger"
	"github.com/songquanpeng/one-api/relay/model"
	"strconv"
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

	return fullRequestURL
}
