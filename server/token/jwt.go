package token

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// WebToken contain timeout show time before timeout
// secret is secret for sign token
type WebToken struct {
	addtime time.Duration
	secret  string
}

// NewWebToken is constructor
func NewWebToken(addtime time.Duration, secret string) *WebToken {
	return &WebToken{addtime: addtime, secret: secret}
}

// CreateToken create a jwt token
// TODO: interface for input
func (w *WebToken) CreateToken(username string) (string, error) {
	var err error

	atClaims := jwt.MapClaims{}

	atClaims["authorized"] = true
	atClaims["username"] = username
	atClaims["exp"] = time.Now().Add(w.addtime).Unix()

	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err := at.SignedString([]byte(w.secret))

	if err != nil {
		return "", err
	}
	return token, nil
}

// VerifyToken verify jwt token
func (w *WebToken) VerifyToken(r *http.Request) (*jwt.Token, error) {
	tokenString := ExtractToken(r)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		//Make sure that the token method conform to "SigningMethodHMAC"
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(w.secret), nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}

// ExtractToken extract from http Header
func ExtractToken(r *http.Request) string {
	bearToken := r.Header.Get("Authorization")
	//normally Authorization the_token_xxx
	strArr := strings.Split(bearToken, " ")
	if len(strArr) == 2 {
		return strArr[1]
	}
	return ""
}
