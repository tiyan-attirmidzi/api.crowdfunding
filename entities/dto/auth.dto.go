package dto

import "github.com/tiyan-attirmidzi/api.crowdfunding/entities"

type SignUp struct {
	Name       string `json:"name" binding:"required"`
	Occupation string `json:"occupation" binding:"required"`
	Email      string `json:"email" binding:"required,email"`
	Password   string `json:"password" binding:"required"`
}

type SignIn struct {
	Email    string `json:"email" form:"email" binding:"required,email"`
	Password string `json:"password" form:"password" binding:"required"`
}

type SignedFormatter struct {
	User  UserFormatter `json:"user"`
	Token string        `json:"token"`
}

func FormatAuth(user entities.User, token string) SignedFormatter {

	formatter := SignedFormatter{
		User: UserFormatter{
			ID:         user.ID,
			Name:       user.Name,
			Occupation: user.Occupation,
			Email:      user.Email,
		},
		Token: token,
	}

	return formatter

}
