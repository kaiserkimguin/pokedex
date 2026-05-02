package pokeapi

import(
	"encoding/json"
	"io"
	"net/http"
	"errors"
)

func (c *Client) GetPokemon (pokeName string) (FullPokemon, error) {
	// declaring the return variable
	var fulPok FullPokemon
	// making sure the argument exists
	if pokeName == "" {
		return FullPokemon{}, errors.New("Pokemon Name cannot be empty")
	}
	// declaring the url variable
	url := baseURL + "/pokemon/" + pokeName
	// checking wether Pokemon is already in Cache
	val, ok := c.cache.Get(url)
	if ok {
		//cache hit - unmarshal existing []byte
		if err := json.Unmarshal(val, &fulPok); err != nil {
			return FullPokemon{}, err 
		}
		// return unmarshaled val
		return fulPok, nil 
	}
	// cache miss, create request and  make the client call
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return FullPokemon{}, nil 
	}
	res, err := c.httpClient.Do(req)
	if err != nil {
		return FullPokemon{}, nil
	}
	
	defer res.Body.Close()

	// read the data and store it in Cache
	data, err := io.ReadAll(res.Body)
	if err != nil {
		return FullPokemon{}, err
	}

 c.cache.Add(url, data)

	// unmarshal the data and return import
	if err := json.Unmarshal(data, &fulPok); err != nil {
		return FullPokemon{}, err
	}
	return fulPok, nil
}
