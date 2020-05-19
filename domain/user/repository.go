package user

import (
	"database/sql"
)

type Repository interface {
	StoreUser(*User) error
}

type sqlRepository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &sqlRepository{db}
}

func (sqlRepo *sqlRepository) StoreUser(u *User) error {
	_, err := sqlRepo.db.Exec("INSERT INTO `user`(user_name, password) values (?, ?)", u.UserName, u.Password)
	return err
}
