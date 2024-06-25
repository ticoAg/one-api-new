package aiproxy

import "github.com/ticoAg/one-api-new/relay/adaptor/openai"

var ModelList = []string{""}

func init() {
	ModelList = openai.ModelList
}
