package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sferawann/test_mnc/model"
	"github.com/sferawann/test_mnc/usecase"
)

type HistoryCon struct {
	HistoryUsecase usecase.HistoryUsecase
}

func NewHistoryController(HistoryUsecase usecase.HistoryUsecase) *HistoryCon {
	return &HistoryCon{HistoryUsecase: HistoryUsecase}
}

func (c *HistoryCon) Create(ctx *gin.Context) {
	insertHistory := model.History{}
	if err := ctx.ShouldBind(&insertHistory); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if insertHistory.AccountID == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "id_account is required"})
		return
	}

	if insertHistory.Amount == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "amount is required"})
		return
	}

	newHistory, err := c.HistoryUsecase.Save(insertHistory)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"History": newHistory})
}

func (c *HistoryCon) Update(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updateID := model.History{ID: id}
	if err := ctx.ShouldBindJSON(&updateID); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedHistory, err := c.HistoryUsecase.Update(updateID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"History": updatedHistory})
}

func (c *HistoryCon) Delete(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	_, err = c.HistoryUsecase.Delete(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Successfully deleted History!"})
}

func (c *HistoryCon) FindAll(ctx *gin.Context) {
	Historys, err := c.HistoryUsecase.FindAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"Historys": Historys})
}

func (c *HistoryCon) FindByID(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	History, err := c.HistoryUsecase.FindById(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"History": History})
}
