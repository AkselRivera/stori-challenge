package sender

import (
	"github.com/AkselRivera/stori-challenge/migration-service/pkg/ports"
	"github.com/resend/resend-go/v2"
)

type ResendEmailSender struct {
	Client *resend.Client
}

var _ ports.EmailService = &ResendEmailSender{}
