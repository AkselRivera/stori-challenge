package migration

import (
	"fmt"

	"github.com/AkselRivera/stori-challenge/migration-service/pkg/domain"
)

func (s *Service) Migrate(transactions []*domain.Transaction, fileName string) (*domain.EmailData, error) {

	title := "Good news, successful transactions migration!"
	summay := fmt.Sprintf("Migrated <strong>%d</strong> transactions", len(transactions))

	var globarlErr error

	if err := s.Repo.InsertMany(transactions); err != nil {
		domain.HandleError(err, "failed to insert transactions")
		title = "Bad news, failed transactions migration!"
		summay = "Could not migrate transactions"

		globarlErr = domain.HandleError(err, "failed to insert transactions")
	}

	body := fmt.Sprintf("<h1>FinTech Solutions Inc.</h1> <h2>Summary:</h2> <p>%s from <strong>%s</strong> file</p> <br/><br/>  <p>FinTech Solutions Inc. Development Team</p>", summay, fileName)

	emailData := domain.EmailData{
		To:      []string{"moralesaksel@gmail.com"},
		Subject: title,
		Body:    body,
	}

	s.Sender.Send(emailData)

	return &emailData, globarlErr

}
