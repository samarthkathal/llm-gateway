package common

type Request struct {
}
type Response struct {
}

type Provider string

const (
	Claude Provider = "Claude"
	OpenAI Provider = "OpenAI"
)
