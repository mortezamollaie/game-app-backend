package mysql

import (
	"game-app/entity"
	"game-app/pkg/errmsg"
	"game-app/pkg/richerror"
	"game-app/pkg/slice"
	"strings"
	"time"
)

func (d *MySQLDB) GetUserPermissionsTitle(userID uint) ([]entity.PermissionTitle, error) {
	const op = "mysql.GetUserPermissions"
	user, err := d.GetUserByID(userID)
	if err != nil {
		return nil, richerror.New(op).WithErr(err).
			WithMessage(errmsg.ErrorMsgCantQuery).WithKind(richerror.KindUnexpected)
	}

	roleACL := make([]entity.AccessControl, 0)

	rows, err := d.db.Query("SELECT * FROM access_controls WHERE actor_type = ? and actor_id = ?", entity.RoleActorType, user.Role)
	if err != nil {
		return nil, richerror.New(op).WithErr(err).
			WithMessage(errmsg.ErrorMsgSomeThingWentWrong).WithKind(richerror.KindUnexpected)
	}

	defer rows.Close()

	for rows.Next() {
		acl, err := scanAccessControl(rows)
		if err != nil {
			return nil, richerror.New(op).WithErr(err).
				WithMessage(errmsg.ErrorMsgSomeThingWentWrong).WithKind(richerror.KindUnexpected)
		}

		roleACL = append(roleACL, acl)
	}

	if err := rows.Err(); err != nil {
		return nil, richerror.New(op).WithErr(err).
			WithMessage(errmsg.ErrorMsgSomeThingWentWrong).WithKind(richerror.KindUnexpected)
	}

	userACL := make([]entity.AccessControl, 0)

	userRows, err := d.db.Query("SELECT * FROM access_controls WHERE actor_type = ? and actor_id = ?", entity.UserActorType, user.ID)
	if err != nil {
		return nil, richerror.New(op).WithErr(err).
			WithMessage(errmsg.ErrorMsgSomeThingWentWrong).WithKind(richerror.KindUnexpected)
	}

	defer userRows.Close()

	for userRows.Next() {
		acl, err := scanAccessControl(userRows)
		if err != nil {
			return nil, richerror.New(op).WithErr(err).
				WithMessage(errmsg.ErrorMsgSomeThingWentWrong).WithKind(richerror.KindUnexpected)
		}

		userACL = append(userACL, acl)
	}

	if err := userRows.Err(); err != nil {
		return nil, richerror.New(op).WithErr(err).
			WithMessage(errmsg.ErrorMsgSomeThingWentWrong).WithKind(richerror.KindUnexpected)
	}

	// merge ACLs by permission id
	permissionIDs := make([]uint, 0)
	for _, r := range roleACL {
		if slice.DoesExist(permissionIDs, r.PermissionID) {
			permissionIDs = append(permissionIDs, r.PermissionID)
		}
	}

	if len(permissionIDs) == 0 {
		return nil, nil
	}

	// select * from permissions where ID in permissionIDs
	args := make([]interface{}, len(permissionIDs))

	for i, id := range permissionIDs {
		args[i] = id
	}

	// warning: this query works if we have one or more permission id
	pRows, err := d.db.Query("select * from access_controls where id in (?"+strings.Repeat(",?", len(permissionIDs)-1)+")", args...)
	if err != nil {
		return nil, richerror.New(op).WithErr(err).
			WithMessage(errmsg.ErrorMsgSomeThingWentWrong).WithKind(richerror.KindUnexpected)
	}

	defer pRows.Close()

	permissionTitles := make([]entity.PermissionTitle, 0)

	for pRows.Next() {
		permission, err := scanPermission(pRows)

		if err != nil {
			return nil, richerror.New(op).WithErr(err).
				WithMessage(errmsg.ErrorMsgSomeThingWentWrong).WithKind(richerror.KindUnexpected)
		}

		permissionTitles = append(permissionTitles, permission.Title)
	}

	if err := pRows.Err(); err != nil {
		return nil, richerror.New(op).WithErr(err).
			WithMessage(errmsg.ErrorMsgSomeThingWentWrong).WithKind(richerror.KindUnexpected)
	}

	return permissionTitles, nil

}

func scanAccessControl(scanner Scanner) (entity.AccessControl, error) {
	var createdAt time.Time
	var acl entity.AccessControl

	err := scanner.Scan(&acl.ID, &acl.ActorID, &acl.ActorType, &acl.PermissionID, &createdAt)
	return acl, err
}
