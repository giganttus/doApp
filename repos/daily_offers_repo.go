package repos

import (
	"doApp/helpers"
	"doApp/models"
	"fmt"
)

// Main functions
func (r *Repos) CreateDailyOffer(input models.DailyOffers) error {
	if err := r.db.Create(&input).Error; err != nil {
		return helpers.ErrRepo
	}

	return nil
}

func (r *Repos) GetDailyOffers() ([]models.DailyOffers, error) {
	var dailyOffers []models.DailyOffers

	if res := r.db.Find(&dailyOffers); res.Error != nil {
		return nil, helpers.ErrRepo
	}

	return dailyOffers, nil
}

func (r *Repos) GetDailyOffer(restaurantID int) (models.DailyOffers, error) {
	var dailyOffer models.DailyOffers

	if err := r.db.Where("restaurants_id = ?", restaurantID).First(&dailyOffer).Error; err != nil {
		return dailyOffer, helpers.ErrRepo
	}

	return dailyOffer, nil
}

func (r *Repos) DeleteDailyOffer(dailyOfferID int) error {
	res := r.db.Where("id = ?", dailyOfferID).Delete(&models.DailyOffers{})
	if res.Error != nil || res.RowsAffected == 0 {
		return helpers.ErrRepo
	}

	return nil
}

// Support functions
func (r *Repos) GetDailyOfferByField(field string, value any) (models.DailyOffers, error) {
	var dailyOffers models.DailyOffers

	if err := r.db.Where(fmt.Sprintf("%v = ?", field), value).First(&dailyOffers).Error; err != nil {
		return dailyOffers, helpers.ErrRepo
	}

	return dailyOffers, nil
}

func (r *Repos) DailyOfferExists(date string, restaurantID int) bool {
	res := r.db.Where("date = ? AND restaurants_id = ?", date, restaurantID).First(&models.DailyOffers{}).RowsAffected

	return res != 0
}
