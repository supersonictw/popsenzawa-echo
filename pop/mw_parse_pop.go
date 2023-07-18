// PopSenzawa Echo
// (c) 2023 SuperSonic (https://github.com/supersonictw).

package pop

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/supersonictw/popsenzawa-echo/data"
)

func MiddlewareParseCount(c *gin.Context) {
	claims := c.MustGet("claims").(*jwt.RegisteredClaims)

	ipAddress := data.ParseVisitorIP(claims.Audience[0])
	regionCode := claims.Subject

	count, err := validateRangeFromContext(c)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.Set("pop", &data.VisitorPop{
		IPAddress:  ipAddress,
		RegionCode: regionCode,
		Count:      count,
	})
}
