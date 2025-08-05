package controllers

import (
	"fmt"
	"instagram-clone/config"
	"instagram-clone/models"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CrearPublicacion(c *gin.Context) {
	type StructPost struct {
		PostID    int    `json:"post_id"`
		UserID    int    `json:"user_id"`
		Caption   string `json:"caption"`
		MediaURL  string `json:"media_url"`
		MediaType string `json:"media_type"`
	}

	var data StructPost

	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	fmt.Println(data)

	query := "insert into posts(user_id,caption,media_url,media_type) values ($1,$2,$3,$4)"
	_, err := config.DB.ExecContext(
		c.Request.Context(),
		query,
		data.UserID,
		data.Caption,
		data.MediaURL,
		data.MediaType,
	)
	if err != nil {
		log.Printf("Error al registrar post: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Error al registrar el posts",
			"details": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "posts registrado exitosamente",
	})
}

func ObtenerPublicaciones(c *gin.Context) {
	var publicaciones []models.Posts

	var users []models.User

	query := `SELECT * FROM posts;`

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
		var publicacion models.Posts
		err := rows.Scan(
			&publicacion.PostID,
			&publicacion.UserId,
			&publicacion.Caption,
			&publicacion.MediaUrl,
			&publicacion.MediaType,
			&publicacion.CreatedAt,
		)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error":   "Failed to scan user posts",
				"details": err.Error(),
			})
			return
		}
		query := `
		SELECT 
			user_id, 
			username, 
			email, 
			bio, 
			profile_pic, 
			created_at 
		FROM users 
		WHERE user_id = $1
	`
		var user models.User
		err = config.DB.QueryRow(query, publicacion.UserId).Scan(
			&user.UserID,
			&user.Username,
			&user.Email,
			&user.Bio,
			&user.ProfilePic,
			&user.CreatedAt,
		)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Usuario no encontrado"})
			return
		}

		publicaciones = append(publicaciones, publicacion)
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
		"data": publicaciones,
		"user": users,
	})
}
