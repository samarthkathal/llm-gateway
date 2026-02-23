package main

import (
	"emergent/internal/common"
	"emergent/internal/llmgateway"
	"time"
)

func main() {

	g := llmgateway.NewGateway()

	request := common.Request{}

	for count := 1; count <= 100; count++ {
		_, err := g.SendRequest(request, count)
		time.Sleep(100 * time.Millisecond) // add some delay between requests
		if err != nil {
			println("error calling the requst")
		}
	}
}

// LLM Gateway (like an outbound request filter/middleware)

// if x consecutive failures at primary, switch 95% traffic to secondary, and 5% to primary.
// if there is y consecutive successes at primary, switch back 100% traffic to primary.
