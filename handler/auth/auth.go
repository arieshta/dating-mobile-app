package auth

import (
	"fmt"
	"strings"
	"time"

	"github.com/arieshta/dating-mobile-app/internal/config"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

func CreateToken(userId string) (string, error) {
	// Create a new token object, specifying signing method and the claims
	// you would like it to contain.
	token := jwt.New(jwt.SigningMethodHS256)

	sysConfig := config.Config{}

	// Set claims
	claims := token.Claims.(jwt.MapClaims)
	claims["exp"] = time.Now().Add(time.Duration(time.Minute*500)).Unix()
	claims["userid"] = userId

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte(sysConfig.JWT_SECRET))
	if err != nil {
		return "", err
	}
	return t, nil
}

func VerifyToken(tokenString string) (jwt.MapClaims, error) {
	// Parse the token
	tokenString = strings.Split(tokenString, " ")[1]
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.Config{}.JWT_SECRET), nil
	})
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}
	return claims, nil
}

func Auth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ec echo.Context) error {
		token := ec.Request().Header.Get("Authorization")
		claims, err := VerifyToken(token)
		fmt.Println(claims)
		if err != nil {
			return ec.JSON(401, "Unauthorized")
		}
		ec.Set("userid", claims["userid"])
		return next(ec)
	}
}