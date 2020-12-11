package request

import (
	"github.com/dgrijalva/jwt-go"
	uuid "github.com/satori/go.uuid"
)

type CustomClaims struct {
	UUID       uuid.UUID
	ID         uint
	Username   string
	Nickname   string
	BufferTime int64
	jwt.StandardClaims
}
