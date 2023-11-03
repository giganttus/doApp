package services

import (
	"doApp/helpers"
	"doApp/models"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

func (s *Services) CreateDailyOffer(ctx *gin.Context) error {
	userID := ctx.Keys["userID"].(int)

	if (!s.Repos.AllowAccess(userID, "admin")) && (!s.Repos.AllowAccess(userID, "moderator")) {
		return helpers.ErrService
	}

	restaurant, err := s.Repos.GetRestaurantByField("users_id", userID)
	if err != nil {
		return helpers.ErrService
	}

	currentDate := time.Now().Format("2006-01-02")
	if doExists := s.Repos.DailyOfferExists(currentDate, restaurant.ID); doExists {
		return helpers.ErrService
	}

	link, err := helpers.UploadFile(ctx)
	if err != nil {
		fmt.Println(err)
		fmt.Println(link)
		return helpers.ErrService
	}

	var newDailyOffer = models.DailyOffers{
		Link:          link,
		Date:          currentDate,
		RestaurantsID: restaurant.ID,
	}

	return s.Repos.CreateDailyOffer(newDailyOffer)
}

func (s *Services) GetDailyOffers(ctx *gin.Context) ([]models.DailyOffers, error) {
	return s.Repos.GetDailyOffers()
}

func (s *Services) GetDailyOffer(ctx *gin.Context) (models.DailyOffers, error) {
	userID := ctx.Keys["userID"].(int)

	if (!s.Repos.AllowAccess(userID, "admin")) && (!s.Repos.AllowAccess(userID, "moderator")) {
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

	if (!s.Repos.AllowAccess(userID, "admin")) && (!s.Repos.AllowAccess(userID, "moderator")) {
		return helpers.ErrService
	}

	return s.Repos.DeleteDailyOffer(dailyOfferID)
}
