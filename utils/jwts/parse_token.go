package jwts

import (
	"github.com/golang-jwt/jwt/v5"
	"go_blog/global"
)

// ParseToken 解析token
func ParseToken(tokenString string) (*CustomClaims, error) {
	var MySecret = []byte(global.Config.JWT.Secret)
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return MySecret, nil
	})
	if err != nil {
		global.Log.Error("解析token错误", err)
		return nil, err
	}
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, err
	}
}
