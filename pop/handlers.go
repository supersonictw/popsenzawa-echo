// PopSenzawa Echo
// (c) 2023 SuperSonic (https://github.com/supersonictw).

package pop

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/supersonictw/popsenzawa-echo/data"
)

func PostPops(c *gin.Context) {
	pop := c.MustGet("pop").(*data.VisitorPop)

	if err := pop.Publish(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
	}

	newToken, err := issueJwtFromContext(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
	}

	c.JSON(http.StatusAccepted, &Response{
		NewToken: newToken,
	})
}
