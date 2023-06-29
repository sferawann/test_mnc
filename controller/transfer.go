package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sferawann/test_mnc/model"
	"github.com/sferawann/test_mnc/usecase"
)

type TransferCon struct {
	TransferUsecase usecase.TransferUsecase
	AccountUsecase  usecase.AccountUsecase
}

func NewTransferController(TransferUsecase usecase.TransferUsecase, AccountUsecase usecase.AccountUsecase) *TransferCon {
	return &TransferCon{
		TransferUsecase: TransferUsecase,
		AccountUsecase:  AccountUsecase,
	}
}

func (c *TransferCon) Create(ctx *gin.Context) {
	insertTransfer := model.Transfer{}
	if err := ctx.ShouldBind(&insertTransfer); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if insertTransfer.FromAccountID == 0 {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "from_account_id is required"})
		return
	}
	if insertTransfer.ToAccountID == 0 {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "to_account_id is required"})
		return
	}
	if insertTransfer.Amount <= 0 {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Amount must be greater than 0"})
		return
	}

	_, err := c.AccountUsecase.FindById(insertTransfer.FromAccountID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "from_account_id not found"})
		return
	}

	_, err = c.AccountUsecase.FindById(insertTransfer.ToAccountID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "to_account_id not found"})
		return
	}

	newTransfer, err := c.TransferUsecase.Save(insertTransfer)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"Transfer": newTransfer})
}

func (c *TransferCon) Update(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updateID := model.Transfer{ID: id}
	if err := ctx.ShouldBindJSON(&updateID); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedTransfer, err := c.TransferUsecase.Update(updateID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"Transfer": updatedTransfer})
}

func (c *TransferCon) Delete(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	_, err = c.TransferUsecase.Delete(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Successfully deleted Transfer!"})
}

func (c *TransferCon) FindAll(ctx *gin.Context) {
	Transfers, err := c.TransferUsecase.FindAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"Transfers": Transfers})
}

func (c *TransferCon) FindByID(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	Transfer, err := c.TransferUsecase.FindById(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"Transfer": Transfer})
}
