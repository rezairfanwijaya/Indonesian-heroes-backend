package handler

import (
	"indonesian-heroes/helper"
	"indonesian-heroes/hero"
	"net/http"

	"github.com/gin-gonic/gin"
)

type heroHandler struct {
	heroService hero.IService
}

func NewHeroHandler(heroService hero.IService) *heroHandler {
	return &heroHandler{
		heroService: heroService,
	}
}

func (h *heroHandler) GetAllHero(c *gin.Context) {
	heros, err := h.heroService.GetAll()
	if err != nil {
		response := helper.ResponseAPIFormat(
			"failed",
			err.Error(),
			http.StatusInternalServerError,
			heros,
		)

		c.JSON(http.StatusInternalServerError, response)
		return
	}

	response := helper.ResponseAPIFormat(
		"success",
		"success get all data",
		http.StatusOK,
		heros,
	)

	c.JSON(http.StatusOK, response)
}
