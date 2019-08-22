package main

import (
	"testing"
	pb "github.com/gologic/microservice/user-service/proto/user"
	"github.com/dgrijalva/jwt-go"
	"time"
)


func TestDecodeToken(t *testing.T) {
	signToken := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VyIjp7ImlkIjoiMDQ3MGI4YzMtZmFkYy00YWVlLWI0ZDktMjdkOTc4NDFkYWRjIiwibmFtZSI6IkV3YW4gVmFsZW50aW5lIiwiY29tcGFueSI6IkJCQyIsImVtYWlsIjoiZXdhbi52YWxlbnRpbmU4OUBnbWFpbC5jb20iLCJwYXNzd29yZCI6IiQyYSQxMCQ2ay5lS1VrZDFmWTRjMkl2b1lHQm9lcWJGc2NrRXZHVXNYTW5VeEpKUHZpT3g3TkQyV1RodSJ9LCJleHAiOjE1NjY0NDc0OTksImlzcyI6ImdvLm1pY3JvLnNydi51c2VyIn0.NvlTHz5kx1EQbEdQsNHbbR3OZhmc8P_J_aqfqnwntiU"
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