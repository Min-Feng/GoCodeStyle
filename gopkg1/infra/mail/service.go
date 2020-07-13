package mail

import (
	"gopkg.in/gomail.v2"

	"ddd/gopkg1/infra/configs"
)

type service struct {
	Address string
}

func NewService(cfg configs.Config) *service {
	address := cfg.MailAddress
	return &service{Address: address}
}

func (svc service) Send() {
	gomail.NewDialer(svc.Address, 587, "user", "123456")
	// do something
}
