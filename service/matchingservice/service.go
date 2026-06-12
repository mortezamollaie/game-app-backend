package matchingservice

import (
	"game-app/entity"
	"game-app/param"
	"game-app/pkg/richerror"
)

type Repo interface {
	AddToWaitingList(userID uint, category entity.Category) error
}

type Service struct {
	repo Repo
}

func New() Service {
	return Service{}
}

func (s Service) AddToWaitingList(req param.AddToWaitingListRequest) (param.AddToWaitingListResponse, error) {
	const op = richerror.Op("matchingservice.AddToWaitingList")

	// add user to the waiting list for the given category if not exist
	err := s.repo.AddToWaitingList(req.UserID, req.Category)
	if err != nil {
		return param.AddToWaitingListResponse{}, richerror.New(op).WithErr(err).WithKind(richerror.KindUnexpected)
	}

	// also we can update the waiting timestamp

}
