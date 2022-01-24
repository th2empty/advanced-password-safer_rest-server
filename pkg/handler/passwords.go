package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/th2empty/auth-server/pkg/models"
	"net/http"
	"strconv"
)

func (h *Handler) AddPassword(ctx *gin.Context) {
	userId, err := getUserId(ctx)
	if err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, "user id not found")
		return
	}

	listId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, "invalid list id param")
		return
	}

	var input models.PasswordItem
	if err := ctx.BindJSON(&input); err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.PasswordItem.Add(userId, listId, input)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) GetAllPasswords(ctx *gin.Context) {
	userId, err := getUserId(ctx)
	if err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, "user id not found")
		return
	}

	listId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, "invalid list id param")
		return
	}

	items, err := h.services.PasswordItem.GetAll(userId, listId)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, items)
}

func (h *Handler) GetPassword(ctx *gin.Context) {
	userId, err := getUserId(ctx)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"package":    "handler",
			"file":       "passwords.go",
			"function":   "GetPassword",
			"code_block": 1,
			"error":      err,
		}).Error()
		newErrorResponse(ctx, http.StatusBadRequest, "user id not found")
		return
	}

	passwordId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"package":    "handler",
			"file":       "passwords.go",
			"function":   "GetPassword",
			"code_block": 2,
			"error":      err,
		}).Error()
		newErrorResponse(ctx, http.StatusBadRequest, "invalid password id param")
		return
	}

	password, err := h.services.PasswordItem.GetById(userId, passwordId)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"package":    "handler",
			"file":       "passwords.go",
			"function":   "GetPassword",
			"code_block": 3,
			"error":      err,
		}).Error()
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	logrus.WithFields(logrus.Fields{
		"package":  "handler",
		"file":     "passwords.go",
		"function": "GetPassword",
		"message":  "password received successfully",
	}).Info()
	ctx.JSON(http.StatusOK, password)
}

func (h *Handler) UpdatePassword(ctx *gin.Context) {
	userId, err := getUserId(ctx)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"package":  "handler",
			"file":     "passwords.go",
			"function": "UpdatePassword",
			"message":  err,
		}).Errorf("error while getting user id")
		newErrorResponse(ctx, http.StatusBadRequest, "user id not found")
		return
	}

	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"package":  "handler",
			"file":     "passwords.go",
			"function": "UpdatePassword",
			"message":  err,
		}).Errorf("error while converting 'id' param to int")
		newErrorResponse(ctx, http.StatusBadRequest, "invalid password id param")
		return
	}

	var input models.UpdatePasswordInput
	if err := ctx.BindJSON(&input); err != nil {
		logrus.WithFields(logrus.Fields{
			"package":  "handler",
			"file":     "passwords.go",
			"function": "UpdatePassword",
			"message":  err,
		}).Errorf("error while input binding")
		newErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.services.PasswordItem.Update(userId, id, input); err != nil {
		logrus.WithFields(logrus.Fields{
			"package":  "handler",
			"file":     "passwords.go",
			"function": "UpdatePassword",
			"message":  err,
		}).Errorf("error while input binding")
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, statusResponse{"Ok"})
}

func (h *Handler) DeletePassword(ctx *gin.Context) {
	userId, err := getUserId(ctx)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"package":    "handler",
			"file":       "passwords.go",
			"function":   "DeletePassword",
			"code_block": 1,
			"error":      err,
		}).Error()
		newErrorResponse(ctx, http.StatusBadRequest, "user id not found")
		return
	}

	passwordId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"package":    "handler",
			"file":       "passwords.go",
			"function":   "DeletePassword",
			"code_block": 2,
			"error":      err,
		}).Error()
		newErrorResponse(ctx, http.StatusBadRequest, "invalid password id param")
		return
	}

	err = h.services.PasswordItem.Delete(userId, passwordId)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"package":    "handler",
			"file":       "passwords.go",
			"function":   "DeletePassword",
			"code_block": 3,
			"error":      err,
		}).Error()
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	logrus.WithFields(logrus.Fields{
		"package":  "handler",
		"file":     "passwords.go",
		"function": "DeletePassword",
		"message":  "password deleted successfully",
	}).Info()
	ctx.JSON(http.StatusOK, statusResponse{"Ok"})
}
