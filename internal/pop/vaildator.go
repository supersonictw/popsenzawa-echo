package pop

import (
	"context"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"github.com/dpapathanasiou/go-recaptcha"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"github.com/supersonictw/popcat-echo/internal/config"
	EchoError "github.com/supersonictw/popcat-echo/internal/error"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func ValidateRange(count int) error {
	if count >= 0 && count <= config.PopLimit {
		return nil
	}
	return EchoError.NewError(EchoError.InvalidCountRange)
}

func getJWTIssuer(c *gin.Context) string {
	secret := append([]byte(c.Request.Host), config.JWTCaptchaSecret...)
	hash := sha256.Sum256(secret)
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
		ExpiresAt: now.Add(config.JWTExpired * time.Second).Unix(),
		Id:        uuid.NewString(),
		IssuedAt:  now.Unix(),
		Issuer:    issuer,
		NotBefore: now.Unix(),
		Subject:   regionCode,
	}
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return tokenClaims.SignedString(config.JWTCaptchaSecret)
}

func ValidateCaptcha(ipAddress string, token string) error {
	if !config.ReCaptchaStatus {
		return nil
	}
	if token == "" {
		return EchoError.NewError(EchoError.EmptyCaptchaToken)
	}
	result, err := recaptcha.Confirm(ipAddress, token)
	if err != nil {
		return err
	}
	if !result {
		return EchoError.NewError(EchoError.UnsafeCaptchaToken)
	}
	return err
}

func ValidateJWT(c *gin.Context, token string) (bool, error) {
	tokenClaims, err := jwt.ParseWithClaims(
		token,
		&jwt.StandardClaims{},
		func(token *jwt.Token) (i interface{}, err error) {
			return config.JWTCaptchaSecret, nil
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
		key := fmt.Sprintf("%s:%s", config.CacheNamespaceGeo, ipAddress)
		err := redisClient.Set(ctx, key, value, 3600).Err()
		if err != nil {
			log.Println(err)
			return value, err
		}
		return value, nil
	}
	return "", EchoError.NewError(EchoError.UnknownRegionCode)
}

func queryRegionCodeFromRedis(ctx context.Context, ipAddress string) string {
	key := fmt.Sprintf("%s:%s", config.CacheNamespaceGeo, ipAddress)
	return redisClient.Get(ctx, key).Val()
}

func queryRegionCodeFromAPI(ipAddress string) string {
	resp, err := http.Get("https://restapi.starinc.xyz/basic/ip/geo?ip_addr=" + ipAddress)
	if err != nil {
		log.Panicln(err)
	}
	result := make(map[string]interface{})
	resultBytes, err := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(resultBytes, &result)
	if err != nil {
		log.Panicln(err)
	}
	if data, ok := result["data"].(map[string]interface{}); ok {
		if country, ok := data["country"].(map[string]interface{}); ok {
			return country["iso_code"].(string)
		}
	}
	return ""
}

func ValidateAddressRate(ctx context.Context, address string) error {
	if config.RateLimit == 0 {
		return nil
	}
	sum := GetAddressCountInRefreshInterval(ctx, address)
	if sum > config.RateLimit {
		return EchoError.NewError(EchoError.AddressRateLimited)
	}
	return nil
}
