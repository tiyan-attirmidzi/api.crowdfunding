package dto

import "github.com/tiyan-attirmidzi/api.crowdfunding/entities"

type CheckExistingEmail struct {
	Email string `json:"email" binding:"required,email"`
}

type UserFormatter struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Occupation string `json:"occupation"`
	Email      string `json:"email"`
}

func FormatUser(user entities.User, token string) UserFormatter {

	formatter := UserFormatter{
		ID:         user.ID,
		Name:       user.Name,
		Occupation: user.Occupation,
		Email:      user.Email,
	}

	return formatter

}
