package handler

import (
	"context"
	"crypto/md5"
	"crypto/sha256"
	"douyin_proj/biz/common"
	"douyin_proj/biz/model"
	"douyin_proj/biz/model/user"
	"douyin_proj/biz/service"
	"encoding/hex"
	"fmt"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/json"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"log"
	"strconv"
)

type UserLoginResponse struct {
	common.Response
	UserId uint   `json:"user_id"`
	Token  string `json:"token"`
}

type UserResponse struct {
	common.Response
	UserInfo user.Users
}

// MD5加密
func MD5(str string) string {
	data := []byte(str) // 切片
	has := md5.Sum(data)
	return fmt.Sprintf("%x", has) // 将[]byte转成16进制
}

func Sha256(str string) string {
	m := sha256.New()
	m.Write([]byte(str))
	res := hex.EncodeToString(m.Sum(nil))
	return res
}

func UserLogin(c context.Context, ctx *app.RequestContext) {
	bytes, err := ctx.Body()
	if err != nil {
		panic(err)
	}

	var userinfo user.Users
	if err = json.Unmarshal(bytes, &userinfo); err != nil {
		panic(err)
	}
	// 加密
	userinfo.Password = Sha256(userinfo.Password)
	// 判断数据库是否存在该用户
	err = service.MatchUser(userinfo.Username, userinfo.Password)
	if err != nil {
		ctx.JSON(consts.StatusOK, common.Fail(40001))
		return
	}
	ctx.JSON(consts.StatusOK, UserLoginResponse{
		Response: common.Response{
			200,
			"登录成功",
			nil,
		},
		UserId: userinfo.ID,
		// token可以使用jwt中间件进行处理
		Token: model.Encodetoken(strconv.FormatUint(uint64(userinfo.ID), 10), userinfo.Username),
	})
}

func UserRegister(c context.Context, ctx *app.RequestContext) {
	bytes, err := ctx.Body()
	if err != nil {
		ctx.JSON(consts.StatusOK, common.FailWithMsg(40001, "参数错误"))
		panic(err)
	}
	var userinfo user.Users
	// 将json数据转化为结构体
	err = json.Unmarshal(bytes, &userinfo)
	if err != nil {
		ctx.JSON(consts.StatusOK, common.Fail(50000))
		panic(err)
	}
	// 密码加密
	userinfo.Password = Sha256(userinfo.Password)
	log.Println(userinfo)
	// 在数据库中添加用户
	if err = service.AddUser(userinfo); err != nil {
		ctx.JSON(consts.StatusOK, common.FailWithMsg(50000, err.Error()))
		panic(err)
	}
	ctx.JSON(consts.StatusOK, UserLoginResponse{
		Response: common.Response{
			200,
			"用户创建成功",
			nil,
		},
		UserId: userinfo.ID,
		Token:  model.Encodetoken(strconv.FormatUint(uint64(userinfo.ID), 10), userinfo.Username),
	})
}

// 根据id查询用户信息
func GetUser(c context.Context, ctx *app.RequestContext) {
	userid := ctx.Param("id")
	var userinfo user.Users
	log.Println(userid)
	if err := common.DB.Find(&userinfo, "id = ?", userid); err != nil {
		ctx.JSON(consts.StatusOK, common.FailWithMsg(50001, "该用户未注册"))
		//return
	}
	ctx.JSON(consts.StatusOK, common.Response{
		200, "查询成功", userinfo,
	})
}
