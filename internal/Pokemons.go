package internal

type PokemonData struct {
	Name           string  `json:"name"`
	BaseExperience float64 `json:"base_experience"`
	Height         int     `json:"height"`
	Stats          []Stats `json:"stats"`
	Weight         int     `json:"weight"`
	Types          []Types `json:"types"`
}
type Stats struct {
	BaseStat int `json:"base_stat"`
	Effort   int `json:"effort"`
	Stat     `json:"stat"`
}

type Stat struct {
	Name string `json:"name"`
}

type Types struct {
	Slot int `json:"slot"`
	Type `json:"type"`
}

type Type struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}
