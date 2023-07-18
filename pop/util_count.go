// PopSenzawa Echo
// (c) 2023 SuperSonic (https://github.com/supersonictw).

package pop

import (
	"errors"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

var (
	ErrCountEmpty   = errors.New("count empty")
	ErrCountInvalid = errors.New("count invalid")
)

var (
	maxPopsAppendPerVisitor = viper.GetInt64("count.max_pops_append_per_visitor")
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

	if countInt64 < 0 {
		return 0, ErrCountInvalid
	}
	if countInt64 > maxPopsAppendPerVisitor {
		countInt64 = maxPopsAppendPerVisitor
	}

	return countInt64, nil
}
