package repos

import (
	"doApp/helpers"
	"doApp/models"
	"fmt"
)

// Main functions
func (r *Repos) CreateRestaurant(input models.Restaurants) error {
	if err := r.db.Create(&input).Error; err != nil {
		return helpers.ErrRepo
	}

	return nil
}

func (r *Repos) GetRestaurants() ([]models.Restaurants, error) {
	var restaurants []models.Restaurants

	if res := r.db.Find(&restaurants); res.Error != nil {
		return nil, helpers.ErrRepo
	}

	return restaurants, nil
}

func (r *Repos) UpdateRestaurant(input models.Restaurants) error {
	res := r.db.Model(&models.Restaurants{}).Where("id = ?", input.ID).Updates(&input)
	if res.Error != nil || res.RowsAffected == 0 {
		return helpers.ErrRepo
	}

	return nil
}

func (r *Repos) DeleteRestaurant(restaurantID int) error {
	res := r.db.Where("id = ?", restaurantID).Delete(&models.Restaurants{})
	if res.Error != nil || res.RowsAffected == 0 {
		return helpers.ErrRepo
	}

	return nil
}

// Support functions
func (r *Repos) GetRestaurantByField(field string, value any) (models.Restaurants, error) {
	var restaurants models.Restaurants

	if err := r.db.Where(fmt.Sprintf("%v = ?", field), value).First(&restaurants).Error; err != nil {
		return restaurants, helpers.ErrRepo
	}

	return restaurants, nil
}

func (r *Repos) RestaurantExists(name string) bool {
	res := r.db.Where("name = ?", name).First(&models.Restaurants{}).RowsAffected

	return res != 0
}
