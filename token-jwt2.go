package main

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	_ "github.com/someone1/gcp-jwt-go"
	//"time"
)

func main() {
	token := jwt.New(jwt.GetSigningMethod("AppEngine"))
	fmt.Println(token)
}
