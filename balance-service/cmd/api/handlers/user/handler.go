package user

import "github.com/AkselRivera/stori-challenge/balance-service/internal/ports"

type Handler struct {
	UserService ports.UserService
}
