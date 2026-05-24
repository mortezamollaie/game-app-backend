package uservalidator

import "game-app/entity"

const (
	phoneNumberRegex = "^09[0-9]{9}$"
	passwordRegex    = "^[A-Z]{10,}$"
)

type Repository interface {
	IsPhoneNumberUnique(phoneNumber string) (bool, error)
	GetUserByPhoneNumber(phoneNumber string) (entity.User, error)
}

type Validator struct {
	repo Repository
}

func New(repo Repository) Validator {
	return Validator{repo: repo}
}
