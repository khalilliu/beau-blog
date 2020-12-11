package response

import (
	"beau-blog/model"
)

type SysUserResponse struct {
	User model.SysUser `json:"user"`
}

type LoginResponse struct {
	User      model.SysUser `json:"user"`
	Toke      string        `json:"toke"`
	ExpiredAt int64         `json:"expired_at"`
}
