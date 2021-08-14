package pop

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/supersonictw/popcat-echo/internal"
	"github.com/supersonictw/popcat-echo/internal/config"
	"net/http"
	"strconv"
)

func Response(c *gin.Context) {
	ctx := context.Background()
	token := c.Query("token")
	ipAddress := c.ClientIP()
	if token == "" {
		newToken, err := IssueJWT(c, ctx)
		if err == nil {
			c.JSON(http.StatusOK, gin.H{
				"new_token": newToken,
			})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
			})
		}
		return
	}
	if status, err := ValidateJWT(c, token); !status {
		var message string
		if raised, ok := err.(*jwt.ValidationError); ok {
			message = raised.Error()
		} else if err != nil {
			message = err.Error()
		} else {
			message = internal.ErrorUnknownJWTError
		}
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": message,
		})
		return
	}
	captchaToken := c.Query("captcha_token")
	if !ValidateCaptcha(ipAddress, captchaToken) {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "unsafe captcha token",
		})
		return
	}
	count, err := strconv.Atoi(c.Query("count"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid count",
		})
		return
	}
	regionCode, err := GetRegionCode(ctx, ipAddress)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	pop := NewPop(count, ipAddress, regionCode)
	stepTimestamp := getCurrentStepTimestamp()
	key := fmt.Sprintf("%s:%d", config.CacheNamespacePop, stepTimestamp)
	err = redisClient.LPush(ctx, key, pop.JSON()).Err()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
	}
	newToken, err := IssueJWT(c, ctx)
	if err == nil {
		c.JSON(http.StatusCreated, gin.H{
			"new_token": newToken,
		})
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
	}
}
