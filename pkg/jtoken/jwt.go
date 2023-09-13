package jtoken

import (
	"github.com/golang-jwt/jwt"
	"github.com/quangdangfit/gocommon/logger"
	"goshop/pkg/config"
	"goshop/pkg/utils"
	"strings"
	"time"
)

const (
	ACCESS_TOKEN_EXIRED_TIME   = 5 * 3600       // 5 hours
	REFRESH_TOKEN_EXPIRED_TIME = 30 * 24 * 3600 // 30 days
	ACCESS_TOKEN_TYPE          = "x-access"
	REFRESH_TOKEN_TYPE         = "x-refresh"
)

func GenerateAccessToken(payload map[string]interface{}) string {
	cfg := config.GetConfig()
	payload["type"] = ACCESS_TOKEN_TYPE
	tokenClaims := jwt.MapClaims{
		"payload": payload,
		"exp":     time.Now().Add(time.Second * ACCESS_TOKEN_EXIRED_TIME).Unix(),
	}
	jwtToken := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tokenClaims)
	token, err := jwtToken.SignedString([]byte(cfg.AuthSecret))
	if err != nil {
		logger.Error("Failed to generate access token: ", err)
		return ""
	}

	return token
}

func GenerateRefreshToken(payload map[string]interface{}) string {
	cfg := config.GetConfig()
	payload["type"] = REFRESH_TOKEN_TYPE
	tokenClaims := jwt.MapClaims{
		"payload": payload,
		"exp":     time.Now().Add(time.Second * REFRESH_TOKEN_EXPIRED_TIME).Unix(),
	}
	unSignedToken := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tokenClaims)
	token, err := unSignedToken.SignedString([]byte(cfg.AuthSecret))
	if err != nil {
		logger.Error("Failed to generate refresh token: ", err)
	}

	return token
}

func ValidateToken(jwtToken string) (map[string]interface{}, error) {
	cfg := config.GetConfig()
	jwtFromRequest := strings.Replace(jwtToken, "Bearer ", "", -1)
	tokenClaims := jwt.MapClaims{}
	token, err := jwt.ParseWithClaims(jwtFromRequest, tokenClaims, func(token *jwt.Token) (interface{}, error) {
		return []byte(cfg.AuthSecret), nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, jwt.ErrInvalidKey
	}

	var data map[string]interface{}
	utils.Copy(&data, tokenClaims["payload"])

	return data, nil
}
