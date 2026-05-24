package userservice

import (
	"game-app/param"
	"game-app/pkg/richerror"
)

func (s Service) GetProfile(req param.ProfileRequest) (param.ProfileResponse, error) {
	const op = "userservice.GetProfile"

	// getUserByID
	user, err := s.repo.GetUserByID(req.UserID)
	if err != nil {
		return param.ProfileResponse{}, richerror.New(op).WithErr(err).WithMeta(map[string]interface{}{"req": req})
	}

	return param.ProfileResponse{Name: user.Name}, nil
}
