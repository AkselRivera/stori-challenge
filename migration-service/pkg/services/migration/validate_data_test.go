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

func TestService_Validate_Data(t *testing.T) {
	testTable := map[string]struct {
		setup         func(mock *mocks.MockMigrationRepository)
		data          [][]string
		assertionFunc func(subTest *testing.T, transaction []*domain.Transaction, err error)
	}{
		"empty data": {
			setup: func(mock *mocks.MockMigrationRepository) {
			},
			data: [][]string{},
			assertionFunc: func(subTest *testing.T, transactions []*domain.Transaction, err error) {

				assert.Nil(subTest, transactions)
				assert.NotNil(subTest, err, "expected error, got nil")

				var customErr domain.CustomError

				if errors.As(err, &customErr) {
					assert.Equal(subTest, domain.ErrorCodeBadRequest, customErr.Code)
				}

			},
		},

		"invalid user id": {
			setup: func(mock *mocks.MockMigrationRepository) {},

			data: [][]string{
				{"1", "1A", "1.1", "2021-01-01T00:00:00Z"},
			},

			assertionFunc: func(subTest *testing.T, transactions []*domain.Transaction, err error) {
				assert.Nil(t, transactions)

				var customErr domain.CustomError

				if errors.As(err, &customErr) {
					assert.Equal(t, domain.ErrorCodeBadRequest, customErr.Code)
				}
			},
		},

		"invalid amount": {
			setup: func(mock *mocks.MockMigrationRepository) {},

			data: [][]string{
				{"1", "1", "1.1A", "2021-01-01T00:00:00Z"},
			},

			assertionFunc: func(subTest *testing.T, transactions []*domain.Transaction, err error) {
				assert.Nil(t, transactions)

				var customErr domain.CustomError

				if errors.As(err, &customErr) {
					assert.Equal(t, domain.ErrorCodeBadRequest, customErr.Code)
				}
			},
		},

		"invalid datetime": {
			setup: func(mock *mocks.MockMigrationRepository) {},

			data: [][]string{
				{"1", "1", "1.1", "2021-01-01"},
			},

			assertionFunc: func(subTest *testing.T, transactions []*domain.Transaction, err error) {
				assert.Nil(t, transactions)

				var customErr domain.CustomError

				if errors.As(err, &customErr) {
					assert.Equal(t, domain.ErrorCodeBadRequest, customErr.Code)
				}
			},
		},

		"correct data types": {
			setup: func(mock *mocks.MockMigrationRepository) {},

			data: [][]string{
				{"1", "1", "1.1", "2021-01-01T00:00:00Z"},
				{"2", "2", "2.2", "2021-01-01T00:00:00Z"},
				{"3", "3", "3.3", "2021-01-01T00:00:00Z"},
			},

			assertionFunc: func(subTest *testing.T, transactions []*domain.Transaction, err error) {
				var expectedTransactions = []*domain.Transaction{
					{UserID: 1, Amount: 1.1, DateTime: time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC)},
					{UserID: 2, Amount: 2.2, DateTime: time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC)},
					{UserID: 3, Amount: 3.3, DateTime: time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC)},
				}
				assert.Nil(t, err)
				assert.Equal(t, expectedTransactions, transactions)
			},
		},
	}

	for name, test := range testTable {
		t.Run(name, func(subTest *testing.T) {

			ctrl := gomock.NewController(subTest)
			defer ctrl.Finish()

			mockRepo := mocks.NewMockMigrationRepository(ctrl)
			test.setup(mockRepo)

			s := &migration.Service{
				Repo:   mockRepo,
				Sender: nil,
			}

			transactions, err := s.ValidateData(test.data)

			test.assertionFunc(subTest, transactions, err)
		})
	}
}
