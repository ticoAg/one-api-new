package relay

import (
	"github.com/ticoAg/one-api-new/relay/adaptor"
	"github.com/ticoAg/one-api-new/relay/adaptor/aiproxy"
	"github.com/ticoAg/one-api-new/relay/adaptor/ali"
	"github.com/ticoAg/one-api-new/relay/adaptor/anthropic"
	"github.com/ticoAg/one-api-new/relay/adaptor/aws"
	"github.com/ticoAg/one-api-new/relay/adaptor/baidu"
	"github.com/ticoAg/one-api-new/relay/adaptor/cloudflare"
	"github.com/ticoAg/one-api-new/relay/adaptor/cohere"
	"github.com/ticoAg/one-api-new/relay/adaptor/coze"
	"github.com/ticoAg/one-api-new/relay/adaptor/deepl"
	// "github.com/ticoAg/one-api-new/relay/adaptor/ennew"
	// "/home/tico/workspace/one-api-new/relay/adaptor/ennew"
	"github.com/ticoAg/one-api-new/relay/adaptor/ennew"
	"github.com/ticoAg/one-api-new/relay/adaptor/gemini"
	"github.com/ticoAg/one-api-new/relay/adaptor/ollama"
	"github.com/ticoAg/one-api-new/relay/adaptor/openai"
	"github.com/ticoAg/one-api-new/relay/adaptor/palm"
	"github.com/ticoAg/one-api-new/relay/adaptor/tencent"
	"github.com/ticoAg/one-api-new/relay/adaptor/xunfei"
	"github.com/ticoAg/one-api-new/relay/adaptor/zhipu"
	"github.com/ticoAg/one-api-new/relay/apitype"
)

func GetAdaptor(apiType int) adaptor.Adaptor {
	switch apiType {
	case apitype.AIProxyLibrary:
		return &aiproxy.Adaptor{}
	case apitype.Ali:
		return &ali.Adaptor{}
	case apitype.Anthropic:
		return &anthropic.Adaptor{}
	case apitype.AwsClaude:
		return &aws.Adaptor{}
	case apitype.Baidu:
		return &baidu.Adaptor{}
	case apitype.Gemini:
		return &gemini.Adaptor{}
	case apitype.OpenAI:
		return &openai.Adaptor{}
	case apitype.PaLM:
		return &palm.Adaptor{}
	case apitype.Tencent:
		return &tencent.Adaptor{}
	case apitype.Xunfei:
		return &xunfei.Adaptor{}
	case apitype.Zhipu:
		return &zhipu.Adaptor{}
	case apitype.Ollama:
		return &ollama.Adaptor{}
	case apitype.Coze:
		return &coze.Adaptor{}
	case apitype.Cohere:
		return &cohere.Adaptor{}
	case apitype.Cloudflare:
		return &cloudflare.Adaptor{}
	case apitype.DeepL:
		return &deepl.Adaptor{}
	case apitype.Enniu:
		return &ennew.Adaptor{}
	}
	return nil
}
