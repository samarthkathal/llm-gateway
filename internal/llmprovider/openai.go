package llmprovider

import "emergent/internal/common"

func NewOpenAIProvider() LLMProviderIface {
	return &openai{}
}

type openai struct {
	ErrCounter int
}

func (o *openai) MakeRequest(request common.Request, count int) (common.Response, error) {
	resp, err := OpenAIProvider(request, count)
	if err != nil {
		o.ErrCounter++
	} else {
		o.ErrCounter = 0
	}

	println("making request from openai")
	return resp, err
}

func (o *openai) GetErrorCount() int {
	return o.ErrCounter
}

func (o *openai) IsHealthy() bool {
	return o.ErrCounter <= 5
}
