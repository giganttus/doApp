package handlers

import (
	"doApp/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Register		 godoc
// @Summary      Create an User
// @Description  Create new user for app
// @Tags         Users
// @Accept       application/json
// @Produce      application/json
// @Param        Request body models.Users true "User credentials for registration"
// @Success      201  {object} helpers.ResponseBody
// @Failure      401  {object} helpers.ResponseBody
// @Failure      403  {object} helpers.ResponseBody
// @Failure      404  {object} helpers.ResponseBody
// @Failure      500  {object} helpers.ResponseBody
// @Router       /register [post]
func (h *Handlers) CreateUser(ctx *gin.Context) {
	var newUser models.Users

	if err := ctx.ShouldBind(&newUser); err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"response": "Bad Request"})
		return
	}

	if err := h.Services.CreateUser(ctx, newUser); err != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"response": err.Error()})
		return
	}

	ctx.IndentedJSON(http.StatusCreated, gin.H{"response": "success"})
}

// Login		 godoc
// @Summary      Login
// @Description  Login service
// @Tags         Users
// @Accept       application/json
// @Produce      application/json
// @Param        Request body models.LoginReq true "User credentials for login"
// @Success      200  {object} models.LoginRes
// @Failure      401  {object} helpers.ResponseBody
// @Failure      403  {object} helpers.ResponseBody
// @Failure      404  {object} helpers.ResponseBody
// @Failure      500  {object} helpers.ResponseBody
// @Router       /login [post]
func (h *Handlers) Login(ctx *gin.Context) {
	var loginReq models.LoginReq

	if err := ctx.ShouldBind(&loginReq); err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"response": "Bad Request"})
		return
	}

	res, err := h.Services.Login(ctx, loginReq)
	if err != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"response": err.Error()})
		return
	}

	ctx.IndentedJSON(http.StatusOK, gin.H{"response": res})
}

// GetUsers		 godoc
// @Summary      Get Users
// @Description  Get list of all users
// @Tags         Users
// @Param		 Authorization header string true "JWT Token for Auth"
// @Success      200  {array}  models.Users
// @Failure      401  {object} helpers.ResponseBody
// @Failure      403  {object} helpers.ResponseBody
// @Failure      404  {object} helpers.ResponseBody
// @Failure      500  {object} helpers.ResponseBody
// @Router       /users [get]
func (h *Handlers) GetUsers(ctx *gin.Context) {
	res, err := h.Services.GetUsers(ctx)
	if err != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"response": err.Error()})
		return
	}

	ctx.IndentedJSON(http.StatusOK, gin.H{"response": res})
}
