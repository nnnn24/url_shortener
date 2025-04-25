package handlers

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nnnn24/url_shortener_service/internal/models"
	"github.com/nnnn24/url_shortener_service/internal/service"
	"gorm.io/gorm"
)

type URLHandler struct {
	urlService *service.URLService
}

func NewURLHandler(urlService *service.URLService) *URLHandler {
	return &URLHandler{
		urlService: urlService,
	}
}

func (handler *URLHandler) CreateURL(c *gin.Context) {
	var req models.CreateURLRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Bad request",
			"details": err.Error(),
		})
	}

	resp, err := handler.urlService.CreateShortURL(c.Request.Context(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Failed to create URL",
			"details": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, resp)
}

func (handler *URLHandler) FindByShortCode(c *gin.Context) {
	shortCode, isExist := c.Params.Get("shortCode")

	if !isExist {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "shortCode is required !!!",
		})
		return
	}

	resp, err := handler.urlService.FindURL(c.Request.Context(), shortCode)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Failed to get URL",
			"details": err,
		})
		return
	}

	c.JSON(http.StatusOK, resp)
}

func (handler *URLHandler) UpdateURL(c *gin.Context) {
	shortCode, isExist := c.Params.Get("shortCode")

	if !isExist {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "shortCode is required !!!",
		})
		return
	}

	var req models.UpdateURLRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Bad request url is required",
			"details": err,
		})
		return
	}

	resp, err := handler.urlService.UpdateURL(c.Request.Context(), &req, shortCode)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error",
			"details": err,
		})
		return
	}

	c.JSON(http.StatusOK, resp)
}

func (handler *URLHandler) DeleteURL(c *gin.Context) {
	shortCode, isExist := c.Params.Get("shortCode")

	if !isExist {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "shortCode is required",
		})
	}

	err := handler.urlService.DeleteURL(c.Request.Context(), shortCode)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(
				http.StatusNotFound, gin.H{
					"message": "URL not found !!!",
					"details": err,
				})

			return
		}

		c.JSON(
			http.StatusInternalServerError, gin.H{
				"message": "Error",
				"details": err,
			},
		)

		return
	}

	c.JSON(
		http.StatusOK, gin.H{
			"message": "Deleted",
		},
	)
}
