package services

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/tiyan-attirmidzi/api.crowdfunding/entities"
	"github.com/tiyan-attirmidzi/api.crowdfunding/entities/dto"
	"github.com/tiyan-attirmidzi/api.crowdfunding/repositories"
	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	SignUp(user dto.SignUp) (entities.User, error)
	SignIn(user dto.SignIn) (entities.User, error)
	GenerateToken(data entities.User) (string, error)
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

type jwtCustomClaim struct {
	ID         int    `json:"user_id"`
	Name       string `json:"name"`
	Occupation string `json:"occupation"`
	Email      string `json:"email"`
	jwt.StandardClaims
}

// TODO: Change Later
var secretKey = []byte("S3c12e7_k3Y")

func (s *authService) GenerateToken(data entities.User) (string, error) {
	claims := &jwtCustomClaim{
		ID:         data.ID,
		Name:       data.Name,
		Occupation: data.Occupation,
		Email:      data.Email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(secretKey))

	if err != nil {
		return signedToken, err
	}

	return signedToken, nil

}
