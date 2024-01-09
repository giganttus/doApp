package repos

import (
	"doApp/models"

	"gorm.io/gorm"
)

type Repos struct {
	db *gorm.DB
}

func NewRepos(db *gorm.DB) *Repos {
	return &Repos{
		db: db,
	}
}

// Function located in repo, used by services
type IRepos interface {
	// Users - Main functions
	CreateUser(models.Users) error
	GetUsers() ([]models.Users, error)

	// Users - Support functions
	EmailExists(string) bool
	GetUserByField(string, any) (models.Users, error)
	AllowAccess(int, string) bool

	// Restaurants - Main functions
	CreateRestaurant(models.Restaurants) error
	GetRestaurants() ([]models.Restaurants, error)
	UpdateRestaurant(models.Restaurants) error
	DeleteRestaurant(int) error

	// Restaurants - Support functions
	RestaurantExists(string) bool
	GetRestaurantByField(string, any) (models.Restaurants, error)
	RestaurantUserRelation(int) bool

	// Daily offers - Main functions
	CreateDailyOffer(models.DailyOffers) error
	GetDailyOffers() ([]models.DailyOffers, error)
	GetDailyOffer(int) (models.DailyOffers, error)
	DeleteDailyOffer(int) error

	// Daily offers - Support functions
	DailyOfferExists(string, int) bool
	GetDailyOfferByField(string, any) (models.DailyOffers, error)
}
