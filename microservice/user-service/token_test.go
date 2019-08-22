package main

import (
	"testing"
	pb "github.com/gologic/microservice/user-service/proto/user"
	"github.com/dgrijalva/jwt-go"
	"time"
)

func TestToken(t *testing.T) {

	claims := CustomClaims{
		&pb.User{Name: "张三"},
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute*30).Unix(),
			Issuer:    "go.micro.srv.user",
		},
	}

	// Create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign token and return
	signToken, err := token.SignedString(key)

	if err != nil {
		t.Fatal(err)
	}
	t.Logf("token %s\n", signToken)


	tokenType, err := jwt.ParseWithClaims(signToken, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return key, nil
	})

	if err != nil {
		t.Logf("parse error %v\n", err)
	}

	// Validate the token and return the custom claims
	if claims, ok := tokenType.Claims.(*CustomClaims); ok && tokenType.Valid {
		t.Logf("valid %v \n", claims.User.Name)
	} else {
		t.Logf("%t %t\n", ok, tokenType.Valid)
	}
}