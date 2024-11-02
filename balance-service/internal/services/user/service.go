package user

import "github.com/AkselRivera/stori-challenge/balance-service/internal/ports"

type Service struct {
	Repo ports.UserRepository
}
