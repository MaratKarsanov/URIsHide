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

	var presenter = internal.NewFilePresenter(outputDir)
	var producer = internal.NewFileProducer(inputDir)
	var service = internal.NewHideURIsService(producer, presenter)

	var err = service.Run()
	if err != nil {
		fmt.Println("При работе сервиса произошла ошибка: ", err)
		return
	}

	fmt.Println("Приложение успешно завершило работу")
}
