package main

import (
	"URIsHide/internal"
	"fmt"
	"os"
)

const (
	defaultInputDir  = "./input"
	defaultOutputDir = "./output"
)

func main() {
	inputDir := defaultInputDir
	outputDir := defaultOutputDir

	if len(os.Args) > 1 {
		inputDir = os.Args[1]
	}
	if len(os.Args) > 2 {
		outputDir = os.Args[2]
	}

	presenter := internal.NewFilePresenter(outputDir)
	producer := internal.NewFileProducer(inputDir)
	service := internal.NewHideURIsService(producer, presenter)

	err := service.Run()
	if err != nil {
		fmt.Println("При работе сервиса произошла ошибка: ", err)
		return
	}

	fmt.Println("Приложение успешно завершило работу")
}
