package hero

import (
	"encoding/json"
	"indonesian-heroes/helper"
	"io/ioutil"
	"net/http"
	"strconv"
)

type IService interface {
	GetAll() ([]Hero, error)
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

func (s *service) GetAll() ([]Hero, error) {
	// cek apakah sudah ada data atau belum
	allHero, err := s.repoHero.FindAll()
	if err != nil {
		return allHero, err
	}

	if len(allHero) == 0 {
		response, err := http.Get(URL)
		if err != nil {
			return []Hero{}, err
		}

		respByte, err := ioutil.ReadAll(response.Body)
		if err != nil {
			return []Hero{}, err
		}

		var respHero HeroFromAPI
		if err := json.Unmarshal(respByte, &respHero.Data); err != nil {
			return []Hero{}, err
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
		return heros, err
	}

	return heros, nil
}
