package user_test

import (
	"authapi/api/user"
	"authapi/config"
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
	userDomain "authapi/domain/user"
	_ "github.com/go-sql-driver/mysql"
)

func TestRegisterUser(t *testing.T) {
	t.Run("positive case register user", func(t *testing.T) {
		cfg := config.Get()
		db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:3306)/%s", cfg.DBUser, cfg.DBPass, cfg.DBHost, cfg.DBName))
		if err != nil {
			panic(err)
		}
		userRepo := userDomain.NewRepository(db)
		userSvc := userDomain.NewService(userRepo)
		h := user.NewUserHandler(userSvc)

		body := map[string]interface{}{
			"userName" : "robot"+strconv.Itoa(rand.Int()),
			"password" : "123456",
			"confirmPassword" : "123456",
		}

		bodyJson, _ := json.Marshal(body)
		r, _ := http.NewRequest("POST", "/register", bytes.NewReader(bodyJson))
		r.Header.Set("Content-Type", "application/json")

		w := httptest.NewRecorder()
		h.RegisterUser(w, r)

		log.Println(w.Code)

		if w.Code != 200 {
			t.Error("Header response must be 200!")
		}


	})
}
