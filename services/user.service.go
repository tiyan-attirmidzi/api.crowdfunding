package services

import (
	"errors"

	"github.com/tiyan-attirmidzi/api.crowdfunding/entities"
	"github.com/tiyan-attirmidzi/api.crowdfunding/repositories"
)

type UserService interface {
	IsEmailAvailable(email string) (bool, error)
	SaveAvatar(ID int, fileLocation string) (entities.User, error)
}

type userService struct {
	userRepository repositories.UserRepository
}

func NewUserService(userRepository repositories.UserRepository) *userService {
	return &userService{userRepository}
}

func (s *userService) IsEmailAvailable(email string) (bool, error) {

	user, err := s.userRepository.FindByEmail(email)

	if err != nil {
		return false, err
	}

	if user.ID == 0 {
		return true, nil
	}

	return false, errors.New("email has been registered")

}

func (s *userService) SaveAvatar(ID int, fileLocation string) (entities.User, error) {

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
