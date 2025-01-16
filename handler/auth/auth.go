package auth

import (
	"fmt"
	"strings"
	"time"

	"github.com/arieshta/dating-mobile-app/internal/config"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

func CreateToken(userId string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	sysConfig := config.Config{}

	claims := token.Claims.(jwt.MapClaims)
	claims["exp"] = time.Now().Add(time.Duration(time.Minute * 500)).Unix()
	claims["userid"] = userId

	t, err := token.SignedString([]byte(sysConfig.JWT_SECRET))
	if err != nil {
		return "", err
	}
	return t, nil
}

func VerifyToken(tokenString string) (jwt.MapClaims, error) {
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
		if token == "" {
			return ec.JSON(401, "Unauthorized")
		}

		claims, err := VerifyToken(token)
		if err != nil {
			return ec.JSON(401, "Unauthorized")
		}

		ec.Set("userid", claims["userid"])

		return next(ec)
	}
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	fmt.Println("=======", err)
	return err == nil
}