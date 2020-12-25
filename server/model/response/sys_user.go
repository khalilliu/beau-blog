package response

import (
	"beau-blog/model"
)

type UserResponse struct {
	User model.User `json:"user"`
}

type LoginResponse struct {
	User      model.User `json:"user"`
	Token      string     `json:"toke"`
	ExpiredAt int64      `json:"expired_at"`
}
