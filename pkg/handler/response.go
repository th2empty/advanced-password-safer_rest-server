package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/th2empty/auth-server/pkg/models"
)

type errorResponse struct {
	Message string `json:"message"`
}

type statusResponse struct {
	Status string `json:"status"`
}

func newErrorResponse(ctx *gin.Context, statusCode int, message string) {
	logrus.Error(message)
	ctx.AbortWithStatusJSON(statusCode, errorResponse{message})
}

func newErrorResponseDetails(ctx *gin.Context, statusCode int, msg models.Error) {
	logrus.WithFields(logrus.Fields{
		"package": msg.Package,
		"file":    msg.File,
		"func":    msg.Func,
		"message": msg.Err,
	}).Errorf(msg.Message)
	ctx.AbortWithStatusJSON(statusCode, errorResponse{msg.Message})
}
