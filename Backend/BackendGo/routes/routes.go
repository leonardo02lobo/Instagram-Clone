package routes

import (
	"instagram-clone/controllers"

	"github.com/gin-gonic/gin"
)

func Rutas(r *gin.Engine) {
	api := r.Group("/api")
	{
		api.POST("/registrar", controllers.RegistrarUsuario)
		api.POST("/login", controllers.IniciarSesion)
		api.GET("/ObtenerUsuario", controllers.GetUsers)
		api.POST("/perfil", controllers.PerfilUsuario)
		api.POST("/CrearPublicacion", controllers.CrearPublicacion)
		api.POST("/ObtenerUsuarioPorNombre", controllers.GetUserByName)
		api.GET("/ObtenerPublicaciones", controllers.ObtenerPublicaciones)
	}
}
