package llmprovider

import "emergent/internal/common"

func NewClaudeProvider() LLMProviderIface {
	return &claude{}
}

type claude struct {
	ErrCounter int
}

func (c *claude) MakeRequest(request common.Request, count int) (common.Response, error) {
	resp, err := ClaudeProvider(request, count)
	if err != nil {
		println("claude server returned an error")
		c.ErrCounter++
	} else {
		println("making request from claude")
		c.ErrCounter = 0
	}

	return resp, err
}

func (c *claude) GetErrorCount() int {
	return c.ErrCounter
}

func (c *claude) IsHealthy() bool {
	return c.ErrCounter <= 5
}
