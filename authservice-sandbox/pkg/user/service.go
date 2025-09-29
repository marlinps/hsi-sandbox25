package user

import (
	"authservice-sandbox/pkg/entities"
)

type AutRepository interface {
	FindUserByUsername(username string) (*entities.User, error)
	SaveUser(user *entities.User) error
}

type TokenGenerator interface {
	GenerateAccessToken(username string) (string, error)
	ValidateToken(token string) (string, error)
}

type UserService struct {
	repo           AutRepository
	tokenGenerator TokenGenerator
}

func NewUserService(repo AutRepository, tokenGen TokenGenerator) *UserService {
	return &UserService{
		repo:           repo,
		tokenGenerator: tokenGen,
	}
}

func (s *UserService) Login(username, password string) (string, error) {
	// TODO: 1. Cari user berdasarkan username
	user, err := s.repo.FindUserByUsername(username)
	if err != nil {
		// TODO: 2. Jika user tidak ditemukan
		return "Invalid Credential", err
	}

	// TODO: 3. Jika user ditemukan, maka membuat token JWT
	accessToken, err := s.tokenGenerator.GenerateAccessToken(user.Username)
	if err != nil {
		return "Failed to generate token", err
	}
	return accessToken, nil
}
