package jwts

import (
	"github.com/golang-jwt/jwt/v5"
	"go_blog/global"
	"time"
)

// CreateToken 创建token
func CreateToken(user JwtPayLoad) (string, error) {
	var MySecret = []byte(global.Config.JWT.Secret)
	claim := CustomClaims{
		user,
		jwt.RegisteredClaims{
			Issuer:    global.Config.JWT.Issuer,
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * time.Duration(global.Config.JWT.Expires))),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	return token.SignedString(MySecret)
}
