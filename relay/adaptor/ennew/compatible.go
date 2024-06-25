package ennew

import (
	"github.com/ticoAg/one-api-new/relay/adaptor/ai360"
	"github.com/ticoAg/one-api-new/relay/adaptor/baichuan"
	"github.com/ticoAg/one-api-new/relay/adaptor/deepseek"
	"github.com/ticoAg/one-api-new/relay/adaptor/doubao"
	"github.com/ticoAg/one-api-new/relay/adaptor/groq"
	"github.com/ticoAg/one-api-new/relay/adaptor/lingyiwanwu"
	"github.com/ticoAg/one-api-new/relay/adaptor/minimax"
	"github.com/ticoAg/one-api-new/relay/adaptor/mistral"
	"github.com/ticoAg/one-api-new/relay/adaptor/moonshot"
	"github.com/ticoAg/one-api-new/relay/adaptor/stepfun"
	"github.com/ticoAg/one-api-new/relay/adaptor/togetherai"
	"github.com/ticoAg/one-api-new/relay/channeltype"
	// "github.com/ticoAg/one-api-new/relay/adaptor/ennew"
)

var CompatibleChannels = []int{
	channeltype.Azure,
	channeltype.AI360,
	channeltype.Moonshot,
	channeltype.Baichuan,
	channeltype.Minimax,
	channeltype.Doubao,
	channeltype.Mistral,
	channeltype.Groq,
	channeltype.LingYiWanWu,
	channeltype.StepFun,
	channeltype.DeepSeek,
	channeltype.TogetherAI,
	channeltype.Enniu,
}

func GetCompatibleChannelMeta(channelType int) (string, []string) {
	switch channelType {
	case channeltype.Azure:
		return "azure", ModelList
	case channeltype.AI360:
		return "360", ai360.ModelList
	case channeltype.Moonshot:
		return "moonshot", moonshot.ModelList
	case channeltype.Baichuan:
		return "baichuan", baichuan.ModelList
	case channeltype.Minimax:
		return "minimax", minimax.ModelList
	case channeltype.Mistral:
		return "mistralai", mistral.ModelList
	case channeltype.Groq:
		return "groq", groq.ModelList
	case channeltype.LingYiWanWu:
		return "lingyiwanwu", lingyiwanwu.ModelList
	case channeltype.StepFun:
		return "stepfun", stepfun.ModelList
	case channeltype.DeepSeek:
		return "deepseek", deepseek.ModelList
	case channeltype.TogetherAI:
		return "together.ai", togetherai.ModelList
	case channeltype.Doubao:
		return "doubao", doubao.ModelList
	case channeltype.Enniu:
		return "ennew", togetherai.ModelList
	default:
		return "openai", ModelList
	}
}
