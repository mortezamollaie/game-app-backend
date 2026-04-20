package mysql

import (
	"database/sql"
	"errors"
	"game-app/entity"
	"game-app/pkg/errmsg"
	"game-app/pkg/richerror"
	"time"
)

func (d MySQLDB) IsPhoneNumberUnique(phoneNumber string) (bool, error) {
	const op = "mysql.IsPhoneNumberUnique"
	row := d.db.QueryRow(`select * from users where phone_number = ?`, phoneNumber)
	_, err := scanUser(row)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return true, nil
		}

		return false, richerror.New(op).WithErr(err).WithMessage(errmsg.ErrorMsgCantQuery).WithKind(richerror.KindUnexpected)
	}

	return false, err
}

func (d MySQLDB) Register(u entity.User) (entity.User, error) {
	const op = "mysql.Register"

	res, err := d.db.Exec(`insert into users(name, phone_number, password) values (?, ?, ?)`, u.Name, u.PhoneNumber, u.Password)
	if err != nil {
		return entity.User{}, richerror.New(op).WithErr(err).WithMessage(errmsg.ErrorMsgCantInsert).WithKind(richerror.KindUnexpected)
	}

	id, _ := res.LastInsertId()
	u.ID = uint(id)
	return u, nil
}

func (d MySQLDB) GetUserByPhoneNumber(phoneNumber string) (entity.User, bool, error) {
	const op = "mysql.GetUserByPhoneNumber"
	row := d.db.QueryRow(`select * from users where phone_number = ?`, phoneNumber)
	user, err := scanUser(row)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return entity.User{}, false, nil
		}

		return entity.User{}, false, richerror.New(op).WithErr(err).WithMessage(errmsg.ErrorMsgCantQuery).WithKind(richerror.KindUnexpected)
	}

	return user, true, nil
}

func (d MySQLDB) GetUserByID(id uint) (entity.User, error) {
	const op = "mysql.GetUserByID"
	row := d.db.QueryRow(`select * from users where id = ?`, id)
	user, err := scanUser(row)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return entity.User{}, nil
		}

		return entity.User{}, richerror.New(op).WithErr(err).WithMessage(errmsg.ErrorMsgNotFound).WithKind(richerror.KindNotFound)
	}

	return user, nil
}

func scanUser(row *sql.Row) (entity.User, error) {
	var createdAt time.Time
	var user entity.User
	err := row.Scan(&user.ID, &user.Name, &user.PhoneNumber, &createdAt, &user.Password)
	return user, err
}
