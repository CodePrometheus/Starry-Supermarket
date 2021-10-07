package models

import (
	"github.com/dgrijalva/jwt-go"
)

type JwtClaims struct {
	Id          uint
	Nickname    string
	AuthorityId uint
	jwt.StandardClaims
}
