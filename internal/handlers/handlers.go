package handlers

import (
	"net/http"
)

func Login(w http.ResponseWriter, r *http.Request)        {}
func Signup(w http.ResponseWriter, r *http.Request)       {}
func RefreshToken(w http.ResponseWriter, r *http.Request) {}
func Logout(w http.ResponseWriter, r *http.Request)       {}
