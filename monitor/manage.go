package monitor

import (
	"github.com/ticoAg/one-api-new/common/config"
	"github.com/ticoAg/one-api-new/relay/model"
	"net/http"
	"strings"
)

func ShouldDisableChannel(err *model.Error, statusCode int) bool {
	if !config.AutomaticDisableChannelEnabled {
		return false
	}
	if err == nil {
		return false
	}
	if statusCode == http.StatusUnauthorized {
		return true
	}
	switch err.Type {
	case "insufficient_quota":
		return true
	// https://docs.anthropic.com/claude/reference/errors
	case "authentication_error":
		return true
	case "permission_error":
		return true
	case "forbidden":
		return true
	}
	if err.Code == "invalid_api_key" || err.Code == "account_deactivated" {
		return true
	}
	if strings.HasPrefix(err.Message, "Your credit balance is too low") { // anthropic
		return true
	} else if strings.HasPrefix(err.Message, "This organization has been disabled.") {
		return true
	}
	//if strings.Contains(err.Message, "quota") {
	//	return true
	//}
	if strings.Contains(err.Message, "credit") {
		return true
	}
	if strings.Contains(err.Message, "balance") {
		return true
	}
	return false
}

func ShouldEnableChannel(err error, openAIErr *model.Error) bool {
	if !config.AutomaticEnableChannelEnabled {
		return false
	}
	if err != nil {
		return false
	}
	if openAIErr != nil {
		return false
	}
	return true
}
