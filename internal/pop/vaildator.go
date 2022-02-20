// PopCat Echo
// (c) 2021 SuperSonic (https://github.com/supersonictw).

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
	if count >= 0 && count <= config.PopLimitRedisPopCount {
		return nil
	}
	return EchoError.InvalidCountRange
}

func getJWTIssuer(c *gin.Context) string {
	secret := append([]byte(c.Request.Host), config.PopJWTSecret...)
	hash := sha256.Sum256(secret)
	return fmt.Sprintf("%x", hash)
}

func IssueJWT(issuer, ipAddress, regionCode string) (string, error) {
	now := time.Now()
	claims := jwt.StandardClaims{
		Audience:  ipAddress,
		ExpiresAt: now.Add(config.PopJWTExpired * time.Second).Unix(),
		Id:        uuid.NewString(),
		IssuedAt:  now.Unix(),
		Issuer:    issuer,
		NotBefore: now.Unix(),
		Subject:   regionCode,
	}
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return tokenClaims.SignedString(config.PopJWTSecret)
}

func ValidateCaptcha(ipAddress, token string) error {
	if !config.PopReCaptchaStatus {
		return nil
	}
	if token == "" {
		return EchoError.EmptyCaptchaToken
	}
	result, err := recaptcha.Confirm(ipAddress, token)
	if err != nil {
		return err
	}
	if !result {
		return EchoError.UnsafeCaptchaToken
	}
	return err
}

func ValidateJWT(c *gin.Context, ctx context.Context, token string) (bool, *jwt.StandardClaims, error) {
	tokenClaims, err := jwt.ParseWithClaims(
		token,
		&jwt.StandardClaims{},
		func(token *jwt.Token) (i interface{}, err error) {
			return config.PopJWTSecret, nil
		},
	)
	if err != nil {
		return false, nil, err
	}
	if claims, ok := tokenClaims.Claims.(*jwt.StandardClaims); ok && tokenClaims.Valid {
		ipAddress := c.ClientIP()
		issuer := getJWTIssuer(c)
		regionCode, err := GetRegionCode(ctx, ipAddress)
		if err != nil {
			return false, nil, err
		}
		return claims.Issuer == issuer &&
			claims.Subject == regionCode &&
			claims.Audience == ipAddress, claims, nil
	}
	return false, nil, nil
}

func GetRegionCode(ctx context.Context, ipAddress string) (string, error) {
	if gin.Mode() != gin.ReleaseMode {
		return "##", nil
	}
	if value := queryRegionCodeFromRedis(ctx, ipAddress); value != "" {
		return value, nil
	}
	if value := queryRegionCodeFromAPI(ipAddress); value != "" {
		key := fmt.Sprintf("%s:%s", config.CacheNamespaceGeo, ipAddress)
		err := redisClient.Set(ctx, key, value, time.Hour).Err()
		if err != nil {
			log.Println(err)
			return value, err
		}
		return value, nil
	}
	return "", EchoError.UnknownRegionCode
}

func queryRegionCodeFromRedis(ctx context.Context, ipAddress string) string {
	key := fmt.Sprintf("%s:%s", config.CacheNamespaceGeo, ipAddress)
	return redisClient.Get(ctx, key).Val()
}

func queryRegionCodeFromAPI(ipAddress string) string {
	resp, err := http.Get("https://restapi.starinc.xyz/basic/network/ip/geo?ip_addr=" + ipAddress)
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
	if config.PopLimitRedisDuration == 0 {
		return nil
	}
	sum := GetAddressCountInRefreshInterval(ctx, address)
	if sum > config.PopLimitRedisPopCount {
		return EchoError.AddressRateLimited
	}
	return nil
}
