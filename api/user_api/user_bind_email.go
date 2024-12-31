package user_api

import (
	"fmt"
	"regexp"
	"time"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"go_blog/global"
	"go_blog/model"
	"go_blog/model/res"
	"go_blog/plugins/email"
	"go_blog/utils/jwts"
	"go_blog/utils/pwd"
	"go_blog/utils/random"
)

type BindEmailRequest struct {
	Email    string  `json:"email" binding:"required" msg:"邮箱非法"`
	Code     *string `json:"code"`
	Password string  `json:"password"`
}

func (UserAPI) UserBindEmailView(c *gin.Context) {
	_claims, _ := c.Get("claims")
	claims := _claims.(*jwts.CustomClaims)

	var cr BindEmailRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithError(err, &cr, c)
		return
	}

	session := sessions.Default(c)

	if cr.Code == nil {
		// 第一次请求：发送验证码

		// 校验邮箱格式
		emailRegex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
		matched, _ := regexp.MatchString(emailRegex, cr.Email)
		if !matched {
			res.FailWithMessage("邮箱格式不正确", c)
			return
		}

		// 限制验证码发送频率
		lastSent := session.Get("email_last_sent")
		if lastSent != nil {
			if time.Since(lastSent.(time.Time)) < time.Minute {
				res.FailWithMessage("验证码发送过于频繁，请稍后再试", c)
				return
			}
		}

		// 生成验证码并存入 session
		code := random.RandomString(4)
		session.Set("email_code", code)
		session.Set("email_target", cr.Email)
		session.Set("email_last_sent", time.Now())
		err := session.Save()
		if err != nil {
			global.Log.Error("Session 保存失败:", err)
			res.FailWithMessage("验证码发送失败，请稍后重试", c)
			return
		}

		// 发送验证码
		err = email.NewCode().Send(cr.Email, "你的验证码是: "+code)
		if err != nil {
			global.Log.Error("邮件发送失败:", err)
			res.FailWithMessage("验证码发送失败，请稍后重试", c)
			return
		}

		res.OkWithMessage("验证码已发送，请在5分钟内完成验证", c)
		return
	}

	// 第二次请求：验证验证码并绑定邮箱
	sessionCode := session.Get("email_code")
	sessionEmail := session.Get("email_target")

	fmt.Println("sessionCode:", sessionCode)
	fmt.Println("sessionEmail:", sessionEmail)

	if sessionCode == nil || sessionEmail == nil {
		res.FailWithMessage("验证码不存在或已过期", c)
		return
	}

	// 验证码和邮箱校验
	if sessionCode.(string) != *cr.Code || sessionEmail.(string) != cr.Email {
		res.FailWithMessage("验证码错误或邮箱不一致", c)
		return
	}

	// 验证码校验通过后，清除 session 中的验证码
	session.Delete("email_code")
	session.Delete("email_target")
	err = session.Save()
	if err != nil {
		global.Log.Error("Session 清除失败:", err)
	}

	// 查询当前用户信息
	var user model.UserModel
	err = global.DB.Take(&user, claims.UserID).Error
	if err != nil {
		res.FailWithMessage("用户不存在", c)
		return
	}

	// 校验密码复杂性
	passwordRegex := `^(?=.*[a-z])(?=.*[A-Z])(?=.*\d)(?=.*[@$!%*?&])[A-Za-z\d@$!%*?&]{8,}$`
	matched, _ := regexp.MatchString(passwordRegex, cr.Password)
	if !matched {
		res.FailWithMessage("密码必须包含大小写字母、数字和特殊字符，且长度不少于8位", c)
		return
	}

	// 加密密码
	hashPwd := pwd.HashPwd(cr.Password)

	// 使用事务更新用户信息
	tx := global.DB.Begin()
	err = tx.Model(&user).Updates(map[string]any{
		"email":    cr.Email,
		"password": hashPwd,
	}).Error
	if err != nil {
		tx.Rollback()
		global.Log.Error("更新用户信息失败:", err)
		res.FailWithMessage("邮箱绑定失败", c)
		return
	}
	tx.Commit()

	res.OkWithMessage("邮箱绑定成功", c)
}
