package main

import (
	"fmt"

	"github.com/dgrijalva/jwt-go"
)

func main() {
	// This token is expired
	var tokenString = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJmb28iOiJiYXIiLCJleHAiOjE1MDAwLCJpc3MiOiJ0ZXN0In0.HE7fK0xOQwFEr4WDgRWj4teRPZ6i3GLwD5YCm6Pwu_c"

	token, _, err := new(jwt.Parser).ParseUnverified(tokenString, jwt.MapClaims{})
	if err != nil {
		fmt.Println(err)
		return
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		fmt.Println(claims["foo"], claims["exp"])
	} else {
		fmt.Println(err)
	}
}
