package middleware

import (
	"beau-blog/global"
	"beau-blog/model/request"
	"beau-blog/model/response"
	"beau-blog/service"
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"strconv"
	"time"
)

type JWT struct {
	SigningKey []byte
}

var (
	TokenExpired   = errors.New("Token is expired")
	TokenNotValid  = errors.New("Token not active yet")
	TokenMalformed = errors.New("That's not even a token")
	TokenInvalid   = errors.New("Couldn't handle this token:")
)

func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("x-token")
		if token == "" {
			response.FailWithDetailed(gin.H{"reload": true}, "未登录或非法访问", c )
			c.Abort()
			return
		}
		j := NewJWT()
		// parse token
		claims, err := j.ParseToken(token)
		if err != nil {
			if err == TokenExpired {
				response.FailWithDetailed(gin.H{"reload": true}, "授权已过期", c)
				c.Abort()
				return
			}
			response.FailWithDetailed(gin.H{"reload": true}, err.Error(), c)
			c.Abort()
			return
		}
		if err, _  = service.FindUserByUuid(claims.UUID.String()); err != nil {
			response.FailWithDetailed(gin.H{"reload": true}, err.Error(), c)
			c.Abort()
		}
		if claims.ExpiresAt - time.Now().Unix() < claims.BufferTime {
			claims.ExpiresAt = time.Now().Unix() + global.BB_CONFIG.JWT.ExpiresTime
			newToken, _ := j.CreateToken(*claims)
			newClaims, _ := j.ParseToken(newToken)
			c.Header("new-token", newToken)
			c.Header("new-expires-at", strconv.FormatInt(newClaims.ExpiresAt, 10))
			if global.BB_CONFIG.System.UseMultipoint {

			}
		}
		c.Set("claims", claims)
		c.Next()
	}
}

func NewJWT() *JWT {
	return &JWT{
		SigningKey: []byte(global.BB_CONFIG.JWT.SigningKey),
	}
}

// 创建一个token
func (j *JWT) CreateToken(claims request.ReqCustomClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(j.SigningKey)
}

// 解析token
func (j *JWT) ParseToken(tokenStr string) (*request.ReqCustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &request.ReqCustomClaims{}, func(token *jwt.Token) (i interface{}, e error) {
		return j.SigningKey, nil
	})
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, TokenMalformed
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				return nil, TokenExpired
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, TokenNotValid
			} else {
				return nil, TokenInvalid
			}
		}
	}
	if token != nil {
		if claim, ok := token.Claims.(*request.ReqCustomClaims); ok && token.Valid {
			return claim, nil
		}
		return nil, TokenInvalid
	} else {
		return nil, TokenInvalid
	}
}
