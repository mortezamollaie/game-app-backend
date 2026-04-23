package userservice

import (
	"fmt"
	"game-app/dto"
	"game-app/entity"
	"game-app/pkg/richerror"

	"golang.org/x/crypto/bcrypt"
)

type Repository interface {
	Register(u entity.User) (entity.User, error)
	GetUserByPhoneNumber(phoneNumber string) (entity.User, bool, error)
	GetUserByID(id uint) (entity.User, error)
}

type AuthGenerator interface {
	CreateAccessToken(user entity.User) (string, error)
	CreateRefreshToken(user entity.User) (string, error)
}

type Service struct {
	auth AuthGenerator
	repo Repository
}

func New(authGenerator AuthGenerator, repo Repository) Service {
	return Service{auth: authGenerator, repo: repo}
}

func (s Service) Register(req dto.RegisterRequest) (dto.RegisterResponse, error) {
	// TODO - we should verify phone number by verification code

	pass := []byte(req.Password)
	pass, err := bcrypt.GenerateFromPassword(pass, 0)
	if err != nil {
		return dto.RegisterResponse{}, err
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
		return dto.RegisterResponse{}, fmt.Errorf("unexpected err: %w", err)
	}

	return dto.RegisterResponse{User: struct {
		ID          uint   `json:"id"`
		Name        string `json:"name"`
		PhoneNumber string `json:"phone_number"`
	}{ID: createdUser.ID, PhoneNumber: createdUser.PhoneNumber, Name: createdUser.Name}}, nil
}

func (s Service) Login(req dto.LoginRequest) (dto.LoginResponse, error) {
	// TODO - it would be better to user two separate method for existence check and getUserByPhoneNumber
	const op = "userservice.Login"

	// check the existence of phone number from repository
	// get the user by phone number

	user, exits, err := s.repo.GetUserByPhoneNumber(req.PhoneNumber)
	if err != nil {
		return dto.LoginResponse{}, richerror.New(op).WithErr(err).WithMeta(map[string]interface{}{"phone_number": req.PhoneNumber})
	}

	if !exits {
		return dto.LoginResponse{}, fmt.Errorf("username or password is not correct")
	}

	accessToken, err := s.auth.CreateAccessToken(user)
	if err != nil {
		return dto.LoginResponse{}, fmt.Errorf("unexpected err: %w", err)
	}

	refreshToken, err := s.auth.CreateRefreshToken(user)
	if err != nil {
		return dto.LoginResponse{}, fmt.Errorf("unexpected err: %w", err)
	}

	// compare the user password with the req.password

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

func (s Service) GetProfile(req dto.ProfileRequest) (dto.ProfileResponse, error) {
	const op = "userservice.GetProfile"

	// getUserByID
	user, err := s.repo.GetUserByID(req.UserID)
	if err != nil {
		return dto.ProfileResponse{}, richerror.New(op).WithErr(err).WithMeta(map[string]interface{}{"req": req})
	}

	return dto.ProfileResponse{Name: user.Name}, nil
}
