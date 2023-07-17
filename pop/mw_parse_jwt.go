// PopSenzawa Echo
// (c) 2023 SuperSonic (https://github.com/supersonictw).

package pop

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

func MiddlewareParseJwt(c *gin.Context) {
	if claims, err := validateJwtFromContext(c); errors.Is(err, ErrJwtEmpty) {
		newToken, err := issueJwtFromContext(c)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
			return
		}
		c.AbortWithStatusJSON(http.StatusOK, Response{
			NewToken: newToken,
		})
	} else if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
	} else {
		c.Set("claims", claims)
		c.Next()
	}
}
