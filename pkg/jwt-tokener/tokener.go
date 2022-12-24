package jwttokener

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"encoding/json"
	"time"

	"github.com/RapidCodeLab/AuthService/internal/interfaces"
	"github.com/cristalhq/jwt/v4"
	"github.com/google/uuid"
)

type RefreshTokenStorage interface {
	Set(token, user []byte) error
	Get(token []byte) (user []byte, err error)
	Delete(token []byte) error
}

type JWTUserClaims struct {
	jwt.RegisteredClaims
	UserID int64                 `json:"user_id"`
	Email  string                `json:"email"`
	Roles  []interfaces.UserRole `json:"roles"`
}
type tokener struct {
	jwtTokenBuilder *jwt.Builder
	jwtSigner       jwt.Signer
	PublicKey       []byte
	rtStorage       RefreshTokenStorage
}

func New() (t *tokener, err error) {

	t = &tokener{}
	err = t.tokenBuilderUpdate()
	if err != nil {
		return
	}

	go func() {
		ticker := time.NewTicker(3600 * time.Second)
		for range ticker.C {
			err = t.tokenBuilderUpdate()
			if err != nil {
				//log error
			}
		}
	}()

	return
}

func (t *tokener) GetPublicKey() []byte {
	return t.PublicKey
}

func (t *tokener) NewJWT(u interfaces.User) (r []byte, err error) {
	userToken, err := buildUserJWTToken(u, t.jwtTokenBuilder)
	if err != nil {
		return
	}

	refreshToken, err := buildRefreshJWTToken(t.jwtTokenBuilder)
	if err != nil {
		return
	}
	r, err = json.Marshal(map[string]string{
		"token":         userToken.String(),
		"refresh_token": refreshToken.String(),
	})
	if err != nil {
		return
	}

	return
}

func (t *tokener) RefreshJWT(rt interfaces.RefreshToken) (r []byte, err error) {

	return
}

func (t *tokener) GetRefreshToken(
	rt interfaces.RefreshToken,
) (u interfaces.User, err error) {
	data, err := t.rtStorage.Get(rt.RefreshToken)
	if err != nil {
		return
	}

	err = json.Unmarshal(data, &u)

	return
}

func (t *tokener) SetRefreshToken(
	rt interfaces.RefreshToken,
	u interfaces.User) (err error) {

	data, err := json.Marshal(u)
	if err != nil {
		return
	}

	return t.rtStorage.Set(rt.RefreshToken, data)
}

func (t *tokener) DeleteRefreshToken(
	rt interfaces.RefreshToken) (err error) {

	return t.rtStorage.Delete(rt.RefreshToken)
}

func (t *tokener) tokenBuilderUpdate() (err error) {

	privateKey, err :=
		ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		return
	}

	publicKey := privateKey.PublicKey
	signer, err := jwt.NewSignerES(jwt.ES256, privateKey)
	if err != nil {
		return
	}
	t.jwtSigner = signer
	t.jwtTokenBuilder = jwt.NewBuilder(signer)
	t.PublicKey = elliptic.Marshal(&publicKey, publicKey.X, publicKey.Y)
	return
}

func buildUserJWTToken(u interfaces.User, jwtBuilder *jwt.Builder) (
	token *jwt.Token, err error) {
	claims := JWTUserClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			ID:       uuid.NewString(),
			Audience: []string{"users"},
			ExpiresAt: jwt.NewNumericDate(
				time.Now().Add(time.Minute * 15)),
		},
		UserID: u.ID,
		Email:  u.Email,
		Roles:  u.Roles,
	}

	return jwtBuilder.Build(claims)
}

func buildRefreshJWTToken(jwtBuilder *jwt.Builder) (
	token *jwt.Token, err error) {
	claims := jwt.RegisteredClaims{
		Subject:   "1",
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
	}
	return jwtBuilder.Build(claims)
}
