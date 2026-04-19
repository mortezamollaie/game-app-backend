package mysql

import (
	"database/sql"
	"errors"
	"game-app/entity"
	"game-app/pkg/richerror"
	"time"
)

func (d MySQLDB) IsPhoneNumberUnique(phoneNumber string) (bool, error) {
	row := d.db.QueryRow(`select * from users where phone_number = ?`, phoneNumber)
	_, err := scanUser(row)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return true, nil
		}

		return false, richerror.New(err, "mysql.IsPhoneNumberUnique", "can't scan query result", richerror.KindUnexpected, nil)
	}

	return false, err
}

func (d MySQLDB) Register(u entity.User) (entity.User, error) {
	res, err := d.db.Exec(`insert into users(name, phone_number, password) values (?, ?, ?)`, u.Name, u.PhoneNumber, u.Password)
	if err != nil {
		return entity.User{}, richerror.New(err, "mysql.Register", "can't scan query result", richerror.KindUnexpected, nil)
	}

	id, _ := res.LastInsertId()
	u.ID = uint(id)
	return u, nil
}

func (d MySQLDB) GetUserByPhoneNumber(phoneNumber string) (entity.User, bool, error) {
	row := d.db.QueryRow(`select * from users where phone_number = ?`, phoneNumber)
	user, err := scanUser(row)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return entity.User{}, false, nil
		}

		return entity.User{}, false, richerror.New(err, "mysql.GetUserByPhoneNumber", "can't scan query result", richerror.KindUnexpected, nil)
	}

	return user, true, nil
}

func (d MySQLDB) GetUserByID(id uint) (entity.User, error) {
	row := d.db.QueryRow(`select * from users where id = ?`, id)
	user, err := scanUser(row)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return entity.User{}, nil
		}

		return entity.User{}, richerror.New(err, "mysql.GetUserByID", "record not found", richerror.KindNotFound, nil)
	}

	return user, nil
}

func scanUser(row *sql.Row) (entity.User, error) {
	var createdAt time.Time
	var user entity.User
	err := row.Scan(&user.ID, &user.Name, &user.PhoneNumber, &createdAt, &user.Password)
	return user, err
}
