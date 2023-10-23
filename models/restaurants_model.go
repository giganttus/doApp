package models

type Restaurants struct {
	ID      int     `json:"id" gorm:"primaryKey"`
	Name    string  `json:"name" binding:"max=100"`
	Address string  `json:"address" binding:"max=100"`
	Glink   string  `json:"glink" binding:"max=100"`
	Lat     float64 `json:"lat" binding:"max=100"`
	Lon     float64 `json:"lon" binding:"max=100"`
	UsersID int     `json:"users_id" gorm:"default:null"`
}
