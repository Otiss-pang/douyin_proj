package model

import (
	"github.com/golang-jwt/jwt"
	"log"
	"time"
)

// 连接数据库

func Connect2sql() {

}

// Encodetoken 此函数用于做jwt编码
func Encodetoken(userid string, username string) string {
	keyinfo := []byte("fzuirpangyifei+0.618+otiss")
	temp := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userid":   userid,
		"username": username,
		// exp: jwt的过期时间，这个过期时间必须要大于签发时间
		"exp": time.Now().Unix() + 3600*24,
		// iss: jwt签发者
		"iss": "daniel",
		// nbf: 定义在什么时间之前，该jwt都是不可用的.
		"nbf": time.Now().Unix(),
		// sub: jwt所面向的用户
		// aud: 接收jwt的一方
		// iat: jwt的签发时间
		// jti: jwt的唯一身份标识，主要用来作为一次性token,从而回避重放攻击。
	})
	token, err := temp.SignedString(keyinfo)
	if err != nil {
		log.Println(err.Error())
		return ""
	}
	return token
}

// Decodetoken 此函数用于做jwt解码，返回解码后得到的用户id与用户名
func Decodetoken(token string) []string {
	keyinfo := []byte("fzuirpangyifei+0.618+otiss")
	parse, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		return keyinfo, nil
	})
	if err != nil {
		log.Println(err.Error())
		return []string{"", ""}
	}
	return []string{parse.Claims.(jwt.MapClaims)["userid"].(string), parse.Claims.(jwt.MapClaims)["username"].(string)}

}
