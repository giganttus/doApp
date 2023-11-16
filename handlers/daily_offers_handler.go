package handlers

import (
	"doApp/helpers"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// CreateDailyOffer		godoc
// @Summary      		Create Daily Offer
// @Description  		Create new daily offer
// @Tags         		Daily Offer
// @Accept 				multipart/form-data
// @Produce      		application/json
// @Param		 		Authorization header string true "JWT Token for Auth"
// @Param 				file formData file true "this is a test file"
// @Success      		201  {object} helpers.ResponseBody
// @Failure      		401  {object} helpers.ResponseBody
// @Failure      		403  {object} helpers.ResponseBody
// @Failure      		404  {object} helpers.ResponseBody
// @Failure      		500  {object} helpers.ResponseBody
// @Router       		/dailyOffer [post]
func (h *Handlers) CreateDailyOffer(ctx *gin.Context) {
	err := h.Services.CreateDailyOffer(ctx)
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

// GetDailyOffers	godoc
// @Summary      	Get Daily Offers
// @Description  	Get list of all daily offers
// @Tags         	Daily Offer
// @Param		 	Authorization header string true "JWT Token for Auth"
// @Success      	200  {array}  models.DailyOffers
// @Failure      	401  {object} helpers.ResponseBody
// @Failure      	403  {object} helpers.ResponseBody
// @Failure      	404  {object} helpers.ResponseBody
// @Failure      	500  {object} helpers.ResponseBody
// @Router       	/dailyOffers [get]
func (h *Handlers) GetDailyOffers(ctx *gin.Context) {
	res, err := h.Services.GetDailyOffers(ctx)
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

// GetDailyOfferByUser	godoc
// @Summary      		Get Daily Offer
// @Description  		Get daily offer for logged user
// @Tags         		Daily Offer
// @Param		 		Authorization header string true "JWT Token for Auth"
// @Success      		200  {array}  models.DailyOffers
// @Failure      		401  {object} helpers.ResponseBody
// @Failure      		403  {object} helpers.ResponseBody
// @Failure      		404  {object} helpers.ResponseBody
// @Failure      		500  {object} helpers.ResponseBody
// @Router       		/dailyOffer [get]
func (h *Handlers) GetDailyOffer(ctx *gin.Context) {
	res, err := h.Services.GetDailyOffer(ctx)
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

// GetDailyOfferByRID	godoc
// @Summary      		Get Daily Offer
// @Description  		Get daily offer by restaurant ID
// @Tags         		Daily Offer
// @Param		 		Authorization header string true "JWT Token for Auth"
// @Success      		200  {array}  models.DailyOffers
// @Failure      		401  {object} helpers.ResponseBody
// @Failure      		403  {object} helpers.ResponseBody
// @Failure      		404  {object} helpers.ResponseBody
// @Failure      		500  {object} helpers.ResponseBody
// @Router       		/dailyOffer/{restaurantId} [get]
func (h *Handlers) GetDailyOfferByRestaurantID(ctx *gin.Context) {
	restaurantID, err := strconv.Atoi(ctx.Param("restaurantId"))
	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"response": "Bad Request"})
		return
	}

	res, err := h.Services.GetDailyOfferByRestaurantID(ctx, restaurantID)
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

// DeleteDailyOffer		godoc
// @Summary      		Delete Daily Offer
// @Description  		Delet daily offer by offer id
// @Tags         		Daily Offer
// @Param		 		Authorization header string true "JWT Token for Auth"
// @Success      		200  {array}  models.DailyOffers
// @Failure      		401  {object} helpers.ResponseBody
// @Failure      		403  {object} helpers.ResponseBody
// @Failure      		404  {object} helpers.ResponseBody
// @Failure      		500  {object} helpers.ResponseBody
// @Router       		/dailyOffer/{id} [delete]
func (h *Handlers) DeleteDailyOffer(ctx *gin.Context) {
	dailyOfferID, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"response": "Bad Request"})
		return
	}

	err = h.Services.DeleteDailyOffer(ctx, dailyOfferID)
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
