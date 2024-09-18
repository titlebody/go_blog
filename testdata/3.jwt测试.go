package main

import (
	"fmt"
	"go_blog/core"
	"go_blog/global"
	"go_blog/utils/jwts"
)

func main() {
	core.InitConfig()
	global.Log = core.InitLogger()
	if global.Config.JWT.Secret == "" {
		fmt.Println("请先配置jwt的secret")
		return
	}
	token, err := jwts.CreateToken(jwts.JwtPayLoad{NickName: "xxx", Role: 1, UserID: 1, UserName: "chenxi"})
	if err != nil {
		fmt.Println("创建token错误", err)
		return
	}
	fmt.Println(token)
	clasime, er := jwts.ParseToken("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxLCJyb2xlIjoxLCJ1c2VyX25hbWUiOiJjaVueGkiLCJuaWNrX25hbWUiOiJ4eHgiLCJpc3MiOiIxMjMiLCJleHAiOjE3MjY4MjMzNjl9.JGoAsUzluw738qtqQqp8nCyyGVpz7akJOMdNhH9zaos")
	if er != nil {
		fmt.Println("解析token错误", er)
		return
	}
	fmt.Println(clasime)

}
