package pokeapi

// the following are slim representations of what the structs actually hold
// for the purpose of readablility and efficiency, only used data is parsed
// for full data see:
// curl "https://pokeapi.co/api/v2/location-area/canalave-city-area/"
type LocationArea struct {
    PokemonEncounters []PokemonEncounters `json:"pokemon_encounters"`
}

type PokemonEncounters struct {
    Pokemon Pokemon `json:"pokemon"`
}

type Pokemon struct {
    Name string `json:"name"`
}
