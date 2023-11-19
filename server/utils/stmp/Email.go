package stmp

import (
	"context"
	"errors"
	"fmt"
	"github.com/dashixiong47/KK_BBS/config"
	"github.com/dashixiong47/KK_BBS/db"
	"github.com/dashixiong47/KK_BBS/utils/klog"
	"log"
	"math/rand"
	"net/smtp"
	"time"
)

// EmailSender 包含发送邮件所需的配置
type EmailSender struct {
	SMTPHost     string
	SMTPPort     string
	SenderEmail  string
	SenderPasswd string
}

// NewEmailSender 创建一个新的 EmailSender 实例
func NewEmailSender(host, port, email, passwd string) *EmailSender {
	return &EmailSender{
		SMTPHost:     host,
		SMTPPort:     port,
		SenderEmail:  email,
		SenderPasswd: passwd,
	}
}

// SendEmail 发送电子邮件到指定的接收者
func (es *EmailSender) SendEmail(to []string, subject, body string) error {
	auth := smtp.PlainAuth("", es.SenderEmail, es.SenderPasswd, es.SMTPHost)

	header := "From: " + es.SenderEmail + "\r\n" +
		"To: " + to[0] + "\r\n" +
		"Subject: " + subject + "\r\n\r\n"

	message := []byte(header + body)

	err := smtp.SendMail(es.SMTPHost+":"+es.SMTPPort, auth, es.SenderEmail, to, message)
	if err != nil {
		return err
	}

	return nil
}

// GenerateSecureVerificationCode 生成一个加密安全的随机验证码
func GenerateSecureVerificationCode(email string, length int) (string, error) {
	const digits = "0123456789"
	code := make([]byte, length)
	for i := range code {
		var b [1]byte
		_, err := rand.Read(b[:])
		if err != nil {
			return "", err
		}
		code[i] = digits[int(b[0])%len(digits)]
	}

	// 存到 Redis 中，使用用户的邮箱作为键
	err := SetVerificationCode(email, string(code))
	if err != nil {
		klog.Error("SetVerificationCode", err)
		return "", err
	}
	return string(code), nil
}

// SetVerificationCode 将验证码存到 Redis 中，键为用户邮箱
func SetVerificationCode(email, code string) error {
	ctx := context.Background()
	key := fmt.Sprintf("verifCode:%s", email)              // 使用格式化字符串构造键名
	err := db.Rdb.Set(ctx, key, code, 5*time.Minute).Err() // 设置有效期为5分钟
	if err != nil {
		return err
	}
	return nil
}

// SendVerificationCode 发送包含验证码的电子邮件
func (es *EmailSender) SendVerificationCode(to, username string) error {
	code, err := GenerateSecureVerificationCode(to, 6) // 生成6位数的验证码，使用邮箱作为键
	if err != nil {
		return err
	}
	subject := "您的注册验证码"
	body := fmt.Sprintf("尊敬的 %s，\n\n您的注册验证码是：%s\n\n请在网站上输入此验证码以完成注册。\n\n祝好，\n团队", username, code)

	return es.SendEmail([]string{to}, subject, body)
}

func SendVerificationEmail(to, username string) error {
	log.Println("Verification email test------------------------")
	// 创建 EmailSender 实例
	emailSender := NewEmailSender("us2.smtp.mailhostbox.com", "587", "admin@hxsj.in", "fmzqUps2")
	//emailSender := NewEmailSender(config.SettingsConfig.Smtp.Host, config.SettingsConfig.Smtp.Port, config.SettingsConfig.Smtp.User, config.SettingsConfig.Smtp.Password)
	log.Println(to, username, config.SettingsConfig.Smtp)
	// 发送验证码邮件
	err := emailSender.SendVerificationCode(to, username)
	if err != nil {
		klog.Error("Failed to send verification email:", err)
		return errors.New("failed_to_send_verification_email")
	}

	fmt.Println("Verification email sent successfully------------------------")
	return nil
}
