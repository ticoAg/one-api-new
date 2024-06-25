package doubao

import (
	"fmt"
	"github.com/ticoAg/one-api-new/relay/meta"
	"github.com/ticoAg/one-api-new/relay/relaymode"
)

func GetRequestURL(meta *meta.Meta) (string, error) {
	if meta.Mode == relaymode.ChatCompletions {
		return fmt.Sprintf("%s/api/v3/chat/completions", meta.BaseURL), nil
	}
	return "", fmt.Errorf("unsupported relay mode %d for doubao", meta.Mode)
}
