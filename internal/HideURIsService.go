package internal

import (
	"URIsHide/internal/interfaces"
)

type HideURIsService struct {
	prod interfaces.Producer
	pres interfaces.Presenter
}
