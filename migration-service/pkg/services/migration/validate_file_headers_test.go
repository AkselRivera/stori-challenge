package migration_test

import (
	"errors"
	"testing"

	"github.com/AkselRivera/stori-challenge/migration-service/mocks"
	"github.com/AkselRivera/stori-challenge/migration-service/pkg/domain"
	"github.com/AkselRivera/stori-challenge/migration-service/pkg/services/migration"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestService_Validate_File_Headers(t *testing.T) {

	testTable := map[string]struct {
		setup         func(mock *mocks.MockMigrationRepository)
		data          []string
		assertionFunc func(subTest *testing.T, err error)
	}{

		"less than 3 columns": {
			setup: func(mock *mocks.MockMigrationRepository) {},
			data:  []string{"column1", "column2"},

			assertionFunc: func(subTest *testing.T, err error) {

				assert.NotNil(subTest, err)

				var customErr domain.CustomError
				if errors.As(err, &customErr) {
					assert.Equal(subTest, domain.ErrorCodeBadRequest, customErr.Code)
				}
			},
		},

		"invalid headers": {
			setup: func(mock *mocks.MockMigrationRepository) {},
			data:  []string{"column1", "column2", "column3", "column4"},

			assertionFunc: func(subTest *testing.T, err error) {
				assert.NotNil(subTest, err)

				var customErr domain.CustomError
				if errors.As(err, &customErr) {
					assert.Equal(subTest, domain.ErrorCodeBadRequest, customErr.Code)
				}
			},
		},

		"valid headers": {
			setup: func(mock *mocks.MockMigrationRepository) {},
			data:  []string{"id", "user_id", "amount", "datetime"},

			assertionFunc: func(subTest *testing.T, err error) {

				assert.Nil(subTest, err)

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
				Repo:         mockRepo,
				EmailService: nil,
			}

			err := s.ValidateFileHeaders(test.data)

			test.assertionFunc(subTest, err)

		})

	}
}
