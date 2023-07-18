// PopSenzawa Echo
// (c) 2023 SuperSonic (https://github.com/supersonictw).

package pop

import (
	"crypto/sha256"
	"errors"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"github.com/spf13/viper"
	"github.com/supersonictw/popsenzawa-echo/data"
	"golang.org/x/exp/slices"
)

var (
	ErrJwtEmpty   = errors.New("jwt empty")
	ErrJwtInvalid = errors.New("jwt invalid")
)

var (
	configServerAddress      = viper.GetString("server.address")
	configPopJwtSecretString = viper.GetString("jwt.secret")
)

var (
	popJwtSecretBytes = []byte(configPopJwtSecretString)
)

var (
	serverID = sha256.Sum256(append([]byte(configServerAddress), popJwtSecretBytes...))
)

func getJwtMetadataFromContext(c *gin.Context) (string, string, string, error) {
	visitorIP := data.ParseVisitorIP(c.ClientIP())

	ipAddress := visitorIP.NetIP().String()
	regionCode, err := visitorIP.RegionCode()
	if err != nil {
		return "", "", "", err
	}

	issuer := fmt.Sprintf("%x", serverID)
	audience := ipAddress
	subject := regionCode

	return issuer, audience, subject, nil
}

func issueJwt(issuer, audience, subject string) (string, error) {
	now := time.Now()

	issueAt := jwt.NewNumericDate(now)
	notBefore := jwt.NewNumericDate(now.Add(-60 * time.Second))
	expireAt := jwt.NewNumericDate(now.Add(60 * time.Second))

	claims := &jwt.RegisteredClaims{
		ID:     uuid.NewString(),
		Issuer: issuer,
		Audience: jwt.ClaimStrings([]string{
			audience,
		}),
		Subject:   subject,
		IssuedAt:  issueAt,
		NotBefore: notBefore,
		ExpiresAt: expireAt,
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return tokenClaims.SignedString(popJwtSecretBytes)
}

func issueJwtFromContext(c *gin.Context) (string, error) {
	issuer, audience, subject, err := getJwtMetadataFromContext(c)
	if err != nil {
		return "", err
	}

	token, err := issueJwt(issuer, audience, subject)
	if err != nil {
		return "", err
	}

	return token, nil
}

func validateJwt(issuer, audience, subject, token string) (*jwt.RegisteredClaims, error) {
	if token == "" {
		return nil, ErrJwtEmpty
	}

	tokenClaims, err := jwt.ParseWithClaims(
		token,
		&jwt.RegisteredClaims{},
		func(token *jwt.Token) (i interface{}, err error) {
			return popJwtSecretBytes, nil
		},
	)
	if err != nil {
		return nil, err
	}

	claims, ok := tokenClaims.Claims.(*jwt.RegisteredClaims)
	if !ok || !tokenClaims.Valid {
		return nil, ErrJwtInvalid
	}

	if claims.Issuer != issuer ||
		!slices.Contains(claims.Audience, audience) ||
		claims.Subject != subject {
		return nil, ErrJwtInvalid
	}

	return claims, nil
}

func validateJwtFromContext(c *gin.Context) (*jwt.RegisteredClaims, error) {
	issuer, audience, subject, err := getJwtMetadataFromContext(c)
	if err != nil {
		return nil, err
	}

	token := c.Query("token")
	return validateJwt(issuer, audience, subject, token)
}
