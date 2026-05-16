package mysql

import (
	"database/sql"
	"game-app/entity"
	"game-app/pkg/errmsg"
	"game-app/pkg/richerror"
)

func (d *MySQLDB) GetUserACL(userID uint) ([]entity.AccessControl, error) {
	const op = "mysql.GetUserACL"
	user, err := d.GetUserByID(userID)
	if err != nil {
		return nil, richerror.New(op).WithErr(err).
			WithMessage(errmsg.ErrorMsgCantQuery).WithKind(richerror.KindUnexpected)
	}

	rows, err := d.db.Query("SELECT * FROM access_controls WHERE actor_type = 'role' and actor_id = ?", user.Role)
	if err != nil {
		return nil, richerror.New(op).WithErr(err).
			WithMessage(errmsg.ErrorMsgSomeThingWentWrong).WithKind(richerror.KindUnexpected)
	}

	for _, row := range rows {

	}

	user, err := scanUser(row)
	if err != nil {
		if err == sql.ErrNoRows {
			return entity.User{}, richerror.New(op).WithErr(err).
				WithMessage(errmsg.ErrorMsgNotFound).WithKind(richerror.KindNotFound)
		}

		return entity.User{}, richerror.New(op).WithErr(err).
			WithMessage(errmsg.ErrorMsgCantQuery).WithKind(richerror.KindUnexpected)
	}
}
