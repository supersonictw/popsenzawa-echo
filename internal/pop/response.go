package pop

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/supersonictw/popcat-echo/internal"
	"strconv"
)

func Response(c *gin.Context) {
	token := c.Param("token")
	if token == "" {
		c.JSON(200, gin.H{
			"new_token": IssueJWT(c),
		})
	}
	if !ValidateJWT(token) {
		c.JSON(401, gin.H{
			"message": "invalid jwt token",
		})
		return
	}
	captchaToken := c.Param("captcha_token")
	if !ValidateCaptcha(captchaToken) {
		c.JSON(401, gin.H{
			"message": "unsafe captcha token",
		})
		return
	}
	count, err := strconv.Atoi(c.Param("count"))
	if err != nil {
		c.JSON(400, gin.H{
			"message": "invalid count",
		})
		return
	}
	pop := NewPop(count, "", "")
	ctx := context.Background()
	internal.RDB.LPush(ctx, "", pop.JSON())
	c.JSON(201, gin.H{
		"new_token": IssueJWT(c),
	})
}
