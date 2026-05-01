package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) ListLocations (pageURL *string ) (Locations, error) {
	// declaring the return variable
	var loc Locations 
	// Setting the correct url to recieve locations 
	url := baseURL + "/location-area"
	if  pageURL 	!= nil {
		url = *pageURL
	}
	// Checking wether this url is in cache already
	val, ok := c.cache.Get(url)
	if ok {
		// cache hit - decode the existing val []byte 
		if err := json.Unmarshal(val, &loc); err != nil {
			return Locations{},err
		} 
	return loc, nil
	}
	// cache miss - make the Client call and store the res in cache
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return Locations{}, err
	}
	res, err := c.httpClient.Do(req)
	if err != nil {
		return Locations{}, err
	}
	defer res.Body.Close()
	
	data, err := io.ReadAll(res.Body)
	if err != nil {
		return Locations{}, err
	}
	// store the res.Body in cache
	c.cache.Add(url, data)
	// unmarshall the data and return the Locations
	// err exists already, so no : here 
	err = json.Unmarshal(data, &loc)
	if err != nil {
		return Locations{}, err 
	}
	return loc, nil
}
