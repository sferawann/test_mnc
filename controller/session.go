package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sferawann/test_mnc/model"
	"github.com/sferawann/test_mnc/usecase"
)

type SessionCon struct {
	SessionUsecase usecase.SessionUsecase
}

func NewSessionController(SessionUsecase usecase.SessionUsecase) *SessionCon {
	return &SessionCon{SessionUsecase: SessionUsecase}
}

func (c *SessionCon) Create(ctx *gin.Context) {
	insertSession := model.Session{}
	if err := ctx.ShouldBind(&insertSession); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if insertSession.UserID == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "id_user is required"})
		return
	}

	if insertSession.Token == "" {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "token is required"})
		return
	}

	newSession, err := c.SessionUsecase.Save(insertSession)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"Session": newSession})
}

func (c *SessionCon) Update(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updateID := model.Session{ID: id}
	if err := ctx.ShouldBindJSON(&updateID); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedSession, err := c.SessionUsecase.Update(updateID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"Session": updatedSession})
}

func (c *SessionCon) Delete(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	_, err = c.SessionUsecase.Delete(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Successfully deleted Session!"})
}

func (c *SessionCon) FindAll(ctx *gin.Context) {
	Sessions, err := c.SessionUsecase.FindAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"Sessions": Sessions})
}

func (c *SessionCon) FindByID(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	Session, err := c.SessionUsecase.FindById(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"Session": Session})
}
