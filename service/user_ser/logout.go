package user_ser

import (
	"go_blog/service/redis"
	"go_blog/utils/jwts"
	"time"
)

func (UserService) Logout(claims *jwts.CustomClaims, token string) error {
	exp := claims.ExpiresAt
	now := time.Now()
	diff := exp.Time.Sub(now)
	return redis.Logout(token, diff)
}
