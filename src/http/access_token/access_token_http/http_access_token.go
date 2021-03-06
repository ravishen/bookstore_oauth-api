package http

import (
	"net/http"
	"strings"

	"github.com/ravishen/bookstore_oauth-api/src/utils/errors"

	"github.com/gin-gonic/gin"
	"github.com/ravishen/bookstore_oauth-api/src/domain/access_token"
)

type AccessTokenHandler interface {
	GetById(*gin.Context)
	Create(*gin.Context)
}
type accessTokenHandler struct {
	service access_token.Service
}

func (handler *accessTokenHandler) GetById(c *gin.Context) {
	accessTokenId := strings.TrimSpace(c.Param("access_token_id"))
	access_token, err := handler.service.GetById(accessTokenId)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusOK, access_token)
}

func (handler *accessTokenHandler) Create(c *gin.Context) {
	var at access_token.AccessToken

	if err := c.ShouldBindJSON(&at); err != nil {
		restErr := errors.NewBadRequestError("invalid json body")
		c.JSON(http.StatusBadRequest, restErr)
		return
	}
	if err := handler.service.Create(at); err != nil {
		restErr := errors.NewBadRequestError(err.Error)
		c.JSON(http.StatusInternalServerError, restErr)
		return
	}
	c.JSON(http.StatusOK, at)
}

func NewHandler(service access_token.Service) AccessTokenHandler {
	return &accessTokenHandler{
		service: service,
	}
}
