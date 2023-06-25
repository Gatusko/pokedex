package internal

type ExplorePokemon struct {
	EncounterMethodRates `json:"encounter_method_rates"`
	GameIndex            int      `json:"game_index"`
	ID                   int      `json:"id"`
	Location             location `json:"location"`
	Name                 string   `json:"name"`
	Names                `json:"names"`
	PokemonEncounters    `json:"pokemon_encounters"`
}

type EncounterMethodRates []struct {
	EncounterMethod         encounterMethod         `json:"encounter_method"`
	VersionDetailsEncounter versionDetailsEncounter `json:"version_details"`
}

type encounterMethod struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type versionDetailsEncounter []struct {
	Rate    int `json:"rate"`
	Version struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"version"`
}

type location struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type Language struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type Names []struct {
	Language `json:"language"`
	Name     string `json:"name"`
}

type Version struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type EncounterDetails []struct {
	Chance          int           `json:"chance"`
	ConditionValues []interface{} `json:"condition_values"`
	MaxLevel        int           `json:"max_level"`
	Method          `json:"method"`
	MinLevel        int `json:"min_level"`
}

type Method struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type VersionDetails []struct {
	EncounterDetails `json:"encounter_details"`
	MaxChance        int `json:"max_chance"`
	Version          `json:"version"`
}

type Pokemon struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type PokemonEncounters []struct {
	Pokemon        `json:"pokemon"`
	VersionDetails `json:"version_details"`
}
