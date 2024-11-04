package user

import "github.com/AkselRivera/stori-challenge/balance-service/pkg/ports"

type Service struct {
	Repo ports.UserRepository
}
