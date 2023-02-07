package service

import (
	"errors"
	"fmt"

	//"fmt"
	"douyin_proj/biz/common"
	"douyin_proj/biz/model/user"
)

func MatchUser(username, password string) error {
	userinfo := user.Users{}
	err := common.DB.First(&userinfo, "username = ? and password = ?", username, password).Error
	fmt.Println(userinfo.Password)
	return err

}

func AddUser(users user.Users) error {
	tempUser := user.Users{}
	common.DB.Find(&tempUser, "username = ?", users.Username)
	if tempUser != (user.Users{}) {
		return errors.New("账号重复")
	}
	err := common.DB.Create(&users).Error
	if err != nil {
		return err
	}
	return nil
}
