package llm

type LLM interface {
	Transform(text, tone string) (string, error)
}