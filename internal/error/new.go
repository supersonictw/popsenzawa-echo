package error

import (
	"errors"
)

func NewError(code Code) error {
	return errors.New(string(code))
}
