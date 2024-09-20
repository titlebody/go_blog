package main

import (
	"fmt"
	"go_blog/core"
	"go_blog/global"
	"gopkg.in/gomail.v2"
)

type Subject string

const (
	Code  Subject = "平台验证码"
	Note  Subject = "操作通知"
	Alarm Subject = "告警通知"
)

type Api struct {
	Subject Subject
}

func (a Api) Send(name, body string) error {
	return send(name, string(a.Subject), body)
}

func NewCode() Api {
	return Api{
		Subject: Code,
	}
}

func NewNote() Api {
	return Api{
		Subject: Note,
	}
}

func NewAlarm() Api {
	return Api{
		Subject: Alarm,
	}
}

// send 发送邮件 发给谁 邮件主题 内容
func send(name, subject, body string) error {

	e := global.Config.Email
	fmt.Println(e)
	return sendMail(
		e.User,
		e.Password,
		e.Host,
		e.Port,
		name,
		e.DefaultFromEmail,
		subject,
		body,
	)
}

// 文本邮件
func sendMail(userName, authCode, host string, port int, mailTo, sendName string, subject, body string) error {
	m := gomail.NewMessage()
	m.SetHeader("From", m.FormatAddress(userName, sendName))
	m.SetHeader("To", mailTo)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", body)
	d := gomail.NewDialer(host, port, userName, authCode)
	return d.DialAndSend(m)

}

func main() {
	core.InitConfig()
	core.InitLogger()
	NewCode().Send("2387360024@qq.com", "验证码是：2056")
}
