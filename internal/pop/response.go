package pop

import (
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func Response(c *gin.Context) {
	count, err := strconv.Atoi(c.Param("count"))
	if err != nil {
		panic(err)
	}
	timestamp := time.Now().Unix()
	stepTimestamp := timestamp / 100 * 100
}
