package interfaces

type JWTokener interface {
	NewJWT(u User) ([]byte, error)
	UpdateRT(rt RT) ([]byte, error)
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

type RT struct {
	RefreshToken []byte `json:"refresh_token"`
}

type UserService interface {
	GetUser(email, password string) error
}

type Logger interface{}
