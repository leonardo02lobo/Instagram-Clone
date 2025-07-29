package controllers

import (
	"instagram-clone/config"
	"instagram-clone/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func RegistrarUsuario(c *gin.Context) {
	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if user.Username == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error en el nombre de usuario": "El nombre de usuario no puede estar vacio"})
		return
	}

	if user.Email == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error en el email": "El email no puede estar vacio"})
		return
	}

	if len(user.Password) < 8 {
		c.JSON(http.StatusBadRequest, gin.H{"error en la contraseña": "La contraseña debe tener al menos 8 caracteres"})
		return
	}

	bytes, err1 := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)

	if err1 != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al generar el hash de la contraseña"})
		return
	}

	query := "INSERT INTO users (username,email,password_hash) VALUES ($1,$2,$3) RETURNING user_id"
	err := config.DB.QueryRow(query, user.Username, user.Email, bytes).Scan(&user.ID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Usuario Creado", "user": user})
}
