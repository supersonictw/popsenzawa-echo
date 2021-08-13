package error

const (
	UnknownRegionCode        = "unknown region code"
	JWTMalformed             = "token is malformed"
	JWTSignatureUnacceptable = "token could not be verified because of signing problems"
	JWTSignatureInvalid      = "signature is invalid"
	JWTExpired               = "token is expired"
	JWTNotValidYet           = "token haven't been valid yet"
	JWTInvalid               = "invalid jwt token"
)
