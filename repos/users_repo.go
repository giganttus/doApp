package repos

import (
	"doApp/helpers"
	"doApp/models"
	"fmt"
)

// Main functions
func (r *Repos) CreateUser(input models.Users) error {
	if err := r.db.Omit("roles_id").Create(&input).Error; err != nil {
		return helpers.ErrRepo
	}

	return nil
}

func (r *Repos) GetUsers() ([]models.Users, error) {
	var users []models.Users

	if res := r.db.Find(&users); res.Error != nil {
		return nil, helpers.ErrRepo
	}

	return users, nil
}

// Support functions
func (r *Repos) GetUserByField(field string, value any) (models.Users, error) {
	var user models.Users

	if err := r.db.Where(fmt.Sprintf("%v = ?", field), value).First(&user).Error; err != nil {
		return user, helpers.ErrRepo
	}

	return user, nil
}

func (r *Repos) EmailExists(email string) bool {
	res := r.db.Where("email = ?", email).First(&models.Users{}).RowsAffected

	return res != 0
}

func (r *Repos) AllowAccess(userID int, roleTitle string) bool {
	var user models.Users

	if err := r.db.Where("id = ?", userID).First(&user).Error; err != nil {
		return false
	}

	if res := r.db.Find(&models.Roles{}, "id = ? AND title = ?", user.RolesID, roleTitle).RowsAffected; res != 1 {
		return false
	}

	return true
}
