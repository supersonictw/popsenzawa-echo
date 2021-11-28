// PopCat Echo
// (c) 2021 SuperSonic (https://github.com/supersonictw).

package error

import "errors"

type Code string

var (
	InvalidCount       = errors.New("invalid count")
	InvalidCountRange  = errors.New("invalid count range")
	AddressRateLimited = errors.New("address rate limited")
	EmptyCaptchaToken  = errors.New("empty captcha token")
	UnsafeCaptchaToken = errors.New("unsafe captcha token")
	UnknownRegionCode  = errors.New("unknown region code")
	UnknownJWTError    = errors.New("unknown jwt error")
)
