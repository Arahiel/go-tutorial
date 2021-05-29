package main

import (
	"errors"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var key = []byte("my secret key 435321")

type myClaims struct {
	jwt.StandardClaims
	SessionId string
}

func createToken(sid string) (string, error) {
	claims := myClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * 5).Unix(),
		},
		SessionId: sid,
	}

	t := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	ss, err := t.SignedString(key)
	if err != nil {
		return "", fmt.Errorf("couldn't sign token in createToken", err)
	}
	return ss, nil
}

func parseToken(ss string) (string, error) {
	token, err := jwt.ParseWithClaims(ss, &myClaims{}, func(t *jwt.Token) (interface{}, error) {
		if t.Method.Alg() != jwt.SigningMethodHS256.Alg() {
			return nil, errors.New("wrong signing algorithm")
		}
		return key, nil
	})

	if err != nil {
		return "", err
	}

	if claims, ok := token.Claims.(*myClaims); ok && token.Valid {
		return claims.SessionId, nil
	} else {
		return "", errors.New("claims or token is not valid")
	}
}
