package controllers

import (
	"instagram-clone/auth"
	"instagram-clone/config"
	"instagram-clone/models"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func RegistrarUsuario(c *gin.Context) {
	type UserInput struct {
		Username string `json:"username" binding:"required"`
		Email    string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required,min=8"`
	}

	var input UserInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al generar el hash de la contraseña"})
		return
	}

	query := `
        INSERT INTO users (username, email, password_hash)
        VALUES ($1, $2, $3)
        RETURNING user_id, username, email, created_at
    `

	var user models.User
	err = config.DB.QueryRowContext(
		c.Request.Context(),
		query,
		input.Username,
		input.Email,
		string(hashedPassword),
	).Scan(
		&user.UserID,
		&user.Username,
		&user.Email,
		&user.CreatedAt,
	)

	if err != nil {
		log.Printf("Error al registrar usuario: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Error al registrar el usuario",
			"details": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Usuario registrado exitosamente",
		"user": gin.H{
			"user_id":    user.UserID,
			"username":   user.Username,
			"email":      user.Email,
			"created_at": user.CreatedAt,
		},
	})
}

func IniciarSesion(c *gin.Context) {
	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if user.Email == "" || user.PasswordHash == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Email y contraseña son requeridos"})
		return
	}

	var storedUser models.User
	query := "SELECT user_id, username, password_hash FROM users WHERE email = $1"
	err := config.DB.QueryRow(query, user.Email).Scan(&storedUser.UserID, &storedUser.Username, &storedUser.PasswordHash)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Credenciales incorrectas"})
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(storedUser.PasswordHash), []byte(user.PasswordHash))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Credenciales incorrectas"})
		return
	}
	token, err2 := auth.CreateAuthCookie(storedUser)
	if err2 != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al crear sesión"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Inicio de sesión exitoso", "user": storedUser, "Token": token})
}

func GetUsers(c *gin.Context) {
	var users []models.User

	query := `
		SELECT 
			user_id, 
			username, 
			email, 
			password_hash, 
			bio, 
			profile_pic, 
			created_at 
		FROM users
	`

	rows, err := config.DB.Query(query)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Failed to fetch users",
			"details": err.Error(),
		})
		return
	}
	defer rows.Close()

	for rows.Next() {
		var user models.User
		err := rows.Scan(
			&user.UserID,
			&user.Username,
			&user.Email,
			&user.PasswordHash,
			&user.Bio,
			&user.ProfilePic,
			&user.CreatedAt,
		)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error":   "Failed to scan user data",
				"details": err.Error(),
			})
			return
		}
		users = append(users, user)
	}

	if err = rows.Err(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Error after scanning rows",
			"details": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": users,
	})
}
func PerfilUsuario(c *gin.Context) {
	type PerfilInput struct {
		JWTDecode string `json:"jwtDecode"`
		Perfil    bool   `json:"perfil"`
	}

	var input PerfilInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Ahora puedes usar input.JWTDecode y input.Perfil
	dataJWT, err := auth.DecodeJWTFromBody(input.JWTDecode)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if input.Perfil {
		c.JSON(http.StatusOK, gin.H{
			"nombrePerfil": dataJWT,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"nombrePerfil": dataJWT.Username,
		})
	}
}
