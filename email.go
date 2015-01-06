package email

import (
	"github.com/xybstone/go_log"
	"net/smtp"
	"regexp"
	"strings"
)

var (
	Smtp_host     string
	Smtp_port     string
	Smtp_username string
	Smtp_nickname string
	Smtp_password string
)

//检验图纸有效性
func ValidEmail(in string) bool {
	matched, _ := regexp.MatchString(`^(\w)+(\.\w+)*@(\w)+((\.\w+)+)$`, in)
	return matched
}

//发送电子邮件
func SendEmail(to, subject, body string) {
	vaild := ValidEmail(to)
	if !vaild {
		log.Error("Invalid email " + to)
		return
	}

	smtp_address := Smtp_host + ":" + Smtp_port
	mail_list := strings.Split(to, ";")
	auth := smtp.PlainAuth("", Smtp_username, Smtp_password, Smtp_host)
	log.Info("smtp_address=%s Smtp_username=%s Smtp_password=%s",
		smtp_address, Smtp_username, Smtp_password)
	content_type := "Content-Type: text/html; charset=UTF-8"
	msg := []byte("To: " + to + "\r\nFrom: " + Smtp_nickname + "<" + Smtp_username +
		">\r\nSubject: " + subject + "\r\n" + content_type + "\r\n\r\n" + body)
	err := smtp.SendMail(smtp_address, auth, Smtp_username, mail_list, msg)
	if err != nil {
		log.Error("SendMail : %v", err)
	}
}
