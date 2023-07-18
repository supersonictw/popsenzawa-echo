// PopSenzawa Echo
// (c) 2023 SuperSonic (https://github.com/supersonictw).

package pop

import (
	"errors"

	"github.com/dpapathanasiou/go-recaptcha"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

var (
	ErrRecaptchaTokenEmpty  = errors.New("recaptcha token empty")
	ErrRecaptchaTokenUnsafe = errors.New("recaptcha token unsafe")
)

var (
	configrecaptchaSecret = viper.GetString("recaptcha.secret")
)

func init() {
	if configrecaptchaSecret == "" {
		return
	}
	recaptcha.Init(configrecaptchaSecret)
}

func validateRecaptcha(ipAddress, token string) error {
	if configrecaptchaSecret == "" {
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

func validateRecaptchaFromContext(c *gin.Context) error {
	captchaToken := c.Query("captcha_token")
	ipAddress := c.ClientIP()

	return validateRecaptcha(captchaToken, ipAddress)
}
