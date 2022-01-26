package middlewares

import (
	// "net/http"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type JwtCitizenClaims struct {
	ID string `json:"id"`
	jwt.StandardClaims
}

type JwtAdminClaims struct {
	ID string `json:"id"`
	jwt.StandardClaims
}

type ConfigJWT struct {
	SecretJWT       string
	ExpiresDuration int
}

func (jwtConf *ConfigJWT) Init() middleware.JWTConfig {
	return middleware.JWTConfig{
		Claims:     &JwtCitizenClaims{},
		SigningKey: []byte(jwtConf.SecretJWT),
	}
}

func (jwtConf *ConfigJWT) GenerateToken(CitizenID string) string {
	claims := JwtCitizenClaims{
		CitizenID,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(int64(jwtConf.ExpiresDuration))).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	jwtToken, _ := token.SignedString([]byte(jwtConf.SecretJWT))

	return jwtToken
}

func GetCitizen(c echo.Context) *JwtCitizenClaims {
	citizen := c.Get("citizen").(*jwt.Token)
	claims := citizen.Claims.(*JwtCitizenClaims)
	return claims
}

func (jwtConf *ConfigJWT) GenerateTokenAdmin(AdminID string) string {
	claims := JwtAdminClaims{
		AdminID,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(int64(jwtConf.ExpiresDuration))).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	jwtToken, _ := token.SignedString([]byte(jwtConf.SecretJWT))

	return jwtToken
}
