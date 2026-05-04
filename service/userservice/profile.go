package userservice

import (
	"game-app/dto"
	"game-app/pkg/richerror"
)

func (s Service) GetProfile(req dto.ProfileRequest) (dto.ProfileResponse, error) {
	const op = "userservice.GetProfile"

	// getUserByID
	user, err := s.repo.GetUserByID(req.UserID)
	if err != nil {
		return dto.ProfileResponse{}, richerror.New(op).WithErr(err).WithMeta(map[string]interface{}{"req": req})
	}

	return dto.ProfileResponse{Name: user.Name}, nil
}
