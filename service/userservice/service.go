package userservice

import (
	"fmt"
	"game-app/entity"
	"game-app/pkg/phonenumber"
)

type Repository interface {
	IsPhoneNumberUnique(phoneNumber string) (bool, error)
	Register(u entity.User) (entity.User, error)
}

type Service struct {
	repo Repository
}

type RegisterRequest struct {
	Name        string
	PhoneNumber string
}

type RegisterResponse struct {
	User entity.User
}

func New(repo Repository) Service {
	return Service{repo: repo}
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

	// create new user in storage
	user := entity.User{
		ID:          0,
		PhoneNumber: req.PhoneNumber,
		Name:        req.Name,
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
