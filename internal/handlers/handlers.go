package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/RapidCodeLab/AuthService/internal/interfaces"
)

type SigninUserDTO struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type SignupUserDTO struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Signin(
	w http.ResponseWriter,
	r *http.Request,
	jwTokener interfaces.JWTokener,
	userService interfaces.UserService) {

	var signinUserDTO SigninUserDTO
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	err := decoder.Decode(&signinUserDTO)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		//reason to body
		return
	}

	//get user from UserService gRPC
	user, err := userService.GetUser(r.Context(),
		signinUserDTO.Email,
		signinUserDTO.Password)
	if err != nil {
		w.WriteHeader(http.StatusBadGateway)
		//reason to body
		return
	}

	token, err := jwTokener.NewJWT(user)
	if err != nil {
		w.WriteHeader(http.StatusBadGateway)
		//reason to body
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Header().Add("Content-Type", "application/json")
	_, err = w.Write(token)
	if err != nil {
		w.WriteHeader(http.StatusBadGateway)
		//reason to body
	}
}

func Signup(w http.ResponseWriter, r *http.Request) {

	var signupUserDTO SignupUserDTO
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	err := decoder.Decode(&signupUserDTO)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		//reason to body
		return
	}

	//create user
	var u interfaces.User

	res := struct {
		Email string
	}{
		Email: u.Email,
	}

	w.WriteHeader(http.StatusCreated)
	w.Header().Add("Content-Type", "application/json")
	enc := json.NewEncoder(w)
	err = enc.Encode(res)
	if err != nil {
		w.WriteHeader(http.StatusBadGateway)
		//reason to body
		return
	}

}

func RefreshToken(w http.ResponseWriter, r *http.Request) {}
func Logout(w http.ResponseWriter, r *http.Request)       {}
