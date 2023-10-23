package handlers

import (
	"doApp/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handlers) CreateRestaurant(ctx *gin.Context) {
	var newRestaurant models.Restaurants

	if err := ctx.ShouldBind(&newRestaurant); err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"response": "Bad Request"})
		return
	}

	err := h.Services.CreateRestaurants(ctx, newRestaurant)
	if err != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"response": err.Error()})
		return
	}

	ctx.IndentedJSON(http.StatusCreated, gin.H{"response": "success"})
}

func (h *Handlers) GetRestaurants(ctx *gin.Context) {
	res, err := h.Services.GetRestaurants(ctx)
	if err != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"response": err.Error()})
		return
	}

	ctx.IndentedJSON(http.StatusOK, gin.H{"response": res})
}

func (h *Handlers) UpdateRestaurant(ctx *gin.Context) {
	var upResta models.Restaurants

	if err := ctx.ShouldBind(&upResta); err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"response": "Bad Request"})
		return
	}

	err := h.Services.UpdateRestaurants(ctx, upResta)
	if err != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"response": err.Error()})
		return
	}

	ctx.IndentedJSON(http.StatusOK, gin.H{"response": "success"})
}

func (h *Handlers) DeleteRestaurant(ctx *gin.Context) {
	restaurantID, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"response": "Bad Request"})
		return
	}

	err = h.Services.DeleteRestaurant(ctx, restaurantID)
	if err != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"response": err.Error()})
		return
	}

	ctx.IndentedJSON(http.StatusOK, gin.H{"response": "success"})
}
