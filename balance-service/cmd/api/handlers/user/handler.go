package user

import "github.com/AkselRivera/stori-challenge/balance-service/pkg/ports"

type Handler struct {
	UserService ports.UserService
}
