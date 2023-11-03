package services

type Service struct {
	Responder
}

func NewService() *Service {
	return &Service{
		Responder: NewSingleResponder(),
	}
}
