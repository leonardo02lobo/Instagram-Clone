package routes

import (
	"instagram-clone/controllers"

	"github.com/gin-gonic/gin"
)

func Rutas(r *gin.Engine) {
	api := r.Group("/api")
	{
		api.POST("/registrar", controllers.RegistrarUsuario)
	}
}
