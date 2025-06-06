package brass

import (
	"encoding/json"
	"net/http"
	"path/filepath"

	"github.com/mikerybka/constants"
)

func NewAPI(appHost string) *API {
	api := &API{
		appHost:     appHost,
		authManager: NewManager[*Auth](filepath.Join(constants.DataDir, appHost, "auth.json")),
		metaManager: NewManager[*Metadata](filepath.Join(constants.DataDir, appHost, "meta.json")),
	}
	return api
}

type API struct {
	appHost     string
	authManager *Manager[*Auth]
	metaManager *Manager[*Metadata]
}

func (api *API) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "UserID, Token")
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/auth/join", api.join)
	mux.HandleFunc("/auth/login", api.login)
	mux.HandleFunc("/auth/logout", api.logout)
	mux.HandleFunc("/auth/change-password", api.changePassword)
	mux.HandleFunc("/auth/delete-account", api.deleteAccount)
	mux.Handle("/meta", api.metaManager.Get())
	// mux.HandleFunc("/data/{owner}/", func(w http.ResponseWriter, r *http.Request) {
	// 	// Authenticate
	// 	userID := api.authManager.Get().GetUserID(r)

	// 	// Authorize
	// 	owner := r.PathValue("owner")
	// 	if owner != userID { // TODO: set up user groups
	// 		w.WriteHeader(http.StatusUnauthorized)
	// 		return
	// 	}

	// 	// Proxy request to database
	// 	_, id, _ := util.PopPath(r.URL.Path)
	// 	path := api.appHost + id
	// 	req, err := http.NewRequest(r.Method, fmt.Sprintf("http://localhost:4000/%s", path), r.Body)
	// 	if err != nil {
	// 		http.Error(w, err.Error(), http.StatusInternalServerError)
	// 		return
	// 	}
	// 	res, err := http.DefaultClient.Do(req)
	// 	if err != nil {
	// 		http.Error(w, err.Error(), http.StatusInternalServerError)
	// 		return
	// 	}
	// 	w.WriteHeader(res.StatusCode)
	// 	_, err = io.Copy(w, res.Body)
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// })
	mux.ServeHTTP(w, r)
}

func (a *API) join(w http.ResponseWriter, r *http.Request) {
	req := &struct {
		Username        string `json:"username"`
		Password        string `json:"password"`
		ConfirmPassword string `json:"confirmPassword"`
	}{}
	err := json.NewDecoder(r.Body).Decode(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	auth := a.authManager.Get()
	token, err := auth.Join(req.Username, req.Password, req.ConfirmPassword)
	a.authManager.Set(auth)

	res := &struct {
		Token string `json:"token"`
		Error error  `json:"error"`
	}{
		Token: token,
		Error: err,
	}
	json.NewEncoder(w).Encode(res)
}

func (a *API) login(w http.ResponseWriter, r *http.Request) {
	req := &struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}{}
	err := json.NewDecoder(r.Body).Decode(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	auth := a.authManager.Get()
	token, err := auth.Login(req.Username, req.Password)
	a.authManager.Set(auth)

	res := &struct {
		Token string `json:"token"`
		Error error  `json:"error"`
	}{
		Token: token,
		Error: err,
	}
	json.NewEncoder(w).Encode(res)
}

func (a *API) logout(w http.ResponseWriter, r *http.Request)         {}
func (a *API) changePassword(w http.ResponseWriter, r *http.Request) {}
func (a *API) deleteAccount(w http.ResponseWriter, r *http.Request)  {}
