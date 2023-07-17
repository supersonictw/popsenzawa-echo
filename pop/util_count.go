// PopSenzawa Echo
// (c) 2023 SuperSonic (https://github.com/supersonictw).

package pop

import (
	"errors"
	"strconv"

	"github.com/gin-gonic/gin"
)

var (
	ErrCountEmpty   = errors.New("count empty")
	ErrCountInvalid = errors.New("count invalid")
)

func ValidateRangeFromContext(c *gin.Context) (int64, error) {
	countString := c.Query("count")
	if countString == "" {
		return 0, ErrCountEmpty
	}

	countInt64, err := strconv.ParseInt(countString, 10, 0)
	if err != nil {
		return 0, err
	}

	if countInt64 > 800 {
		return 0, ErrCountInvalid
	}

	return countInt64, nil
}
