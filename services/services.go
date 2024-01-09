package services

import (
	"doApp/models"
	"doApp/repos"

	"github.com/gin-gonic/gin"
)

type Services struct {
	Repos repos.IRepos
}

func NewServices(repos repos.IRepos) *Services {
	return &Services{
		Repos: repos,
	}
}

// Function located in services, used by handlers
type IServices interface {
	// Users
	CreateUser(*gin.Context, models.Users) error
	Login(*gin.Context, models.LoginReq) (models.LoginRes, error)
	GetUsers(*gin.Context) ([]models.Users, error)

	// Restaurants
	CreateRestaurants(*gin.Context, models.Restaurants) error
	GetRestaurants(*gin.Context) ([]models.Restaurants, error)
	UpdateRestaurant(*gin.Context, models.Restaurants) error
	DeleteRestaurant(*gin.Context, int) error

	// Daily Offers
	CreateDailyOffer(*gin.Context) error
	GetDailyOffers(*gin.Context) ([]models.DailyOffers, error)
	GetDailyOffer(*gin.Context) (models.DailyOffers, error)
	GetDailyOfferByRestaurantID(*gin.Context, int) (models.DailyOffers, error)
	DeleteDailyOffer(*gin.Context, int) error
}
