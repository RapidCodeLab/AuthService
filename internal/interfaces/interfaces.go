package interfaces

import "context"

type JWTokener interface {
	NewJWT(User) ([]byte, error)
	RefreshJWT(context.Context, RefreshToken) ([]byte, error)
	GetPublicKey() []byte
	GetRefreshToken(context.Context, RefreshToken) (User, error)
	SetRefreshToken(context.Context, RefreshToken, User) error
	DeleteRefreshToken(context.Context, RefreshToken) error
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
	ID       string
	Email    string
	Password string
	Roles    []UserRole
	Status   UserStatus
}

type RefreshToken struct {
	RefreshToken []byte `json:"refresh_token"`
}

type UserService interface {
	GetUser(ctx context.Context, email, password string) (User, error)
	CreateUser(ctx context.Context, email, password string, role int) (User, error)
}

type Configurator interface {
	GetHTTPServerListenNetwork() string
	GetHTTPServerListenAddr() string
	GetGRPCServerListenNetwork() string
	GetGRPCServerListenAddr() string
	GetGRPCUserServiceAddr() string
}

type Logger interface{}

type EventProducer interface {
	Send(msg []byte)
}
