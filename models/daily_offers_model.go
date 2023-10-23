package models

type DailyOffers struct {
	ID            int    `json:"id" gorm:"primaryKey"`
	Link          string `json:"link" binding:"max=100"`
	Date          string `json:"date" binding:"max=100"`
	RestaurantsID int    `json:"restaurants_id"`
}
