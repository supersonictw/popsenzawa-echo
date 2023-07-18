// PopSenzawa Echo
// (c) 2023 SuperSonic (https://github.com/supersonictw).

package pop

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func MiddlewareCheckRecaptcha(c *gin.Context) {
	if err := validateRecaptchaFromContext(c); err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.Next()
}
