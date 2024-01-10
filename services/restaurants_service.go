package services

import (
	"doApp/helpers"
	"doApp/models"

	"github.com/gin-gonic/gin"
)

func (s *Services) CreateRestaurants(ctx *gin.Context, input models.Restaurants) error {
	userID := ctx.Keys["userID"].(int)

	if !s.Repos.AllowAccess(userID, "admin") {
		return helpers.ErrForbidden
	}

	if restauEx := s.Repos.RestaurantExists(input.Name); restauEx {
		return helpers.ErrService
	}

	if input.UsersID != 0 {
		if restUserRelEx := s.Repos.RestaurantUserRelation(input.UsersID); restUserRelEx {
			return helpers.ErrService
		}
	}

	return s.Repos.CreateRestaurant(input)
}

func (s *Services) GetRestaurants(ctx *gin.Context) ([]models.Restaurants, error) {
	return s.Repos.GetRestaurants()
}

func (s *Services) UpdateRestaurant(ctx *gin.Context, input models.Restaurants) error {
	userID := ctx.Keys["userID"].(int)

	if !s.Repos.AllowAccess(userID, "admin") {
		return helpers.ErrService
	}

	if input.Name != "" {
		if restauEx := s.Repos.RestaurantExists(input.Name); restauEx {
			return helpers.ErrService
		}
	}

	if input.UsersID != 0 {
		if restUserRelEx := s.Repos.RestaurantUserRelation(input.UsersID); restUserRelEx {
			return helpers.ErrService
		}
	}

	restaurant, err := s.Repos.GetRestaurantByField("id", input.ID)
	if err != nil || input == restaurant {
		return helpers.ErrService
	}

	return s.Repos.UpdateRestaurant(input)
}

func (s *Services) DeleteRestaurant(ctx *gin.Context, RestaurantID int) error {
	userID := ctx.Keys["userID"].(int)

	if !s.Repos.AllowAccess(userID, "admin") {
		return helpers.ErrService
	}

	relEx, err := s.Repos.GetRestaurantByField("id", RestaurantID)
	if err != nil {
		return helpers.ErrService
	}

	if relEx.UsersID != 0 {
		return helpers.ErrService
	}

	return s.Repos.DeleteRestaurant(RestaurantID)
}
