package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) DeleteUser(ctx *gin.Context) {
	_, ok := ctx.Get(userCtx)
	if !ok {
		newErrorResponse(ctx, http.StatusInternalServerError, "user id not found")
		return
	}
}
