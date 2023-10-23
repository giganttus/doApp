package services

import (
	"doApp/helpers"
	"doApp/models"
	"time"

	"github.com/gin-gonic/gin"
)

func (s *Services) CreateDailyOffer(ctx *gin.Context, input models.DailyOffers) error {
	userID := ctx.Keys["userID"].(int)

	if (s.Repos.AllowAccess(userID, "admin") == false) || (s.Repos.AllowAccess(userID, "admin") == false) {
		return helpers.ErrService
	}

	currentDate := time.Now().Format("2006-01-02")
	if doExists := s.Repos.DailyOfferExists(currentDate, input.RestaurantsID); doExists {
		return helpers.ErrService
	}

	input.Date = currentDate

	return s.Repos.CreateDailyOffer(input)
}

func (s *Services) GetDailyOffers(ctx *gin.Context) ([]models.DailyOffers, error) {
	return s.Repos.GetDailyOffers()
}

func (s *Services) GetDailyOffer(ctx *gin.Context) (models.DailyOffers, error) {
	userID := ctx.Keys["userID"].(int)

	if (s.Repos.AllowAccess(userID, "admin") == false) || (s.Repos.AllowAccess(userID, "admin") == false) {
		return models.DailyOffers{}, helpers.ErrService
	}

	restaurant, err := s.Repos.GetRestaurantByField("users_id", userID)
	if err != nil {
		return models.DailyOffers{}, helpers.ErrService
	}

	return s.Repos.GetDailyOffer(restaurant.ID)
}

func (s *Services) GetDailyOfferByRestaurantID(ctx *gin.Context, restaurantID int) (models.DailyOffers, error) {
	return s.Repos.GetDailyOfferByField("restaurants_id", restaurantID)
}

func (s *Services) DeleteDailyOffer(ctx *gin.Context, dailyOfferID int) error {
	userID := ctx.Keys["userID"].(int)

	if (s.Repos.AllowAccess(userID, "admin") == false) || (s.Repos.AllowAccess(userID, "admin") == false) {
		return helpers.ErrService
	}

	return s.Repos.DeleteDailyOffer(dailyOfferID)
}
