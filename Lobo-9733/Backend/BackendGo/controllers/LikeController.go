package controllers

import (
	"database/sql"
	"instagram-clone/auth"
	"instagram-clone/config"
	"net/http"

	"github.com/gin-gonic/gin"
)

type LikeRequest struct {
	PostID int `json:"post_id"`
	UserID int `json:"user_id"`
}

func DarLike(c *gin.Context) {
	var LikeInput LikeRequest

	if err := c.ShouldBindJSON(&LikeInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Datos inválidos",
			"details": err.Error(),
		})
		return
	}

	var exists bool
	err := config.DB.QueryRow(
		"SELECT EXISTS(SELECT 1 FROM likes WHERE user_id = $1 AND post_id = $2)",
		LikeInput.UserID, LikeInput.PostID,
	).Scan(&exists)

	if err != nil && err != sql.ErrNoRows {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Error al verificar like existente",
		})
		return
	}

	var query string
	var action string

	if exists {
		query = "DELETE FROM likes WHERE user_id = $1 AND post_id = $2"
		action = "unliked"
	} else {
		query = "INSERT INTO likes (user_id, post_id) VALUES ($1, $2)"
		action = "liked"
	}

	_, err = config.DB.Exec(query, LikeInput.UserID, LikeInput.PostID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Error al actualizar like",
		})
		return
	}

	var likesCount int
	err = config.DB.QueryRow(
		"SELECT COUNT(*) FROM likes WHERE post_id = $1",
		LikeInput.PostID,
	).Scan(&likesCount)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Error al obtener conteo de likes",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success":     true,
		"action":      action,
		"likes_count": likesCount,
	})
}

func GetLikesCount(c *gin.Context) {
	var request LikeRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Datos inválidos",
			"details": err.Error(),
		})
		return
	}

	if request.PostID <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "El ID de publicación debe ser un número positivo",
		})
		return
	}

	var likesCount int
	err := config.DB.QueryRow(
		"SELECT COUNT(*) FROM likes WHERE post_id = $1",
		request.PostID,
	).Scan(&likesCount)

	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusOK, gin.H{
				"post_id":     request.PostID,
				"likes_count": 0,
			})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Error al obtener likes",
			"details": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"post_id":     request.PostID,
		"likes_count": likesCount,
	})
}

type CheckLikeRequest struct {
	PostID int    `json:"post_id"`
	Token  string `json:"token"`
}

func CheckUserLike(c *gin.Context) {
	var req CheckLikeRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Datos inválidos",
			"details": err.Error(),
		})
		return
	}

	dataJWT, err := auth.DecodeJWTFromBody(req.Token)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error":   "Token inválido",
			"details": err.Error(),
		})
		return
	}

	var exists bool
	err = config.DB.QueryRow(
		"SELECT EXISTS(SELECT 1 FROM likes WHERE user_id = $1 AND post_id = $2)",
		dataJWT.UserID, req.PostID,
	).Scan(&exists)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Error al verificar like",
			"details": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"post_id":   req.PostID,
		"user_id":   dataJWT.ID,
		"has_liked": exists,
	})
}
