package mail

import (
	"gopkg.in/gomail.v2"

	"ddd/gopkg2/domain"
)

type Service struct {
	Address string
}

func NewService(cfg domain.Config) *Service {
	address := cfg.MailAddress
	return &Service{Address: address}
}

func (svc Service) Send() {
	gomail.NewDialer(svc.Address, 587, "user", "123456")
	// do something
}
