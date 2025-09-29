package user

import (
	"authservice-sandbox/pkg/entities"

	"gorm.io/gorm"
)

// TODO: Repository itu bertanggung jawab sebagai operasi database untuk entitas User
type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (r *UserRepository) FindUserByUsername(username string) (*entities.User, error) {
	var user entities.User
	// TODO: SELECT * FROM users WHERE username = ?
	if err := r.db.Where("username = ?", username).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) SaveUser(user *entities.User) error {
	if err := r.db.Save(user).Error; err != nil {
		return err
	}
	return nil
}
