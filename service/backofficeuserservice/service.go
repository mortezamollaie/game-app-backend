package backofficeuserservice

import "game-app/entity"

type Service struct{}

func New() Service {
	return Service{}
}

func (s Service) ListAllUsers() ([]entity.User, error) {
	// TODO - implement me

	list := make([]entity.User, 0)

	list = append(list, entity.User{
		ID:          1,
		PhoneNumber: "09900580684",
		Name:        "fake",
		Password:    "123456",
		Role:        entity.SuperAdminRole,
	})

	return list, nil
}
