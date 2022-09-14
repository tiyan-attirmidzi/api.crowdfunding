package services

import (
	"github.com/tiyan-attirmidzi/api.crowdfunding/entities"
	"github.com/tiyan-attirmidzi/api.crowdfunding/entities/dto"
	"github.com/tiyan-attirmidzi/api.crowdfunding/repositories"
	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	SignUp(user dto.SignUp) (entities.User, error)
}

type authService struct {
	userRepository repositories.UserRepository
}

func NewAuthService(userRepository repositories.UserRepository) *authService {
	return &authService{userRepository}
}

func (s *authService) SignUp(data dto.SignUp) (entities.User, error) {

	user := entities.User{}
	user.Name = data.Name
	user.Email = data.Email
	user.Occupation = data.Occupation
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(data.Password), bcrypt.MinCost)

	if err != nil {
		return user, err
	}

	user.PasswordHash = string(passwordHash) // byte => string
	user.Role = "user"

	newUser, err := s.userRepository.Save(user)

	if err != nil {
		return newUser, err
	}

	return newUser, nil

}
