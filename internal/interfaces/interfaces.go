package interfaces

import "context"

type JWTokener interface {
	NewJWT(u User) ([]byte, error)
	RefreshJWT(rt RefreshToken) ([]byte, error)
	GetPublicKey() []byte
}

type UserRole int

const (
	UserRoleSuper UserRole = iota
	UserRoleAdmin
	UserRoleRegular
)

type UserStatus int

const (
	UserStatusNew UserStatus = iota
	UserStatusActive
	UserStatusBanned
)

type User struct {
	ID       int64
	Email    string
	Password string
	Roles    []UserRole
	Status   UserStatus
}

type RefreshToken struct {
	RefreshToken []byte `json:"refresh_token"`
}

type RefreshTokenService interface {
	Set(RefreshToken) error
	Get(RefreshToken) (User, error)
	Delete(RefreshToken) error
}

type UserService interface {
	GetUser(ctx context.Context, email, password string) (User, error)
}

type Configurator interface {
	GetHTTPServerListenNetwork() string
	GetHTTPServerListenAddr() string
	GetGRPCServerListenNetwork() string
	GetGRPCServerListenAddr() string
	GetGRPCUserServiceAddr() string
}

type Logger interface{}
