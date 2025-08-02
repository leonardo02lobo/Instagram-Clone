package auth

import (
	"database/sql"
	"instagram-clone/config"
	"instagram-clone/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

var jwtSecret = []byte(config.Config.JWTSecret)

type Claims struct {
	UserID     int            `json:"user_id"`
	Username   string         `json:"username"`
	Email      string         `json:"email"`
	Bio        sql.NullString `json:"Bio"`
	ProfilePic sql.NullString `json:"profile_pic"`
	CreatedAt  time.Time      `json:"creadted_at"`
	jwt.RegisteredClaims
}

func CreateAuthCookie(user models.User) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour)

	claims := &Claims{
		UserID:     user.UserID,
		Username:   user.Username,
		Email:      user.Email,
		Bio:        user.Bio,
		ProfilePic: user.ProfilePic,
		CreatedAt:  user.CreatedAt,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func JWTAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		cookie, err := c.Request.Cookie("auth_token")
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "No autorizado - Cookie no encontrada",
			})
			return
		}

		tokenString := cookie.Value
		claims := &Claims{}

		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return jwtSecret, nil
		})

		if err != nil || !token.Valid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "No autorizado - Token inválido",
			})
			return
		}

		c.Set("userID", claims.UserID)
		c.Set("username", claims.Username)
		c.Set("email", claims.Email)

		c.Next()
	}
}

func DecodeJWTFromBody(token string) (*Claims, error) {

	if token == "" {
		return nil, nil
	}

	parsedToken, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := parsedToken.Claims.(*Claims); ok && parsedToken.Valid {
		return claims, nil
	} else {
		return nil, err
	}
}
