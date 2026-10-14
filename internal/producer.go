package internal

type Producer interface {
	Produce() ([]string, error)
}
