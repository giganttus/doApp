package routes

import (
	"doApp/handlers"
	"doApp/middlewares"
	"doApp/repos"
	"doApp/services"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func InitRoutes(db *gorm.DB) *gin.Engine {
	router := gin.Default()
	router.Use(middlewares.CorsMiddleware())
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	rep := repos.NewRepos(db)
	svc := services.NewServices(rep)
	hlr := handlers.NewHandlers(svc)

	apiRouter := router.Group("/api")
	noAuthRouter := apiRouter.Group("/")
	authRouter := apiRouter.Group("/").Use(middlewares.AuthMiddleware())

	// Routes wo Auth
	// Users
	noAuthRouter.POST("register", hlr.CreateUser)
	noAuthRouter.POST("login", hlr.Login)

	// Routes w Auth
	// Users
	authRouter.GET("users", hlr.GetUsers)

	// Restaurants
	authRouter.POST("restaurant", hlr.CreateRestaurant)
	authRouter.GET("restaurant", hlr.GetRestaurants)
	authRouter.PUT("restaurant", hlr.UpdateRestaurant)
	authRouter.DELETE("restaurant/:id", hlr.DeleteRestaurant)

	// Daily Offers
	authRouter.POST("dailyOffer", hlr.CreateDailyOffer)
	authRouter.GET("dailyOffers", hlr.GetDailyOffers)
	authRouter.GET("dailyOffer", hlr.GetDailyOffer)
	authRouter.GET("dailyOffer/:restaurantId", hlr.GetDailyOfferByRestaurantID)
	authRouter.DELETE("dailyOffer/:id", hlr.DeleteDailyOffer)

	return router
}
