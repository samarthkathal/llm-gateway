package llmprovider

import (
	"emergent/internal/common"
	"errors"
)

func ClaudeProvider(request common.Request, count int) (common.Response, error) {
	return common.Response{}, errors.New("Claude API error")
	// return common.Response{}, nil
}

func OpenAIProvider(request common.Request, count int) (common.Response, error) {
	return common.Response{}, nil
}
