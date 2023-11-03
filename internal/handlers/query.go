package handlers

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"openai/internal/domain"
)

type ErrResp struct {
	Message string `json:"error"`
}

// sendSingleQuery обработчик для одного вопроса
func (h *Handler) sendSingleQuery(ctx *gin.Context) {
	var q domain.SingleQuery

	if err := ctx.BindJSON(&q); err != nil {
		log.Printf("Некорректное тело запроса %v", err)

		ctx.AbortWithStatusJSON(http.StatusBadRequest, ErrResp{Message: "incorrect request body"})
		return
	}

	resp, err := h.service.Responder.ResponseSingle(q)
	if err != nil {
		log.Printf("Ошибка в одиночном вопросе - %v", err)

		ctx.AbortWithStatusJSON(http.StatusInternalServerError, ErrResp{Message: err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, resp)
}
