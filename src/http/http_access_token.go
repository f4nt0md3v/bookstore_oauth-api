package http

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/f4nt0md3v/bookstore_oauth-api/src/domain/access_token"
	"github.com/f4nt0md3v/bookstore_oauth-api/src/utils/errors"
)

func NewHandler(service access_token.Service) AccessTokenHandler {
	return &accessTokenHandler{service: service}
}

type AccessTokenHandler interface {
	GetById(*gin.Context)
	Create(*gin.Context)
	UpdateExpirationTime(c *gin.Context)
}

type accessTokenHandler struct {
	service access_token.Service
}

func (handler *accessTokenHandler) GetById(c *gin.Context) {
	token, err := handler.service.GetById(c.Param("access_token_id"))
	if err != nil {
		c.JSON(err.StatusCode, err)
		return
	}
	c.JSON(http.StatusOK, token)
}

func (handler *accessTokenHandler) Create(c *gin.Context) {
	var token access_token.AccessToken
	if err := c.ShouldBindJSON(&token); err != nil {
		restErr := errors.NewBadRequestError("invalid json body")
		c.JSON(restErr.StatusCode, restErr)
		return
	}
	if err := handler.service.Create(token); err != nil {
		c.JSON(err.StatusCode, err)
		return
	}
	c.JSON(http.StatusCreated, token)
}

func (handler *accessTokenHandler) UpdateExpirationTime(c *gin.Context) {
	c.String(http.StatusNotImplemented, "not implemented yet")
}
