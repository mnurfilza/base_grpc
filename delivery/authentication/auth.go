package authentication

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type MyClaims struct {
	jwt.RegisteredClaims
	Username string `json:"Username"`
}

var APPLICATION_NAME = "e-money-svc"
var LOGIN_EXPIRATION_DURATION = time.Duration(1) * time.Hour
var JWT_SIGNING_METHOD = jwt.SigningMethodHS256
var JWT_SIGNATURE_KEY = []byte("the secret of kalimdor")

func CreateToken(username string) (string, error) {
	claims := MyClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    APPLICATION_NAME,
			ExpiresAt: &jwt.NumericDate{time.Now().Add(LOGIN_EXPIRATION_DURATION)},
		},
		Username: username,
	}

	token := jwt.NewWithClaims(JWT_SIGNING_METHOD, claims)

	signedToken, err := token.SignedString(JWT_SIGNATURE_KEY)

	if err != nil {
		return "", err
	}

	return signedToken, nil

}
