package handlers

import (
	"doApp/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

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

func (h *Handlers) GetUsers(ctx *gin.Context) {
	res, err := h.Services.GetUsers(ctx)
	if err != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"response": err.Error()})
		return
	}

	ctx.IndentedJSON(http.StatusOK, gin.H{"response": res})
}
