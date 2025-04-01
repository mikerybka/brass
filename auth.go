package brass

import (
	"fmt"
	"net/http"

	"github.com/mikerybka/util"
	"golang.org/x/crypto/bcrypt"
)

type Auth struct {
	Users map[string]User
}

func (auth *Auth) GetUserID(r *http.Request) string {
	userID := r.Header.Get("UserID")
	user, ok := auth.Users[userID]
	if !ok {
		return "public"
	}

	sessionID := r.Header.Get("Token")
	if !user.Sessions[sessionID] {
		return "public"
	}

	return userID
}

func (a *Auth) Join(username, password, confirmPassword string) (string, error) {
	if _, ok := a.Users[username]; ok {
		return "", fmt.Errorf("username taken")
	}
	if password != confirmPassword {
		return "", fmt.Errorf("passwords don't match")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	sessionToken := util.RandomToken(16)
	user := User{
		ID:           username,
		PasswordHash: string(hashedPassword),
		Sessions: map[string]bool{
			sessionToken: true,
		},
	}

	if a.Users == nil {
		a.Users = map[string]User{}
	}
	a.Users[username] = user

	return sessionToken, nil
}
