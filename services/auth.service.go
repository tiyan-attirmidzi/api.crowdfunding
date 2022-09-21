package services

import (
	"errors"

	"github.com/tiyan-attirmidzi/api.crowdfunding/entities"
	"github.com/tiyan-attirmidzi/api.crowdfunding/entities/dto"
	"github.com/tiyan-attirmidzi/api.crowdfunding/repositories"
	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	SignUp(user dto.SignUp) (entities.User, error)
	SignIn(user dto.SignIn) (entities.User, error)
	IsEmailAvailable(user dto.CheckEmail) (bool, error)
	SaveAvatar(ID int, fileLocation string) (entities.User, error)
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

func (s *authService) SignIn(data dto.SignIn) (entities.User, error) {

	email := data.Email
	password := data.Password

	user, err := s.userRepository.FindByEmail(email)

	if err != nil {
		return user, err
	}

	if user.ID == 0 {
		return user, errors.New("no user found on that email")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))

	if err != nil {
		return user, errors.New("password incorrect")
	}

	return user, nil

}

func (s *authService) IsEmailAvailable(data dto.CheckEmail) (bool, error) {

	email := data.Email

	user, err := s.userRepository.FindByEmail(email)

	if err != nil {
		return false, err
	}

	if user.ID == 0 {
		return true, nil
	}

	return false, errors.New("email has been registered")

}

func (s *authService) SaveAvatar(ID int, fileLocation string) (entities.User, error) {

	user, err := s.userRepository.FindByID(ID)

	if err != nil {
		return user, err
	}

	user.AvatarFileName = fileLocation

	updatedUser, err := s.userRepository.Update(user)

	if err != nil {
		return updatedUser, err
	}

	return updatedUser, nil
}
