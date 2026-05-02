 package pokeapi

 // Get full representation of the pokemon struct. Different call then Pokemon in encounters, hence 
 // different struct. All sublayers are removed for the sake of simplicity.
 // for full struct see:
 // curl "https://pokeapi.co/api/v2/pokemon/clefairy/"
 
type FullPokemon struct {
	BaseExperience         int             `json:"base_experience"`
	Height                 int             `json:"height"`
	ID                     int             `json:"id"`
	LocationAreaEncounters string          `json:"location_area_encounters"`
	Name                   string          `json:"name"`
	Order                  int             `json:"order"`
	Weight                 int             `json:"weight"`
	Stats                  []Stats         `json:"stats"`
	Types                  []Types         `json:"types"`
}

type Stat struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}
type Stats struct {
	BaseStat int  `json:"base_stat"`
	Effort   int  `json:"effort"`
	Stat     Stat `json:"stat"`
}
type Type struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}
type Types struct {
	Slot int  `json:"slot"`
	Type Type `json:"type"`
}
