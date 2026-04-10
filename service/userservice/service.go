package userservice

import (
	"fmt"
	"game-app/entity"
	"game-app/pkg/phonenumber"

	"golang.org/x/crypto/bcrypt"
)

type Repository interface {
	IsPhoneNumberUnique(phoneNumber string) (bool, error)
	Register(u entity.User) (entity.User, error)
}

type Service struct {
	repo Repository
}

func New(repo Repository) Service {
	return Service{repo: repo}
}

type RegisterRequest struct {
	Name        string `json:"name"`
	PhoneNumber string `json:"phone_number"`
	Password    string `json:"password"`
}

type RegisterResponse struct {
	User entity.User
}

func (s Service) Register(req RegisterRequest) (RegisterResponse, error) {
	// TODO - we should verify phone number by verification code
	// validate phone number
	if !phonenumber.IsValid(req.PhoneNumber) {
		return RegisterResponse{}, fmt.Errorf("phone number %s is not valid", req.PhoneNumber)
	}

	// check uniqueness of phone number
	if isUnique, err := s.repo.IsPhoneNumberUnique(req.PhoneNumber); err != nil || !isUnique {
		if err != nil {
			return RegisterResponse{}, fmt.Errorf("unexpected err: %w", err)
		}

		return RegisterResponse{}, fmt.Errorf("phone number %s is not unique", req.PhoneNumber)
	}

	// validate name
	if len(req.Name) < 3 {
		return RegisterResponse{}, fmt.Errorf("name lenght should be at least 3 characters long")
	}

	// TODO - check the password with the regex pattern
	// validate password
	if len(req.PhoneNumber) < 8 {
		return RegisterResponse{}, fmt.Errorf("phone_number lenght should be at least 8 characters long")
	}

	pass := []byte(req.Password)
	pass, err := bcrypt.GenerateFromPassword(pass, 0)
	if err != nil {
		return RegisterResponse{}, err
	}

	// create new user in storage
	user := entity.User{
		ID:          0,
		PhoneNumber: req.PhoneNumber,
		Name:        req.Name,
		Password:    string(pass),
	}

	createdUser, err := s.repo.Register(user)
	if err != nil {
		return RegisterResponse{}, fmt.Errorf("unexpected err: %w", err)
	}

	// return created user
	return RegisterResponse{
		User: createdUser,
	}, nil
}
