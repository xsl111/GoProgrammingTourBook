package middleware

import (
	"GoprogrammingTourBook/blog-service/global"
	"GoprogrammingTourBook/blog-service/pkg/app"
	"GoprogrammingTourBook/blog-service/pkg/email"
	"GoprogrammingTourBook/blog-service/pkg/errcode"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)


func Recovery() gin.HandlerFunc {
	defailtMailer := email.NewEmail(&email.SMTPInfo{
		Host:     global.EmailSetting.Host,
		Port:     global.EmailSetting.Port,
		IsSSL:    global.EmailSetting.IsSSL,
		UserName: global.EmailSetting.UserName,
		Password: global.EmailSetting.Password,
		From:     global.EmailSetting.From,
	})
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				s := "panic recover err: %v"
				global.Logger.WithCallersFrames().Errorf(s, err)

				err := defailtMailer.SendMail(
					global.EmailSetting.To,
					fmt.Sprintf("异常抛出,发生时间: %d", time.Now().Unix()),
					fmt.Sprintf("错误信息: %v", err),
				)
				if err != nil {
					global.Logger.Panicf(c, "mail.SendMail err: %v", err)
				}

				app.NewResponse(c).ToErrorResponse(errcode.ServerError)
				c.Abort()
			}
		}()
		c.Next()
	}
}
