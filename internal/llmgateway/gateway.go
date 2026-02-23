package llmgateway

import (
	"emergent/internal/common"
	"emergent/internal/llmprovider"
	"math/rand"
)

type Gateway interface {
	SendRequest(request common.Request, count int) (common.Response, error)
}

func NewGateway() Gateway {
	claudeProvider := llmprovider.NewClaudeProvider()
	openAIProvider := llmprovider.NewOpenAIProvider()

	providers := make(map[common.Provider]llmprovider.LLMProviderIface)
	providers[common.Claude] = claudeProvider
	providers[common.OpenAI] = openAIProvider

	return &gateway{
		providers: providers,
	}
}

type gateway struct {
	providers map[common.Provider]llmprovider.LLMProviderIface
}

func (g *gateway) SendRequest(request common.Request, count int) (common.Response, error) {
	// primary := config.GetPrimaryProvider() // get primary provider from config
	// secondary := config.GetSecondaryProvider() // get secondary provider from config
	// hardcoding for now
	primary := common.Claude
	secondary := common.OpenAI

	if g.providers[primary].IsHealthy() {
		resp, err := g.providers[primary].MakeRequest(request, count)
		println(g.providers[primary].GetErrorCount())
		return resp, err
	} else {
		if rand.Intn(100) < 95 {
			openAIProvider := g.providers[secondary]
			return openAIProvider.MakeRequest(request, count)
		} else {
			claudeProvider := g.providers[primary]
			return claudeProvider.MakeRequest(request, count)
		}
	}
}

// func (g *gateway) SendRequest(request common.Request, count int) (common.Response, error) {
// 	// primary := config.GetPrimaryProvider() // get primary provider from config
// 	// secondary := config.GetSecondaryProvider() // get secondary provider from config
// 	// x := config.GetPrimaryConfiguredFailureThreshold()

// 	// hardcoding for now
// 	primary := common.Claude
// 	secondary := common.OpenAI

// 	claudeErrCount := g.providers[primary].GetErrorCount()
// 	if claudeErrCount <= 5 {
// 		req, err := g.providers[primary].MakeRequest(request, count)
// 		println(err)
// 		return req, err
// 	} else {
// 		req, err := g.randomizeRequest(count)
// 		println(err)
// 		return req, err
// 	}
// }

func (g *gateway) randomizeRequest(count int) (common.Response, error) {
	// 95% traffic to openAI, 5% to claude
	// rand.Intn(100) < 95
	num := rand.Intn(100)

	if num < 50 {
		openAIProvider := g.providers[common.OpenAI]
		return openAIProvider.MakeRequest(common.Request{}, count)
	} else {
		claudeProvider := g.providers[common.Claude]
		return claudeProvider.MakeRequest(common.Request{}, count)
	}
}
