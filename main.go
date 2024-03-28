package main

import (
	"github.com/janortop5/go-ecommerce-yt/controllers"
	"github.com/janortop5/go-ecommerce-yt/database"
	"github.com/janortop5/go-ecommerce-yt/middleware"
	"github.com/janortop5/go-ecommerce-yt/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	port := os.Getenv("PORT");
	if port == "" {
		port = "8000"
	}

	app:= controllers.NewApplication(database.ProductData(database.Client, "Products"), database.UserData(database.Client, "Users"));

	router:= gin.New(); // 'router' is used to access gin's methods
	router.Use(gin.Logger()); // here we use a gin method by using the 'router' variable which has been set to access gin's methods. 
							  // gin.Logger()' is used to log the requests

	routes.UserRoutes(router); // we pass router to the UserRoutes function in routes/routes.go
	
	router.Use(middleware.Authentication());
	router.GET("/addtocart", app.AddToCart())
	router.GET("/removeitem", app.RemoveItem());
	router.GET("/cartcheckout", app.BuyFromCart());
	router.GET("/instantbuy", app.InstantBuy());

	log.Fatal(router.Run(":" + port))
}