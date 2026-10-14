package internal

import (
	"bufio"
	"os"
)

type FilePresenter struct {
	filePath string
}

func NewFilePresenter(filePath string) *FilePresenter {
	return &FilePresenter{
		filePath: filePath,
	}
}

func (p *FilePresenter) Present(data []string) error {
	file, err := os.Create(p.filePath)
	if err != nil {
		return err
	}

	defer func() {
		closingError := file.Close()
		if closingError != nil {
			/*здесь должна быть запись в лог*/
		}
	}()

	writer := bufio.NewWriter(file)
	for _, line := range data {
		if _, err := writer.WriteString(line + "\n"); err != nil {
			return err
		}
	}

	if err := writer.Flush(); err != nil {
		return err
	}

	return nil
}
