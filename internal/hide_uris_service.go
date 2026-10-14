package internal

type HideURIsService struct {
	prod Producer
	pres Presenter
}

func NewHideURIsService(prod Producer, pres Presenter) *HideURIsService {
	return &HideURIsService{prod, pres}
}

func (s *HideURIsService) Run() error {
	data, err := s.prod.Produce()
	if err != nil {
		return err
	}
	result := make([]string, 0, len(data))

	for _, str := range data {
		result = append(result, hideURIs(str))
	}

	err = s.pres.Present(result)
	if err != nil {
		return err
	}

	return nil
}

func hideURIs(str string) string {
	prefix := "http://"
	prefixLength := len(prefix)
	strLength := len(str)
	buffer := []byte(str)
	sentenceEndings := map[byte]bool{
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

	return string(buffer)
}
