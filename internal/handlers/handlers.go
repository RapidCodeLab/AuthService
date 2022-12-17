package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/RapidCodeLab/AuthService/internal/interfaces"
)

type LoginUserDTO struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Login(
	w http.ResponseWriter,
	r *http.Request,
	jwTokener interfaces.JWTokener) {

	var loginUserDTO LoginUserDTO
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	err := decoder.Decode(&loginUserDTO)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		//reason to body
		return
	}

	//get user from UserService gRPC
	var user interfaces.User

	token, err := jwTokener.NewJWT(user)
	if err != nil {
		w.WriteHeader(http.StatusBadGateway)
		//reason to body
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Header().Add("Content-Type", "application/json")
	w.Write(token)

}

func Signup(w http.ResponseWriter, r *http.Request)       {}
func RefreshToken(w http.ResponseWriter, r *http.Request) {}
func Logout(w http.ResponseWriter, r *http.Request)       {}
