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
}
