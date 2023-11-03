package services

import (
	"doApp/auth"
	"doApp/helpers"
	"doApp/models"

	"github.com/gin-gonic/gin"
)

func (s *Services) CreateUser(ctx *gin.Context, input models.Users) error {
	if emailEx := s.Repos.EmailExists(input.Email); emailEx {
		return helpers.ErrService
	}

	hashedPwd, err := models.HashPassword(input.Password)
	if err != nil {
		return helpers.ErrService
	}

	input.Password = hashedPwd

	return s.Repos.CreateUser(input)
}

func (s *Services) Login(ctx *gin.Context, input models.LoginReq) (models.LoginRes, error) {
	var loginRes models.LoginRes

	user, err := s.Repos.GetUserByField("email", input.Email)
	if err != nil {
		return loginRes, helpers.ErrService
	}

	if err = models.ComparePassword(user.Password, input.Password); err != nil {
		return loginRes, helpers.ErrService
	}

	token, err := auth.GenerateJwt(user.ID)
	if err != nil {
		return loginRes, err
	}
	loginRes.Token = token["access_token"]
	loginRes.RolesID = user.RolesID
	return loginRes, nil
}

func (s *Services) GetUsers(ctx *gin.Context) ([]models.Users, error) {
	userID := ctx.Keys["userID"].(int)

	if aA := s.Repos.AllowAccess(userID, "admin"); !aA {
		return nil, helpers.ErrService
	}

	return s.Repos.GetUsers()
}
