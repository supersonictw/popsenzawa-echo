package pop

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/supersonictw/popcat-echo/internal"
	echoErrors "github.com/supersonictw/popcat-echo/internal/error"
	"net/http"
	"strconv"
)

func Response(c *gin.Context) {
	ctx := context.Background()
	token := c.Param("token")
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
			if raised.Errors&jwt.ValidationErrorMalformed != 0 {
				message = echoErrors.JWTMalformed
			} else if raised.Errors&jwt.ValidationErrorUnverifiable != 0 {
				message = echoErrors.JWTSignatureUnacceptable
			} else if raised.Errors&jwt.ValidationErrorSignatureInvalid != 0 {
				message = echoErrors.JWTSignatureInvalid
			} else if raised.Errors&jwt.ValidationErrorExpired != 0 {
				message = echoErrors.JWTExpired
			} else if raised.Errors&jwt.ValidationErrorNotValidYet != 0 {
				message = echoErrors.JWTNotValidYet
			} else {
				message = echoErrors.JWTInvalid
			}
		}
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": message,
		})
		return
	}
	captchaToken := c.Param("captcha_token")
	if !ValidateCaptcha(c.ClientIP(), captchaToken) {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "unsafe captcha token",
		})
		return
	}
	count, err := strconv.Atoi(c.Param("count"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid count",
		})
		return
	}
	pop := NewPop(count, "", "")
	internal.RDB.LPush(ctx, "", pop.JSON())
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
}
