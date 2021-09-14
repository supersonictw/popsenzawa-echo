// PopCat Echo
// (c) 2021 SuperSonic (https://github.com/supersonictw).

package error

import (
	"errors"
)

func NewError(code Code) error {
	return errors.New(string(code))
}
