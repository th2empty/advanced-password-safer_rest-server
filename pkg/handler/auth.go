package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/th2empty/auth-server/pkg/models"
	"net/http"
)

func (h *Handler) SignUp(ctx *gin.Context) {
	var input models.User

	if err := ctx.BindJSON(&input); err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.Authorization.CreateUser(input)
	if err != nil {
		var errStruct = models.Error{
			Package: "handler",
			File:    "auth.go",
			Func:    "SignUp",
			Err:     err.Error(),
		}

		if errStruct.Message = err.Error(); err.Error() == "pq: duplicate key value violates unique constraint \"users_username_key\"" {
			errStruct.Message = models.UserAlreadyRegistered
		}

		if errStruct.StatusCode = http.StatusInternalServerError; errStruct.Message == models.UserAlreadyRegistered {
			errStruct.StatusCode = http.StatusConflict
		}

		newErrorResponseDetails(ctx, errStruct.StatusCode, errStruct)
		return
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

type signInInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (h *Handler) SignIn(ctx *gin.Context) {
	var input signInInput

	if err := ctx.BindJSON(&input); err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	token, err := h.services.Authorization.GenerateToken(input.Username, input.Password)
	if err != nil {
		var errStruct = models.Error{
			Package: "handler",
			File:    "auth.go",
			Func:    "SignIn",
			Err:     err.Error(),
		}

		if errStruct.Message = err.Error(); err.Error() == "sql: no rows in result set" {
			errStruct.Message = models.InvalidPasswordOrUsername
		}

		if errStruct.StatusCode = http.StatusInternalServerError; errStruct.Message == models.InvalidPasswordOrUsername {
			errStruct.StatusCode = http.StatusBadRequest
		}

		newErrorResponseDetails(ctx, errStruct.StatusCode, errStruct)
		return
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{
		"token": token,
	})
}

type deleteInput struct {
	Id int `json:"id" binding:"required"`
}
