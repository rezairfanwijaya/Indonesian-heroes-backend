package hero

type HeroFromAPI struct {
	Data []heroResponse
}

type heroResponse struct {
	Name        string      `json:"name"`
	Birthyear   interface{} `json:"birth_year"`
	DeathYear   interface{} `json:"death_year"`
	Description string      `json:"description"`
}
