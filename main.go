package main

import (
	"authapi/api/user"
	"authapi/config"
	"authapi/server"
	"fmt"
	"time"

	"database/sql"

	_ "github.com/go-sql-driver/mysql"

	userDomain "authapi/domain/user"
	"log"
)

func main() {
	cfg := config.Get()
	log.Printf("%#v", cfg)
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:3306)/%s", cfg.DBUser, cfg.DBPass, cfg.DBHost, cfg.DBName))
	if err != nil {
		panic(err)
	}

	//connection pooling  read : http://go-database-sql.org/connection-pool.html
	db.SetMaxIdleConns(0)
	db.SetMaxOpenConns(5)
	db.SetConnMaxLifetime(time.Minute * 5)

	// Init repo
	userRepo := userDomain.NewRepository(db)

	// Init services
	userSvc := userDomain.NewService(userRepo)

	// Init handler
	userHandler := user.NewUserHandler(userSvc)
	server.ServeHTTP(cfg.Port, userHandler)
}
