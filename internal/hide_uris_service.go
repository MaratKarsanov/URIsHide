package internal

type HideURIsJob struct {
	index int
	str   string
}

type HideURIsResult struct {
	index int
	str   string
}

type HideURIsService struct {
	prod Producer
	pres Presenter
}

func NewHideURIsService(prod Producer, pres Presenter) *HideURIsService {
	return &HideURIsService{prod, pres}
}

func (s *HideURIsService) Run() error {
	const workersCount = 10

	data, err := s.prod.Produce()
	if err != nil {
		return err
	}

	pool := NewWorkerPool(workersCount, hideURIs)

	jobs := make(chan HideURIsJob)

	go func() {
		defer close(jobs)
		for i, str := range data {
			jobs <- HideURIsJob{index: i, str: str}
		}
	}()

	go pool.Process(jobs)

	result := make([]string, len(data))
	for r := range pool.Results() {
		result[r.index] = r.str
	}

	return s.pres.Present(result)
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
