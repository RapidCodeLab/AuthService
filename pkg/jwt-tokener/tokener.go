package jwttokener

import (
	"context"
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
	SetValueByKey(ctx context.Context, key string, user []byte) error
	GetValueByKey(ctx context.Context, key string) (user []byte, err error)
	DeleteByKey(ctx context.Context, key string) error
}

type JWTUserClaims struct {
	jwt.RegisteredClaims
	UserID string                `json:"user_id"`
	Email  string                `json:"email"`
	Roles  []interfaces.UserRole `json:"roles"`
}
type tokener struct {
	jwtTokenBuilder *jwt.Builder
	jwtSigner       jwt.Signer
	PublicKey       *ecdsa.PublicKey
	rtStorage       RefreshTokenStorage
}

func New(s RefreshTokenStorage) (t *tokener, err error) {

	t = &tokener{
		rtStorage: s,
	}

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
	return elliptic.Marshal(elliptic.P256(), t.PublicKey.X, t.PublicKey.Y)
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

	return
}

func (t *tokener) RefreshJWT(ctx context.Context,
	rt interfaces.RefreshToken) (r []byte, err error) {

	defer t.DeleteRefreshToken(ctx, rt)

	_, err = veryfyRefreshToken(rt.RefreshToken, t.PublicKey)
	if err != nil {
		return
	}

	user, err := t.rtStorage.GetValueByKey(ctx, string(rt.RefreshToken))
	if err != nil {
		return
	}
	var u interfaces.User
	err = json.Unmarshal(user, &u)
	if err != nil {
		return
	}

	return t.NewJWT(u)
}

func (t *tokener) GetRefreshToken(
	ctx context.Context,
	rt interfaces.RefreshToken,
) (u interfaces.User, err error) {
	data, err := t.rtStorage.GetValueByKey(ctx, string(rt.RefreshToken))
	if err != nil {
		return
	}

	err = json.Unmarshal(data, &u)

	return
}

func (t *tokener) SetRefreshToken(
	ctx context.Context,
	rt interfaces.RefreshToken,
	u interfaces.User) (err error) {

	data, err := json.Marshal(u)
	if err != nil {
		return
	}

	return t.rtStorage.SetValueByKey(ctx, string(rt.RefreshToken), data)
}

func (t *tokener) DeleteRefreshToken(
	ctx context.Context,
	rt interfaces.RefreshToken) (err error) {

	return t.rtStorage.DeleteByKey(ctx, string(rt.RefreshToken))
}

func (t *tokener) tokenBuilderUpdate() (err error) {

	privateKey, err :=
		ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		return
	}

	signer, err := jwt.NewSignerES(jwt.ES256, privateKey)
	if err != nil {
		return
	}
	t.jwtSigner = signer
	t.jwtTokenBuilder = jwt.NewBuilder(signer)
	t.PublicKey = &privateKey.PublicKey
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

func veryfyRefreshToken(token []byte,
	publicKey *ecdsa.PublicKey) (t *jwt.Token, err error) {

	alg, err := jwt.NewVerifierES(jwt.ES256, publicKey)
	if err != nil {
		return
	}
	return jwt.Parse(token, alg)
}
