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
	heros, totalData, err := h.heroService.GetAll()
	if err != nil {
		response := helper.ResponseAPIFormat(
			"failed",
			err.Error(),
			http.StatusInternalServerError,
			totalData,
			heros,
		)

		c.JSON(http.StatusInternalServerError, response)
		return
	}

	response := helper.ResponseAPIFormat(
		"success",
		"success get all data",
		http.StatusOK,
		totalData,
		heros,
	)

	c.JSON(http.StatusOK, response)
}

func (h *heroHandler) GetHerosByAge(c *gin.Context) {
	age := c.Param("age")
	heros, totalData, err := h.heroService.GetByAge(age)
	if err != nil {
		response := helper.ResponseAPIFormat(
			"failed",
			err.Error(),
			http.StatusInternalServerError,
			totalData,
			heros,
		)

		c.JSON(http.StatusInternalServerError, response)
		return
	}

	response := helper.ResponseAPIFormat(
		"success",
		"success get data",
		http.StatusOK,
		totalData,
		heros,
	)

	c.JSON(http.StatusOK, response)
}
