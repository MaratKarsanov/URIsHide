package internal

import (
	"URIsHide/internal/interfaces"
)

type HideURIsService struct {
	prod interfaces.Producer
	pres interfaces.Presenter
}

func NewHideURIsService(prod interfaces.Producer, pres interfaces.Presenter) *HideURIsService {
	return &HideURIsService{prod, pres}
}

func (service HideURIsService) Run() error {
	data, err := service.prod.Produce()
	if err != nil {
		return err
	}
	var result = make([]string, 0, len(data))

	for _, str := range data {
		var prefix = "http://"
		var prefixLength = len(prefix)
		var strLength = len(str)
		var buffer = []byte(str)
		var sentenceEndings = map[byte]bool{
			'.': true,
			'!': true,
			'?': true,
		}

		for i := 0; i < strLength-prefixLength; i++ {
			if str[i:i+prefixLength] == prefix {
				i += prefixLength
				for j := i; j < strLength; j++ {
					if str[j] == ' ' {
						if sentenceEndings[str[j-1]] {
							buffer[j-1] = str[j-1]
						}
						break
					}
					if j == strLength-1 && sentenceEndings[str[j]] {
						break
					}
					buffer[j] = '*'
				}
			}
		}

		result = append(result, string(buffer))
	}

	err = service.pres.Present(result)
	if err != nil {
		return err
	}

	return nil
}
