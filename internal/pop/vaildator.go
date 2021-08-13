package pop

import "github.com/gin-gonic/gin"

func IssueJWT(c *gin.Context) string {
	return ""
}

func ValidateCaptcha(token string) bool {
	return false
}

func ValidateJWT(token string) bool {
	return false
}

func queryRegionCodeFromRedis(ipAddress string) string {
	return ""
}

func GetRegionCode(ipAddress string) string {
	return ""
}

func queryRegionCodeFromAPI(ipAddress string) string {
	return ""
}
