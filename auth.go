package brass

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	"path/filepath"

	"github.com/mikerybka/util"
	"golang.org/x/crypto/bcrypt"
)

type Auth struct {
	DataDir string
}

func (auth *Auth) user(id string) (*User, bool, error) {
	path := filepath.Join(auth.DataDir, "users", id)
	b, err := os.ReadFile(path)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return nil, false, nil
		} else {
			return nil, false, err
		}
	}
	user := &User{}
	err = json.Unmarshal(b, user)
	if err != nil {
		panic(err)
	}
	return user, true, nil
}

func (auth *Auth) GetUserID(r *http.Request) (string, error) {
	userID := util.Cookie(r, "UserID")
	user, ok, err := auth.user(userID)
	if !ok {
		return "", err
	}
	sessionID := util.Cookie(r, "SessionID")
	if !user.Sessions[sessionID] {
		return "", nil
	}
	return userID, nil
}

func (a *Auth) Join(username, password, confirmPassword string) (string, error) {
	if _, ok := UsernameBlocklist[username]; ok {
		return "", fmt.Errorf("username blocked")
	}
	_, ok, err := a.user(username)
	if err != nil {
		return "", err
	}
	if ok {
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
	user := &User{
		ID:           username,
		PasswordHash: string(hashedPassword),
		Sessions: map[string]bool{
			sessionToken: true,
		},
	}
	err = a.saveUser(username, user)
	if err != nil {
		return "", err
	}
	err = util.Touch(filepath.Join(a.DataDir, "orgs", user.ID, "members", user.ID))
	if err != nil {
		return "", err
	}
	return sessionToken, nil
}

func (a *Auth) saveUser(id string, u *User) error {
	path := filepath.Join(a.DataDir, "users", id)
	return util.WriteJSONFile(path, u)
}

func (a *Auth) Login(username, password string) (string, error) {
	user, ok, err := a.user(username)
	if err != nil {
		return "", err
	}
	if !ok {
		return "", fmt.Errorf("no user")
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))
	if err != nil {
		return "", err
	}
	token := util.RandomToken(16)
	if user.Sessions == nil {
		user.Sessions = map[string]bool{}
	}
	user.Sessions[token] = true
	err = a.saveUser(username, user)
	if err != nil {
		return "", err
	}
	return token, nil
}

func (a *Auth) Allowed(r *http.Request) (bool, error) {
	orgID := r.PathValue("orgID")
	if orgID == "" {
		return true, nil
	}
	userID, err := a.GetUserID(r)
	if err != nil {
		return false, err
	}
	_, err = os.Stat(filepath.Join(a.DataDir, "orgs", orgID, "members", userID))
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return false, nil
		}
		return false, err
	}
	return true, nil
}
