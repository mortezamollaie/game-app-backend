package authorizationservice

import "game-app/entity"

type Repository interface {
	GetUserACL(userID uint) ([]entity.AccessControl, error)
}

type Service struct{}

func (s Service) CheckAccess(userID uint, permission ...entity.PermissionTitle) (bool, error) {
	// get the user role

	// get all ACLs for the given role

	// get all ACLs for the given user

	// merge all ACLs

	// check the access

	return false, nil
}
