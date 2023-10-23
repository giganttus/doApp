package handlers

import (
	"doApp/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handlers) CreateDailyOffer(ctx *gin.Context) {
	var newDailyOffer models.DailyOffers

	if err := ctx.ShouldBind(&newDailyOffer); err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"response": "Bad Request"})
		return
	}

	err := h.Services.CreateDailyOffer(ctx, newDailyOffer)
	if err != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"response": err.Error()})
		return
	}

	ctx.IndentedJSON(http.StatusCreated, gin.H{"response": "success"})
}

func (h *Handlers) GetDailyOffers(ctx *gin.Context) {
	res, err := h.Services.GetDailyOffers(ctx)
	if err != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"response": err.Error()})
		return
	}

	ctx.IndentedJSON(http.StatusOK, gin.H{"response": res})
}

func (h *Handlers) GetDailyOffer(ctx *gin.Context) {
	res, err := h.Services.GetDailyOffer(ctx)
	if err != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"response": err.Error()})
		return
	}

	ctx.IndentedJSON(http.StatusOK, gin.H{"response": res})
}

func (h *Handlers) GetDailyOfferByRestaurantID(ctx *gin.Context) {
	restaurantID, err := strconv.Atoi(ctx.Param("restaurantId"))
	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"response": "Bad Request"})
		return
	}

	res, err := h.Services.GetDailyOfferByRestaurantID(ctx, restaurantID)
	if err != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"response": err.Error()})
		return
	}

	ctx.IndentedJSON(http.StatusOK, gin.H{"response": res})
}

func (h *Handlers) DeleteDailyOffer(ctx *gin.Context) {
	dailyOfferID, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"response": "Bad Request"})
		return
	}

	err = h.Services.DeleteDailyOffer(ctx, dailyOfferID)
	if err != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"response": err.Error()})
		return
	}

	ctx.IndentedJSON(http.StatusOK, gin.H{"response": "success"})
}
