package pokeapi

import(
	"encoding/json"
	"io"
	"net/http"
	"errors"
)

func (c *Client) GetLocationArea (areaName string) (LocationArea, error) {
// declaring the return Variable
	var locArea LocationArea 
	// making sure the argument exists
	if areaName == "" {
		return LocationArea{}, errors.New("areaName cannot be empty")
	}
	// declaring the url Variable
	url := baseURL + "/location-area/" + areaName 
	// checking wether this url is already in cache
	val, ok := c.cache.Get(url)
	if ok {
		// cache hit - decode the existing val []byte
		if err := json.Unmarshal(val, &locArea); err != nil {
			return LocationArea{}, err 
		}
		return locArea, nil 
	}
	// cache miss - make the client call and store the res in cache
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return LocationArea{}, err 
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return LocationArea{}, err 
	}

	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return LocationArea{}, err 	
	}

	// store the data in cache
	c.cache.Add(url, data)

	//Unmarshal the data and return the locArea struct
	// err exists here, so no :
	err = json.Unmarshal(data, &locArea)
	if err != nil {
		return LocationArea{}, err 
	}
	return locArea, nil
}
