package migration_test

import (
	"errors"
	"testing"
	"time"

	"github.com/AkselRivera/stori-challenge/migration-service/mocks"
	"github.com/AkselRivera/stori-challenge/migration-service/pkg/domain"
	"github.com/AkselRivera/stori-challenge/migration-service/pkg/services/migration"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestMigrate(t *testing.T) {

	testTable := map[string]struct {
		setup         func(mock *mocks.MockMigrationRepository, mockEmailService *mocks.MockEmailService)
		data          []*domain.Transaction
		assertionFunc func(subTest *testing.T, email *domain.EmailData, err error)
	}{

		"migration successful": {
			setup: func(mock *mocks.MockMigrationRepository, mockEmailService *mocks.MockEmailService) {
				mock.EXPECT().InsertMany(gomock.Any()).Do(func(t []*domain.Transaction) {
					t[0].ID = 1
					t[0].UserID = 1
					t[0].Amount = 1.1
					t[0].DateTime = time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC)
				}).Return(nil)

				mockEmailService.EXPECT().Send(gomock.Any()).Return(nil)
			},
			data: []*domain.Transaction{
				{
					UserID:   1,
					Amount:   1.1,
					DateTime: time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
				},
			},
			assertionFunc: func(subTest *testing.T, email *domain.EmailData, err error) {
				assert.Nil(subTest, err)
				assert.NotNil(subTest, email)

				assert.Contains(subTest, "Good news, successful transactions migration!", email.Subject)
			},
		},

		"migration failed": {
			setup: func(mock *mocks.MockMigrationRepository, mockEmailService *mocks.MockEmailService) {
				mock.EXPECT().InsertMany(gomock.Any()).Do(func(t []*domain.Transaction) {

				}).Return(errors.New("generic error"))

				mockEmailService.EXPECT().Send(gomock.Any()).Return(nil)
			},

			data: []*domain.Transaction{
				{
					UserID:   1,
					Amount:   1.1,
					DateTime: time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
				},
			},

			assertionFunc: func(subTest *testing.T, emailData *domain.EmailData, err error) {
				assert.NotNil(subTest, err)
				assert.Contains(subTest, emailData.Subject, "Bad news, failed transactions migration!")

				var customErr domain.CustomError
				if errors.Is(err, &customErr) {
					assert.Equal(subTest, domain.ErrorCodeInternalServerError, customErr.Code)
				}
			},
		},
	}

	for name, test := range testTable {

		t.Run(name, func(t *testing.T) {

			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockRepo := mocks.NewMockMigrationRepository(ctrl)
			mockEmailSrv := mocks.NewMockEmailService(ctrl)
			test.setup(mockRepo, mockEmailSrv)

			s := migration.Service{Repo: mockRepo, Sender: mockEmailSrv}

			emailData, err := s.Migrate(test.data, "file.csv")

			test.assertionFunc(t, emailData, err)

		})
	}
}
