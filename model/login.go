package model

import "github.com/golang-jwt/jwt/v4"

type Login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Claim struct {
	Email string `json:"email"`
	jwt.StandardClaims
}
