package jwttokener

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"encoding/json"

	"github.com/RapidCodeLab/AuthService/internal/interfaces"
	"github.com/cristalhq/jwt"
	"github.com/google/uuid"
)

type JWTUserClaims struct {
	jwt.StandardClaims
	UserID int64                 `json:"user_id"`
	Email  string                `json:"email"`
	Roles  []interfaces.UserRole `json:"roles"`
}
type tokener struct {
	jwtBuilder *jwt.TokenBuilder
}

func New() *tokener {
	return &tokener{}
}

func (t *tokener) NewJWT(u interfaces.User) (r []byte, err error) {
	userToken, err := buildUserJWTToken(u, t.jwtBuilder)
	if err != nil {
		return
	}
	//refretsh token

	r, err = json.Marshal(map[string]string{
		"token": userToken.String(),
		//"refresh_token": rtToken.String(),
	})
	if err != nil {
		return nil, err
	}

	return
}
func (t *tokener) UpdateRT(rt interfaces.RT) (r []byte, err error) {
	return
}

func (t *tokener) tokenBuilderUpdate() (err error) {

	privateKey, err :=
		ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		return
	}

	signer, err := jwt.NewES256(&privateKey.PublicKey, privateKey)
	if err != nil {
		return
	}
	t.jwtBuilder = jwt.NewTokenBuilder(signer)
	return
}

func buildUserJWTToken(u interfaces.User, jwtBuilder *jwt.TokenBuilder) (
	token *jwt.Token, err error) {
	claims := JWTUserClaims{
		StandardClaims: jwt.StandardClaims{
			ID:       uuid.NewString(),
			Audience: []string{"users"},
			//ExpiresAt: jwt.,
		},
		UserID: u.ID,
		Email:  u.Email,
		Roles:  u.Roles,
	}

	token, err = jwtBuilder.Build(claims)
	if err != nil {
		return
	}

}
