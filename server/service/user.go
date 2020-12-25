package service

import (
	"beau-blog/global"
	"beau-blog/model"
	"beau-blog/utils"
	"errors"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

func Register(u model.User)(err error, userInter model.User){
	var user model.User
	if errors.Is(global.BB_DB.Where("username = ?", u.Username).First(&user).Error, gorm.ErrRecordNotFound) {
		return errors.New("用户名已注册"), user
	}
	u.Password = utils.MD5V([]byte(u.Password))
	u.UUID = uuid.NewV4()
	err = global.BB_DB.Create(&user).Error
	return  err, u
}

func Login(u *model.User) (err error, userInter model.User) {
	var user model.User
	u.Password = utils.MD5V([]byte(u.Password))
	err = global.BB_DB.Where("username = ? AND password = ?", u.Username, u.Password).First(&user).Error
	return  err, user
}

func SetUserInfo(u model.User) (err error, user model.User) {
	err = global.BB_DB.Updates(&u).Error
	return err, u
}

func ChangePassword(u *model.User, newPassword string) (err error, userInter *model.User) {
	var user model.User
	u.Password = utils.MD5V([]byte(u.Password))
	newPass := utils.MD5V([]byte(newPassword))
	err = global.BB_DB.Where("username = ? AND password = ?", u.Username, u.Password).First(&user).Update("password", newPass).Error
	return  err, u
}

func FindUserById(id int)(err error, user *model.User){
	var u model.User
	err = global.BB_DB.Where("`id` = ?", id).First(&u).Error
	return err, &u
}

func FindUserByUuid(uuid string) (err error, user *model.User) {
	var u model.User
	if err = global.BB_DB.Where("`uuid` = ?", uuid).First(&u).Error; err != nil {
		return errors.New("用户不存在"), &u
	}
	return nil, &u
}