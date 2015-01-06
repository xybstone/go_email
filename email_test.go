package email

import (
	"testing"
)

func StartTest() {
	Smtp_host = "smtp.lcshentu.com"
	Smtp_port = "25"
	Smtp_username = "no-reply@lcshentu.com"
	Smtp_nickname = "励铖数字化审图系统邮件"
	Smtp_password = "licheng1201"
}

func TestSendEmail(t *testing.T) {
	StartTest()
	SendEmail("xybstone@gmail.com", "test", "testtest")
}
