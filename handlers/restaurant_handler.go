package handlers

import (
	"doApp/helpers"
	"doApp/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// CreateRestaurant	godoc
// @Summary      		Create Restaurant
// @Description  		Create new restaurant
// @Tags         		Restaurant
// @Accept       		application/json
// @Produce      		application/json
// @Param		 		Authorization header string true "JWT Token for Auth"
// @Param        		Request body models.Restaurants true "Restaurant info for creation"
// @Success      		201  {object} helpers.ResponseBody
// @Failure      		401  {object} helpers.ResponseBody
// @Failure      		403  {object} helpers.ResponseBody
// @Failure      		404  {object} helpers.ResponseBody
// @Failure      		500  {object} helpers.ResponseBody
// @Router       		/restaurant [post]
func (h *Handlers) CreateRestaurant(ctx *gin.Context) {
	var newRestaurant models.Restaurants

	if err := ctx.ShouldBind(&newRestaurant); err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"response": "Bad Request"})
		return
	}

	err := h.Services.CreateRestaurants(ctx, newRestaurant)
	if err != nil {
		if err.Error() == helpers.ErrForbidden.Error() {
			ctx.IndentedJSON(http.StatusForbidden, gin.H{"response": err.Error()})
			return
		}

		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"response": err.Error()})
		return
	}

	ctx.IndentedJSON(http.StatusCreated, gin.H{"response": "success"})
}

// GetRestaurant	godoc
// @Summary      	Get Restaurants
// @Description  	Get list of all restaurants
// @Tags         	Restaurant
// @Param		 	Authorization header string true "JWT Token for Auth"
// @Success      	200  {array}  models.Users
// @Failure      	401  {object} helpers.ResponseBody
// @Failure      	403  {object} helpers.ResponseBody
// @Failure      	404  {object} helpers.ResponseBody
// @Failure      	500  {object} helpers.ResponseBody
// @Router       	/restaurant [get]
func (h *Handlers) GetRestaurants(ctx *gin.Context) {
	res, err := h.Services.GetRestaurants(ctx)
	if err != nil {
		if err.Error() == helpers.ErrForbidden.Error() {
			ctx.IndentedJSON(http.StatusForbidden, gin.H{"response": err.Error()})
			return
		}

		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"response": err.Error()})
		return
	}

	ctx.IndentedJSON(http.StatusOK, gin.H{"response": res})
}

// GetRestaurants	godoc
// @Summary      	Update Restaurant
// @Description  	Update existing restaurant
// @Tags         	Restaurant
// @Param		 	Authorization header string true "JWT Token for Auth"
// @Param        	Request body models.Restaurants true "Restaurant info for update"
// @Success      	200  {array}  models.Users
// @Failure      	401  {object} helpers.ResponseBody
// @Failure      	403  {object} helpers.ResponseBody
// @Failure      	404  {object} helpers.ResponseBody
// @Failure      	500  {object} helpers.ResponseBody
// @Router       	/restaurant [put]
func (h *Handlers) UpdateRestaurant(ctx *gin.Context) {
	var upResta models.Restaurants

	if err := ctx.ShouldBind(&upResta); err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"response": "Bad Request"})
		return
	}

	err := h.Services.UpdateRestaurant(ctx, upResta)
	if err != nil {
		if err.Error() == helpers.ErrForbidden.Error() {
			ctx.IndentedJSON(http.StatusForbidden, gin.H{"response": err.Error()})
			return
		}

		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"response": err.Error()})
		return
	}

	ctx.IndentedJSON(http.StatusOK, gin.H{"response": "success"})
}

// DeletRestaurant	godoc
// @Summary      	Delete Restaurant
// @Description  	Delete existing restuarant by id
// @Tags         	Restaurant
// @Param		 	Authorization header string true "JWT Token for Auth"
// @Success      	200  {array}  models.Users
// @Failure      	401  {object} helpers.ResponseBody
// @Failure      	403  {object} helpers.ResponseBody
// @Failure      	404  {object} helpers.ResponseBody
// @Failure      	500  {object} helpers.ResponseBody
// @Router       	/restaurant/{id} [delete]
func (h *Handlers) DeleteRestaurant(ctx *gin.Context) {
	restaurantID, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"response": "Bad Request"})
		return
	}

	err = h.Services.DeleteRestaurant(ctx, restaurantID)
	if err != nil {
		if err.Error() == helpers.ErrForbidden.Error() {
			ctx.IndentedJSON(http.StatusForbidden, gin.H{"response": err.Error()})
			return
		}

		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"response": err.Error()})
		return
	}

	ctx.IndentedJSON(http.StatusOK, gin.H{"response": "success"})
}
