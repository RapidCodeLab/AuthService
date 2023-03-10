package handlers_test

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	jwttokener "github.com/RapidCodeLab/AuthService/pkg/jwt-tokener"
	mockconfigurator "github.com/RapidCodeLab/AuthService/pkg/mocks/configurator"
	mockskvtorage "github.com/RapidCodeLab/AuthService/pkg/mocks/kv-storage"
	"github.com/RapidCodeLab/AuthService/pkg/mocks/logger"
	mockuserservice "github.com/RapidCodeLab/AuthService/pkg/mocks/user-service"

	"github.com/RapidCodeLab/AuthService/internal/handlers"
	"github.com/stretchr/testify/assert"
)

func TestSignin(t *testing.T) {

	assert := assert.New(t)

	ctx := context.Background()
	c := mockconfigurator.New()
	kv := mockskvtorage.New()

	jwtTokener, err := jwttokener.New(kv)
	assert.Nil(err)

	us, err := mockuserservice.New(ctx, c)
	assert.Nil(err)

	userDTO := handlers.SigninUserDTO{
		Email:    "test@test.com",
		Password: "qwerty",
	}
	_, err = us.CreateUser(ctx, userDTO.Email, userDTO.Password, 0)
	assert.Nil(err)

	payload, err := json.Marshal(userDTO)
	assert.Nil(err)
	body := bytes.NewReader(payload)

	req, err := http.NewRequest(http.MethodPost, handlers.SigninPath, body)
	res := httptest.NewRecorder()

	l := &logger.Logger{}
	handler := http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			handlers.Signin(l, w, r, jwtTokener, us)
		})

	handler.ServeHTTP(res, req)

	assert.Equal(res.Code, http.StatusCreated)
}

func TestSignup(t *testing.T) {
	assert := assert.New(t)

	ctx := context.Background()
	c := mockconfigurator.New()

	us, err := mockuserservice.New(ctx, c)
	assert.Nil(err)

	userDTO := handlers.SignupUserDTO{
		Email:    "test@test.com",
		Password: "qwerty",
		Role:     0,
	}

	payload, err := json.Marshal(userDTO)
	assert.Nil(err)

	body := bytes.NewReader(payload)

	req, err := http.NewRequest(http.MethodPost, handlers.SignupPath, body)
	res := httptest.NewRecorder()

	l := &logger.Logger{}
	handler := http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			handlers.Signup(l, w, r, us)
		})

	handler.ServeHTTP(res, req)

	assert.Equal(res.Code, http.StatusCreated)

}
