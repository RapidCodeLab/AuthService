package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/RapidCodeLab/AuthService/internal/interfaces"
)

type ErrorReason struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

const (
	SigninPath       = "/auth/signin"
	SignupPath       = "/auth/signup"
	RefreshTokenPath = "/auth/refresh"
	SignoutPath      = "/auth/signout"
)

type SigninUserDTO struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type SignupUserDTO struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     int    `json:"role"`
}

func Signin(
	w http.ResponseWriter,
	r *http.Request,
	jwTokener interfaces.JWTokener,
	userService interfaces.UserService) {

	var (
		signinUserDTO SigninUserDTO
	)

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	err := decoder.Decode(&signinUserDTO)
	if err != nil {
		err = writeErrorReason(w, http.StatusBadRequest, err.Error())
		if err != nil {
			//log err
		}
		return
	}

	//get user from UserService gRPC
	u, err := userService.GetUser(r.Context(),
		signinUserDTO.Email,
		signinUserDTO.Password)
	if err != nil {
		err = writeErrorReason(w, http.StatusBadGateway, err.Error())
		if err != nil {
			//log err
		}
		return
	}

	token, err := jwTokener.NewJWT(u)
	if err != nil {
		err = writeErrorReason(w, http.StatusBadGateway, err.Error())
		if err != nil {
			//log err
		}
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Header().Add("Content-Type", "application/json")
	_, err = w.Write(token)
	if err != nil {
		err = writeErrorReason(w, http.StatusBadGateway, err.Error())
		if err != nil {
			//log err
		}
	}
}

func Signup(
	w http.ResponseWriter,
	r *http.Request,
	userService interfaces.UserService) {

	var (
		signupUserDTO SignupUserDTO
	)
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	err := decoder.Decode(&signupUserDTO)
	if err != nil {
		err = writeErrorReason(w, http.StatusBadGateway, err.Error())
		if err != nil {
			//log err
		}
		return
	}

	//create user
	u, err := userService.CreateUser(r.Context(),
		signupUserDTO.Email,
		signupUserDTO.Password,
		signupUserDTO.Role)
	if err != nil {
		err = writeErrorReason(w, http.StatusBadGateway, err.Error())
		if err != nil {
			//log err
		}
		return
	}

	res := struct {
		Email string `json:"email"`
	}{
		Email: u.Email,
	}

	w.WriteHeader(http.StatusCreated)
	w.Header().Add("Content-Type", "application/json")
	enc := json.NewEncoder(w)
	err = enc.Encode(res)
	if err != nil {
		err = writeErrorReason(w, http.StatusBadGateway, err.Error())
		if err != nil {
			//log err
		}
	}

}

func RefreshToken(w http.ResponseWriter,
	r *http.Request,
	jwtTokener interfaces.JWTokener) {

	var (
		rt interfaces.RefreshToken
	)

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	err := decoder.Decode(&rt)
	if err != nil {
		err = writeErrorReason(w, http.StatusBadGateway, err.Error())
		if err != nil {
			//log err
		}
		return
	}

	token, err := jwtTokener.RefreshJWT(r.Context(), rt)
	if err != nil {
		err = writeErrorReason(w, http.StatusUnauthorized, err.Error())
		if err != nil {
			//log err
		}
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Header().Add("Content-Type", "application/json")
	_, err = w.Write(token)
	if err != nil {
		err = writeErrorReason(w, http.StatusBadGateway, err.Error())
		if err != nil {
			//log err
		}
	}
}

func Logout(w http.ResponseWriter, r *http.Request) {}

func writeErrorReason(w http.ResponseWriter,
	status int,
	message string) error {
	var errorReason ErrorReason

	w.WriteHeader(status)
	errorReason.Code = status
	errorReason.Message = message
	body, err := json.Marshal(errorReason)
	if err != nil {
		return err
	}
	_, err = w.Write(body)
	if err != nil {
		return err
	}
	return nil
}
