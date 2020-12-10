package http

import (
	"microservice_tut/bookstore_oauth_api/src/domain/access_token"

	"net/http"

	"github.com/gin-gonic/gin"
)

type AccessTokenHandler interface {
	GetByID(*gin.Context)
}

type accessTokenHandler struct {
	service access_token.Service
}

func NewHandler(service access_token.Service) AccessTokenHandler {
	return &accessTokenHandler{
		service: service,
	}
}

func (h *accessTokenHandler) GetByID(c *gin.Context) {
	token, err := h.service.GetByID(c.Param("token_id"))

	if err != nil {
		c.JSON(err.Status, err)

	}
	c.JSON(http.StatusNotImplemented, token)
}
