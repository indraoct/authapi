package user

import (
	"database/sql"
	"log"
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
	tx, err := sqlRepo.db.Begin()
	if err != nil{
		log.Printf(err.Error())
		return err
	}
	stmt, err := tx.Prepare("INSERT INTO `user`(user_name, password) values (?, ?)")
	if err != nil{
		tx.Rollback()
		log.Printf(err.Error())
		return err
	}

	_,err = stmt.Exec(u.UserName,u.Password)
	if err != nil{
		tx.Rollback()
		log.Printf(err.Error())
		return err
	}

	err = tx.Commit()
	if err != nil{
		tx.Rollback()
		log.Printf(err.Error())
		return err
	}

	return nil
}
