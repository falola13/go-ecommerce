package main

import (
	"log"
	"os"

	"github.com/falola13/go-ecommerce/controllers"
	"github.com/falola13/go-ecommerce/database"
	"github.com/falola13/go-ecommerce/middleware"
	"github.com/falola13/go-ecommerce/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	app := controllers.NewApplication(database.ProductData(database.Client, " Products"), database.UserData(database.Client, "Users"))

	router := gin.New()
	router.Use(gin.Logger())

	routes.UserRoutes(router)

	authenticated := router.Group("/")
	authenticated.Use(middleware.Authentication())
	{
		authenticated.GET("/addtocart", app.AddToCart())
		authenticated.GET("/removeitem", app.RemoveItem())
		authenticated.GET("/cartcheckout", app.BuyFromCart())
		authenticated.GET("/instantbuy", app.InstantBuy())
		authenticated.POST("/address/:id", controllers.AddAddress())
		authenticated.DELETE("/address/:id", controllers.DeleteAddress())
	}

	log.Fatal(router.Run(":" + port))
}
