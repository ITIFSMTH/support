package middlewares

import (
	"fmt"
	"net/http"
	"strings"
	"support-back/responses"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

var (
	SigningKey = "D0NT_T0UCH_M3!_1_L0V3_C4TS"
)

func AuthHandler(authRoles ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get Token, If WS get Token from URL
		var token string
		if c.Request.Header.Get("Upgrade") == "websocket" {
			token = c.Query("token")
		} else {
			token = c.Request.Header.Get("Authorization")
		}

		// Check if toke in correct format
		// ie Bearer: xx03xllasx
		b := "Bearer: "
		if !strings.Contains(token, b) {
			c.AbortWithStatusJSON(http.StatusForbidden, &responses.Response{
				Error: responses.ErrorAuth,
			})
			return
		}
		t := strings.Split(token, b)
		if len(t) < 2 {
			c.AbortWithStatusJSON(http.StatusForbidden, &responses.Response{
				Error: responses.ErrorAuth,
			})
			return
		}

		// Validate token
		valid, err := ValidateToken(t[1], SigningKey)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusForbidden, &responses.Response{
				Error: responses.ErrorAuth,
			})
			return
		}

		// Authorize
		if !contains(authRoles, fmt.Sprintf("%v", valid.Claims.(jwt.MapClaims)["role"])) {
			c.AbortWithStatusJSON(http.StatusForbidden, &responses.Response{
				Error: responses.ErrorAuth,
			})
			return
		}

		// Set variables
		c.Set("id", valid.Claims.(jwt.MapClaims)["id"])
		c.Set("login", valid.Claims.(jwt.MapClaims)["login"])
		c.Set("role", valid.Claims.(jwt.MapClaims)["role"])

		c.Next()
	}
}

func contains(slice []string, item string) bool {
	set := make(map[string]struct{}, len(slice))
	for _, s := range slice {
		set[s] = struct{}{}
	}

	_, ok := set[item]
	return ok
}

func GenerateToken(key []byte, login string, role string, id int) (string, error) {
	// New Token
	token := jwt.New(jwt.SigningMethodHS256)

	// Claims
	claims := make(jwt.MapClaims)
	claims["id"] = id
	claims["login"] = login
	claims["exp"] = time.Now().Add(time.Hour*72).UnixNano() / int64(time.Millisecond)

	// Set user role
	claims["role"] = role

	token.Claims = claims

	// Sign and get as a string
	tokenString, err := token.SignedString(key)
	return tokenString, err
}

func ValidateToken(tokenString string, key string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(key), nil
	})

	return token, err
}
