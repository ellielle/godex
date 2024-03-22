package pokedex

import (
	"sync"
)

type Pokemon struct {
	BaseExperience int    `json:"base_experience"`
	Name           string `json:"name"`
	Height         int    `json:"height"`
	Weight         int    `json:"weight"`
	Stats          []struct {
		BaseStat int `json:"base_stat"`
		Effort   int `json:"effort"`
		Stat     struct {
			Name string `json:"name"`
			Url  string `json:"url"`
		} `json:"stat"`
	} `json:"stats"`
	Types []struct {
		Slot int `json:"slot"`
		Type struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"type"`
	} `json:"types"`
}

type Pokedex struct {
	Entries map[string]Pokemon
	mu      *sync.RWMutex
}

func NewPokedex() Pokedex {
	return Pokedex{
		Entries: make(map[string]Pokemon),
		mu:      &sync.RWMutex{},
	}
}

func (p *Pokedex) AddPokemon(pokemon Pokemon) error {
	p.mu.Lock()
	defer p.mu.Unlock()

	p.Entries[pokemon.Name] = pokemon
	return nil
}
