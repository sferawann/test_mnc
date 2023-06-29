package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sferawann/test_mnc/model"
	"github.com/sferawann/test_mnc/usecase"
)

type AccountCon struct {
	AccountUsecase usecase.AccountUsecase
}

func NewAccountController(AccountUsecase usecase.AccountUsecase) *AccountCon {
	return &AccountCon{
		AccountUsecase: AccountUsecase,
	}
}

func (c *AccountCon) Create(ctx *gin.Context) {
	insertAccount := model.Account{}
	currentUserID, exists := ctx.Get("currentUserID")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
		return
	}

	insertAccount.UserID = currentUserID.(int64)

	if err := ctx.ShouldBind(&insertAccount); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if insertAccount.UserID == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "id_user is required"})
		return
	}

	if insertAccount.Balance <= 0 {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "balance must be greater than 0"})
		return
	}

	newAccount, err := c.AccountUsecase.Save(insertAccount)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"Account": newAccount})
}

func (c *AccountCon) FindAll(ctx *gin.Context) {
	Accounts, err := c.AccountUsecase.FindAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"Accounts": Accounts})
}

func (c *AccountCon) FindByID(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	Account, err := c.AccountUsecase.FindById(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"Account": Account})
}

func (c *AccountCon) Update(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updateID := model.Account{ID: id}
	if err := ctx.ShouldBindJSON(&updateID); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedAccount, err := c.AccountUsecase.Update(updateID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"Account": updatedAccount})
}

func (c *AccountCon) Delete(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	_, err = c.AccountUsecase.Delete(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Successfully deleted Account!"})
}

func (c *AccountCon) GetByUserID(ctx *gin.Context) {
	currentUserID, exists := ctx.Get("currentUserID")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
		return
	}

	AccountByUserID, err := c.AccountUsecase.FindByUserId(currentUserID.(int64))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"Account": AccountByUserID})
}
