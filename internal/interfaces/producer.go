package interfaces

type Producer interface {
	Produce() ([]string, error)
}
