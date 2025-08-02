package routes

import (
	"instagram-clone/auth"
	"instagram-clone/controllers"

	"github.com/gin-gonic/gin"
)

func Rutas(r *gin.Engine) {
	api := r.Group("/api")
	{
		api.POST("/registrar", controllers.RegistrarUsuario)
		api.POST("/login", controllers.IniciarSesion)
		api.GET("/ObtenerUsuario", controllers.GetUsers)
		api.Use(auth.JWTAuthMiddleware())
		{
			api.GET("/perfil", controllers.PerfilUsuario)
		}

	}
}
