package internal

type AI interface {
	GetCompletion(prompt string) (string, error)
}
