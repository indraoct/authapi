package user_test

import (
	"authapi/api/user"
	"testing"
)

func TestRegisterUser(t *testing.T) {
	t.Run("positive case register user", func(t *testing.T) {
		h := user.NewUserHandler()
		req := &user.RegisterUserRequest{
			UserName:        "Lim",
			Password:        "test123",
			ConfirmPassword: "test123",
		}
		_, err := h.RegisterUser(req)
		if err != nil {
			t.Errorf("err should be nil")
			return
		}
	})
}
