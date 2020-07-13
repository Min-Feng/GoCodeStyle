package domain

import "ddd/gopkg2/domain/mail"

func DoThing(svc mail.Service) {
	svc.Send()
}
