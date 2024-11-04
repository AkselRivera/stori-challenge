package sender

import (
	"fmt"

	"github.com/AkselRivera/stori-challenge/migration-service/pkg/domain"
	"github.com/resend/resend-go/v2"
)

func (r *ResendEmailSender) Send(emailData domain.EmailData) error {

	if len(emailData.To) == 0 {
		return domain.HandleError(fmt.Errorf("email data is invalid"), "failed to send email")
	}

	if emailData.ReplyTo == "" {
		emailData.ReplyTo = "moralesaksel@gmail.com"
	}

	params := &resend.SendEmailRequest{
		From:    "FinTech Solutions <onboarding@resend.dev>",
		To:      emailData.To,
		Html:    emailData.Body,
		Subject: emailData.Subject,
		ReplyTo: emailData.ReplyTo,
	}

	_, err := r.Client.Emails.Send(params)

	if err != nil {
		return domain.HandleError(err, "failed to send email")
	}

	return nil
}
