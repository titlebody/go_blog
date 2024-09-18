package jwts

import (
	"github.com/golang-jwt/jwt/v5"
)

type JwtPayLoad struct {
	UserID   uint   `json:"user_id"`   //用户ID
	Role     int    `json:"role"`      //角色  1.管理员 2.普通用户 3.游客
	UserName string `json:"user_name"` //用户名
	NickName string `json:"nick_name"` //名称
}

type CustomClaims struct {
	JwtPayLoad
	jwt.RegisteredClaims
}
