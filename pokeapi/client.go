package pokeapi

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type Client struct{}

type Pokemon struct {
	Abilities []struct {
		Ability struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"ability"`
		IsHidden bool `json:"is_hidden"`
		Slot     int  `json:"slot"`
	} `json:"abilities"`
	BaseExperience int `json:"base_experience"`
	Forms          []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"forms"`
	GameIndices []struct {
		GameIndex int `json:"game_index"`
		Version   struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"version"`
	} `json:"game_indices"`
	HeldItems []struct {
		Item []struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"item"`
		VersionDetails []struct {
			Rarity  int `json:"rarity"`
			Version []struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	}
	ID                     int    `json:"id"`
	IsDefault              bool   `json:"is_default"`
	LocationAreaEncounters string `json:"location_area_encounters"`
	Moves                  []struct {
		Move struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"moves"`
		VersionGroupDetails []struct {
			LevelLearnedAt  int `json:"level_learned_at"`
			MoveLearnMethod struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"move_learn_method"`
			VersionGroup struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version_group"`
		} `json:"version_group_details"`
	} `json:"moves"`
	Name      string     `json:"name"`
	Order     int        `json:"order"`
	PastTypes []struct{} `json:"past_types"`
	Species   struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"species"`
	Weight int `json:"weight"`
}

func (c Client) Ditto() (Pokemon, error) {
	resp, err := http.Get("https://pokeapi.co/api/v2/pokemon/ditto")
	if err != nil {
		return Pokemon{}, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return Pokemon{}, nil
	}

	pokemon := &Pokemon{}
	err = json.Unmarshal(body, pokemon)
	if err != nil {
		return Pokemon{}, err
	}

	return *pokemon, nil
}
