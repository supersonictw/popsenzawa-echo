package pop

import (
	"context"
	"crypto/sha256"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/dpapathanasiou/go-recaptcha"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"github.com/supersonictw/popcat-echo/internal"
	echoErrors "github.com/supersonictw/popcat-echo/internal/error"
	"io/ioutil"
	"net/http"
	"time"
)

func getJWTIssuer(c *gin.Context) string {
	hash := sha256.Sum256([]byte(c.GetHeader("Host") + internal.JWTCaptchaSecret))
	return fmt.Sprintf("%x", hash)
}

func IssueJWT(c *gin.Context, ctx context.Context) (string, error) {
	now := time.Now()
	ipAddress := c.ClientIP()
	issuer := getJWTIssuer(c)
	regionCode, err := GetRegionCode(ctx, ipAddress)
	if err != nil {
		return "", err
	}
	claims := jwt.StandardClaims{
		Audience:  ipAddress,
		ExpiresAt: now.Add(20 * time.Second).Unix(),
		Id:        uuid.NewString(),
		IssuedAt:  now.Unix(),
		Issuer:    issuer,
		NotBefore: now.Add(10 * time.Second).Unix(),
		Subject:   regionCode,
	}
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return tokenClaims.SignedString(internal.JWTCaptchaSecret)
}

func ValidateCaptcha(ipAddress string, token string) bool {
	if !internal.ReCaptchaStatus {
		return true
	}
	if token == "" {
		return false
	}
	result, err := recaptcha.Confirm(ipAddress, token)
	if err != nil {
		panic(err)
	}
	return result
}

func ValidateJWT(c *gin.Context, token string) (bool, error) {
	tokenClaims, err := jwt.ParseWithClaims(
		token,
		&jwt.StandardClaims{},
		func(token *jwt.Token) (i interface{}, err error) {
			return internal.JWTCaptchaSecret, nil
		},
	)
	if err != nil {
		return false, err
	}
	if claims, ok := tokenClaims.Claims.(*jwt.StandardClaims); ok && tokenClaims.Valid {
		return claims.Issuer == getJWTIssuer(c), nil
	}
	return false, nil
}

func GetRegionCode(ctx context.Context, ipAddress string) (string, error) {
	if value := queryRegionCodeFromRedis(ctx, ipAddress); value != "" {
		return value, nil
	}
	if value := queryRegionCodeFromAPI(ipAddress); value != "" {
		return value, nil
	}
	return "", errors.New(echoErrors.UnknownRegionCode)
}

func queryRegionCodeFromRedis(ctx context.Context, ipAddress string) string {
	key := fmt.Sprintf("%s:%s", internal.CacheNamespaceGeo, ipAddress)
	return internal.RDB.Get(ctx, key).Val()
}

func queryRegionCodeFromAPI(ipAddress string) string {
	resp, err := http.Get("https://restapi.starinc.xyz/basic/ip/geo?ip_addr=" + ipAddress)
	if err != nil {
		panic(err)
	}
	var result map[string]interface{}
	resultBytes, err := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(resultBytes, &result)
	if err != nil {
		panic(err)
	}
	if country, ok := result["country"].(map[string]interface{}); ok {
		return country["iso_code"].(string)
	}
	return ""
}
