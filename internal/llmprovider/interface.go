package llmprovider

import "emergent/internal/common"

type LLMProviderIface interface {
	MakeRequest(request common.Request, count int) (common.Response, error)
	GetErrorCount() int
	IsHealthy() bool
}
