package hero

import (
	"encoding/json"
	"fmt"
	"indonesian-heroes/helper"
	"io/ioutil"
	"net/http"
	"strconv"
)

type IService interface {
	GetAll() ([]Hero, int, error)
	GetByAge(age string) ([]Hero, int, error)
}

type service struct {
	repoHero IRepository
}

const URL = "https://indonesia-public-static-api.vercel.app/api/heroes"

func NewService(repoHero IRepository) *service {
	return &service{
		repoHero: repoHero,
	}
}

func (s *service) GetAll() ([]Hero, int, error) {
	// cek apakah sudah ada data atau belum
	allHero, err := s.repoHero.FindAll()
	if err != nil {
		return allHero, 0, err
	}

	totalData := len(allHero)

	if totalData == 0 {
		response, err := http.Get(URL)
		if err != nil {
			return []Hero{}, totalData, err
		}

		respByte, err := ioutil.ReadAll(response.Body)
		if err != nil {
			return []Hero{}, totalData, err
		}

		var respHero HeroFromAPI
		if err := json.Unmarshal(respByte, &respHero.Data); err != nil {
			return []Hero{}, totalData, err
		}

		var hero Hero
		for _, data := range respHero.Data {
			birthyear := helper.ChangeFromInterface(data.Birthyear)
			deathYear := helper.ChangeFromInterface(data.DeathYear)

			age := ""

			ageCalculate := deathYear - birthyear
			if ageCalculate > 100 {
				age = "tidak diketahui"
			} else {
				age = strconv.Itoa(ageCalculate)
			}

			hero.Age = age
			hero.Name = data.Name
			hero.Description = data.Description
			hero.BirthYear = int32(birthyear)
			hero.DeathYear = int32(deathYear)
			hero.Description = data.Description

			s.repoHero.Save(hero)
		}
	}

	// get all data
	heros, err := s.repoHero.FindAll()
	if err != nil {
		return heros, totalData, err
	}

	return heros, totalData, nil
}

func (s *service) GetByAge(age string) ([]Hero, int, error) {
	heros, err := s.repoHero.FindByAge(age)
	if err != nil {
		return heros, 0, err
	}

	// cek apakah ada data
	totalData := len(heros)
	if totalData == 0 {
		return heros, 0, fmt.Errorf("data tidak ditemukan")
	}

	return heros, totalData, nil
}
