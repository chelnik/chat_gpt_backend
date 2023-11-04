package services

type Service struct {
	ResponseService
}

func NewService() *Service {
	return &Service{
		ResponseService: *NewResponseService(),
	}
}
