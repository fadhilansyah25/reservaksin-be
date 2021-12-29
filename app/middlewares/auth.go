package middlewares

import (
	// "net/http"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type JwtCustomClaims struct {
	ID int `json:"id"`
	jwt.StandardClaims
}

type ConfigJWT struct {
	SecretJWT       string
	ExpiresDuration int
}

func (jwtConf *ConfigJWT) Init() middleware.JWTConfig {
	return middleware.JWTConfig{
		Claims:     &JwtCustomClaims{},
		SigningKey: []byte(jwtConf.SecretJWT),
		// ErrorHandlerWithContext: middleware.JWTErrorHandlerWithContext(func(e error, c echo.Context) error {
		// 	return controllers.NewErrorResponse(c, http.StatusForbidden, e)
		// }),
	}
}

func (jwtConf *ConfigJWT) GenerateToken(CitizenID int) string {
	claims := JwtCustomClaims{
		CitizenID,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(int64(jwtConf.ExpiresDuration))).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	jwtToken, _ := token.SignedString([]byte(jwtConf.SecretJWT))

	return jwtToken
}

func GetCitizen(c echo.Context) *JwtCustomClaims {
	citizen := c.Get("citizen").(*jwt.Token)
	claims := citizen.Claims.(*JwtCustomClaims)
	return claims
}
