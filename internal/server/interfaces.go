package server

import (
	"github.com/cristalhq/jwt/v3"
)

type JWTokener interface {
	NewJWT(u UserDTO) ([]byte, error)
	UpdateRT(rt RT) ([]byte, error)
}

type UserRole int

type UserDTO struct{}

type JWTUserClaims struct {
	jwt.RegisteredClaims
	UserID int64      `json:"user_id"`
	Email  string     `json:"email"`
	Roles  []UserRole `json:"roles"`
}

type RT struct {
	RefreshToken []byte `json:"refresh_token"`
}
