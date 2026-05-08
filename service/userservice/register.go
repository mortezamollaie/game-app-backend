package userservice

import (
	"fmt"
	"game-app/entity"
	"game-app/param"

	"golang.org/x/crypto/bcrypt"
)

func (s Service) Register(req param.RegisterRequest) (param.RegisterResponse, error) {
	// TODO - we should verify phone number by verification code

	pass := []byte(req.Password)
	pass, err := bcrypt.GenerateFromPassword(pass, 0)
	if err != nil {
		return param.RegisterResponse{}, err
	}

	// create new user_handler in storage
	user := entity.User{
		ID:          0,
		PhoneNumber: req.PhoneNumber,
		Name:        req.Name,
		Password:    string(pass),
	}

	createdUser, err := s.repo.Register(user)
	if err != nil {
		return param.RegisterResponse{}, fmt.Errorf("unexpected err: %w", err)
	}

	return param.RegisterResponse{User: struct {
		ID          uint   `json:"id"`
		Name        string `json:"name"`
		PhoneNumber string `json:"phone_number"`
	}{ID: createdUser.ID, PhoneNumber: createdUser.PhoneNumber, Name: createdUser.Name}}, nil
}
