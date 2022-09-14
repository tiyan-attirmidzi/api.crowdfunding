package repositories

import (
	"github.com/tiyan-attirmidzi/api.crowdfunding/entities"
	"gorm.io/gorm"
)

type UserRepository interface {
	Save(user entities.User) (entities.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{db}
}

func (r *userRepository) Save(user entities.User) (entities.User, error) {
	err := r.db.Create(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}
