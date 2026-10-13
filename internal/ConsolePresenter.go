package internal

import "fmt"

type ConsolePresenter struct {
}

func NewConsolePresenter() *ConsolePresenter {
	return &ConsolePresenter{}
}

func (presenter *ConsolePresenter) Present(data []string) error {
	for _, line := range data {
		_, err := fmt.Println(line)
		if err != nil {
			return err
		}
	}

	return nil
}
