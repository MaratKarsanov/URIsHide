package internal

import (
	"bufio"
	"os"
)

type FileProducer struct {
	filePath string
}

func NewFileProducer(filePath string) *FileProducer {
	return &FileProducer{
		filePath: filePath,
	}
}

func (producer *FileProducer) Produce() ([]string, error) {
	file, err := os.OpenFile(producer.filePath, os.O_RDONLY|os.O_CREATE, 0644)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return lines, nil
}
