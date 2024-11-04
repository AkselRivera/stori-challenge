package user_test

import (
	"errors"
	"testing"
	"time"

	"github.com/AkselRivera/stori-challenge/balance-service/mocks"
	"github.com/AkselRivera/stori-challenge/balance-service/pkg/domain"
	"github.com/AkselRivera/stori-challenge/balance-service/pkg/services/user"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestGet(t *testing.T) {

	testTable := map[string]struct {
		setup     func(mock *mocks.MockUserRepository)
		userID    int
		startDate time.Time
		endDate   time.Time
		assertion func(subTest *testing.T, balance domain.UserBalance, err error)
	}{
		"user not found": {
			setup: func(mock *mocks.MockUserRepository) {
				mock.EXPECT().GetUserTransactions(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil, domain.ErrorUserNotFound)
			},

			userID:    1,
			startDate: time.Date(0, 0, 0, 0, 0, 0, 0, time.UTC),
			endDate:   time.Date(0, 0, 0, 0, 0, 0, 0, time.UTC),

			assertion: func(subTest *testing.T, balance domain.UserBalance, err error) {

				assert.Equal(subTest, domain.UserBalance{}, balance)
				assert.NotNil(subTest, err)

				var customErr domain.CustomError
				if errors.As(err, &customErr) {
					assert.Equal(subTest, domain.ErrorCodeBadRequest, customErr.Code)
				}

			},
		},

		"invalid dates": {
			setup: func(mock *mocks.MockUserRepository) {
				mock.EXPECT().GetUserTransactions(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil, errors.New("invalid date"))
			},

			userID:    1,
			startDate: time.Date(2001, 04, 01, 0, 0, 0, 0, time.UTC),
			endDate:   time.Date(0, 0, 0, 0, 0, 0, 0, time.UTC),

			assertion: func(subTest *testing.T, balance domain.UserBalance, err error) {

				assert.Equal(subTest, domain.UserBalance{}, balance)
				assert.NotNil(subTest, err)

				var customErr domain.CustomError
				if errors.As(err, &customErr) {
					assert.Equal(subTest, domain.ErrorCodeInternalServerError, customErr.Code)
				}

			},
		},

		"balance specific user id": {
			setup: func(mock *mocks.MockUserRepository) {
				mock.EXPECT().GetUserTransactions(gomock.Any(), gomock.Any(), gomock.Any()).Return([]domain.Transaction{
					{
						ID:       1,
						UserID:   1,
						Amount:   10,
						DateTime: time.Date(2001, 04, 01, 0, 0, 0, 0, time.UTC),
					},
					{
						ID:       2,
						UserID:   1,
						Amount:   -15,
						DateTime: time.Date(2001, 04, 02, 0, 0, 0, 0, time.UTC),
					},
				}, nil)
			},

			userID:    1,
			startDate: time.Date(0, 0, 0, 0, 0, 0, 0, time.UTC),
			endDate:   time.Date(0, 0, 0, 0, 0, 0, 0, time.UTC),

			assertion: func(subTest *testing.T, balance domain.UserBalance, err error) {

				var expectedBalance = domain.UserBalance{
					Balance:      25,
					TotalCredits: 10,
					TotalDebits:  15,
				}

				assert.Nil(subTest, nil, err)
				assert.Equal(subTest, expectedBalance, balance)
			},
		},

		"balance specific user id and dates": {
			setup: func(mock *mocks.MockUserRepository) {
				mock.EXPECT().GetUserTransactions(gomock.Any(), gomock.Any(), gomock.Any()).Return([]domain.Transaction{
					{
						ID:       1,
						UserID:   1,
						Amount:   10,
						DateTime: time.Date(2001, 04, 01, 0, 0, 0, 0, time.UTC),
					},
					{
						ID:       2,
						UserID:   1,
						Amount:   -15,
						DateTime: time.Date(2001, 04, 02, 0, 0, 0, 0, time.UTC),
					},
					{
						ID:       3,
						UserID:   1,
						Amount:   20,
						DateTime: time.Date(2001, 04, 02, 07, 30, 0, 0, time.UTC),
					},
				}, nil)
			},

			userID:    1,
			startDate: time.Date(2001, 04, 01, 0, 0, 0, 0, time.UTC),
			endDate:   time.Date(2001, 04, 02, 0, 0, 0, 0, time.UTC),

			assertion: func(subTest *testing.T, balance domain.UserBalance, err error) {

				var expectedBalance = domain.UserBalance{
					Balance:      45,
					TotalCredits: 30,
					TotalDebits:  15,
				}

				assert.Nil(subTest, nil, err)
				assert.Equal(subTest, expectedBalance, balance)
			},
		},
	}

	for testName, testCase := range testTable {

		t.Run(testName, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			repoMock := mocks.NewMockUserRepository(ctrl)
			testCase.setup(repoMock)

			userSrv := user.Service{
				Repo: repoMock,
			}

			balance, err := userSrv.GetBalance(testCase.userID, testCase.startDate, testCase.endDate)

			testCase.assertion(t, balance, err)
		})
	}
}
