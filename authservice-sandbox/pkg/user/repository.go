package user

import (
	"authservice-sandbox/pkg/entities"
)

type AuthRepository interface {
	FindUserByUsername(username string) (*entities.User, error)
	SaveUser(user *entities.User) error
	FindUserByRefreshToken(token string) (*entities.User, error)
}
