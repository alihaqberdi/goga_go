package jwt_manager

import (
	"encoding/json"
	"errors"
	"github.com/alihaqberdi/goga_go/internal/dtos"
	"github.com/dgrijalva/jwt-go"
	"strings"
	"time"
)

type JwtManager struct {
	signingKey     []byte
	expiryInterval time.Duration
}

func New(signingKey string, expiryInterval time.Duration) *JwtManager {
	return &JwtManager{signingKey: []byte(signingKey), expiryInterval: expiryInterval}
}

func (t *JwtManager) Generate(user dtos.JwtUser) (string, error) {
	bytes, err := json.Marshal(user)
	if err != nil {
		return "", err
	}

	subject := string(bytes)

	return t.generate(subject)
}

func (t *JwtManager) Parse(tokenStr string) (*dtos.JwtUser, error) {
	tokenStr = removeBearerIfExists(tokenStr)

	subject, err := t.parse(tokenStr)
	if err != nil {
		return nil, err
	}

	user := new(dtos.JwtUser)
	err = json.Unmarshal([]byte(subject), user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (t *JwtManager) generate(sub string) (string, error) {
	now := time.Now()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &jwt.StandardClaims{
		IssuedAt:  now.Unix(),
		ExpiresAt: now.Add(t.expiryInterval).Unix(),
		Subject:   sub,
	})
	return token.SignedString(t.signingKey)
}

func (t *JwtManager) parse(tokenStr string) (string, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}
		return t.signingKey, nil
	})
	if err != nil {
		return "", err
	}

	claims, ok := token.Claims.(*jwt.StandardClaims)
	if !ok {
		return "", errors.New("token claims are not type of  jwt.StandardClaims")
	}
	return claims.Subject, nil
}

func removeBearerIfExists(token string) string {
	arr := strings.Split(token, " ")
	if len(arr) < 2 {
		return token
	}
	if len(arr) == 2 && strings.EqualFold("Bearer", arr[0]) {
		return arr[1]
	}
	return token
}
