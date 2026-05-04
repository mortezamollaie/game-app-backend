package userservice

import (
	"fmt"
	"game-app/dto"
	"game-app/pkg/richerror"

	"golang.org/x/crypto/bcrypt"
)

func (s Service) Login(req dto.LoginRequest) (dto.LoginResponse, error) {
	// TODO - it would be better to user_handler two separate method for existence check and getUserByPhoneNumber
	const op = "userservice.Login"

	// check the existence of phone number from repository
	// get the user_handler by phone number

	user, err := s.repo.GetUserByPhoneNumber(req.PhoneNumber)

	if err != nil {
		return dto.LoginResponse{}, richerror.New(op).WithErr(err).WithMeta(map[string]interface{}{"phone_number": req.PhoneNumber})
	}
	accessToken, err := s.auth.CreateAccessToken(user)
	if err != nil {
		return dto.LoginResponse{}, fmt.Errorf("unexpected err: %w", err)
	}

	refreshToken, err := s.auth.CreateRefreshToken(user)
	if err != nil {
		return dto.LoginResponse{}, fmt.Errorf("unexpected err: %w", err)
	}

	// compare the user_handler password with the req.password

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		return dto.LoginResponse{}, fmt.Errorf("username or password is not correct")
	}

	return dto.LoginResponse{
		Tokens: dto.Tokens{
			AccessToken:  accessToken,
			RefreshToken: refreshToken,
		},
		User: dto.UserInfo{
			ID:          user.ID,
			Name:        user.Name,
			PhoneNumber: user.PhoneNumber,
		},
	}, nil
}
