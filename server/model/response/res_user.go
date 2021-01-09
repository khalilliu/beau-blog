package response

import (
	"beau-blog/model"
)

type ResUser struct {
	User model.User `json:"user"`
}

type ResLogin struct {
	User      model.User `json:"user"`
	Token     string     `json:"toke"`
	ExpiredAt int64      `json:"expired_at"`
}
