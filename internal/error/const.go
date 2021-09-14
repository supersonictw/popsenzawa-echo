// PopCat Echo
// (c) 2021 SuperSonic (https://github.com/supersonictw).

package error

type Code string

const (
	InvalidCount       Code = "invalid count"
	InvalidCountRange  Code = "invalid count range"
	AddressRateLimited Code = "address rate limited"
	EmptyCaptchaToken  Code = "empty captcha token"
	UnsafeCaptchaToken Code = "unsafe captcha token"
	UnknownRegionCode  Code = "unknown region code"
	UnknownJWTError    Code = "unknown jwt error"
)
