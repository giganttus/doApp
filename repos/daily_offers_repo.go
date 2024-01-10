package repos

import (
	"doApp/helpers"
	"doApp/models"
	"fmt"
	"time"
)

// Main functions
func (r *Repos) CreateDailyOffer(input models.DailyOffers) error {
	if err := r.db.Create(&input).Error; err != nil {
		return helpers.ErrRepo
	}

	return nil
}

func (r *Repos) GetDailyOffers(currentDate string) ([]models.DailyOffers, error) {
	var dailyOffers []models.DailyOffers

	if res := r.db.Find(&dailyOffers); res.Error != nil {
		return nil, helpers.ErrRepo
	}

	return dailyOffers, nil
}

func (r *Repos) GetDailyOffer(restaurantID int, currentDate string) (models.DailyOffers, error) {
	var dailyOffer models.DailyOffers

	if err := r.db.Where("restaurants_id = ? AND date = ?", restaurantID, currentDate).First(&dailyOffer).Error; err != nil {
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

func (r *Repos) GetDailyOfferByFieldAndCurrentDate(field string, value any) (models.DailyOffers, error) {
	var dailyOffers models.DailyOffers
	currentDate := time.Now().Format("2006-01-02")

	if err := r.db.Where(fmt.Sprintf("%v = ? AND date = ? ", field), value, currentDate).First(&dailyOffers).Error; err != nil {
		return dailyOffers, helpers.ErrRepo
	}

	return dailyOffers, nil
}

func (r *Repos) DailyOfferExists(date string, restaurantID int) bool {
	res := r.db.Where("date = ? AND restaurants_id = ?", date, restaurantID).First(&models.DailyOffers{}).RowsAffected

	return res != 0
}
