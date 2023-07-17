// PopSenzawa Echo
// (c) 2023 SuperSonic (https://github.com/supersonictw).

package pop

import (
	"errors"

	"github.com/dpapathanasiou/go-recaptcha"
	"github.com/spf13/viper"
)

var (
	ErrRecaptchaTokenEmpty  = errors.New("recaptcha token empty")
	ErrRecaptchaTokenUnsafe = errors.New("recaptcha token unsafe")
)

var (
	recaptchaSecret = viper.GetString("pop_rules.recaptcha_secret")
)

func init() {
	recaptcha.Init(recaptchaSecret)
}

func ValidateRecaptcha(ipAddress, token string) error {
	if recaptchaSecret == "" {
		return nil
	}

	if token == "" {
		return ErrRecaptchaTokenEmpty
	}

	result, err := recaptcha.Confirm(ipAddress, token)
	if err != nil {
		return err
	}

	if !result {
		return ErrRecaptchaTokenUnsafe
	}

	return err
}
