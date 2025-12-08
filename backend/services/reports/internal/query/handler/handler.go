package handler

type IHandlerService interface{}

type HandlerService struct{}

func NewHandlerService() IHandlerService {
	return HandlerService{}
}
