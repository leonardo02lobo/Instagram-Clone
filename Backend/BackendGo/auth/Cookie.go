package auth

import (
	"instagram-clone/config"
	"instagram-clone/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

// Clave secreta para firmar los JWT (debería estar en variables de entorno)
var jwtSecret = []byte(config.Config.JWTSecret)

// Estructura para los claims del JWT
type Claims struct {
	UserID   int    `json:"user_id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	jwt.RegisteredClaims
}

func CreateAuthCookie(c *gin.Context, user models.User) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour)

	claims := &Claims{
		UserID:   user.UserID,
		Username: user.Username,
		Email:    user.Email,
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

// Middleware para validar el JWT en las cookies
func JWTAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Obtener la cookie
		cookie, err := c.Request.Cookie("auth_token")
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "No autorizado - Cookie no encontrada",
			})
			return
		}

		// Parsear el token JWT
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

		// Añadir los claims al contexto para uso posterior
		c.Set("userID", claims.UserID)
		c.Set("username", claims.Username)
		c.Set("email", claims.Email)

		c.Next()
	}
}

// Obtener datos del usuario desde el JWT
func GetUserFromCookie(c *gin.Context) (*Claims, error) {
	cookie, err := c.Request.Cookie("auth_token")
	if err != nil {
		return nil, err
	}

	tokenString := cookie.Value
	claims := &Claims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if err != nil || !token.Valid {
		return nil, err
	}

	return claims, nil
}

func Logout(c *gin.Context) {
	// Crear cookie de expiración inmediata
	cookie := &http.Cookie{
		Name:     "auth_token",
		Value:    "",
		Expires:  time.Unix(0, 0),
		Path:     "/",
		HttpOnly: true,
		Secure:   false,
		SameSite: http.SameSiteLaxMode,
	}

	http.SetCookie(c.Writer, cookie)
	c.JSON(http.StatusOK, gin.H{"message": "Sesión cerrada correctamente"})
}
