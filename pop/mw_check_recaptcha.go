// PopSenzawa Echo
// (c) 2023 SuperSonic (https://github.com/supersonictw).

package pop

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func MiddlewareCheckRecaptcha(c *gin.Context) {
	captchaToken := c.Query("captcha_token")
	ipAddress := c.ClientIP()

	if err := ValidateRecaptcha(ipAddress, captchaToken); err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.Next()
}
