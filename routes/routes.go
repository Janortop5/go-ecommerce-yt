package routes

import (
	"github.com/janortop5/go-ecommerce-yt/controllers"
	"github.com/gin-gonic/gin"

)

func UserRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.POST("users/signup", controllers.SignUP());
	incomingRoutes.POST("users/login", controllers.Login());
	incomingRoutes.POST("/admin/addproduct", controllers.ProductViewerAdmin());
	incomingRoutes.GET("/users/productview",controllers.SearchProduct());
	incomingRoutes.GET("/users/search", controllers.SearchProductByQuery());
}