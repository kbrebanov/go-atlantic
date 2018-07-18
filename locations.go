package atlantic

import "encoding/json"

// ListLocationsResult represents the result from listing locations.
type ListLocationsResult struct {
	Timestamp int `json:"Timestamp"`
	Response  struct {
		Locations map[string]Location `json:"KeysSet"`
		RequestID string              `json:"requestid"`
	} `json:"list-locationsresponse"`
}

// Location struct represents a location.
type Location struct {
	Description string `json:"description"`
	InfoMessage string `json:"info_message"`
	Active      string `json:"is_active"`
	Code        string `json:"location_code"`
	Name        string `json:"location_name"`
}

// ListLocationsOutput represents the output from listing locations.
type ListLocationsOutput struct {
	Locations []Location
}

// ListLocations returns all available locations.
func (client *Client) ListLocations() (*ListLocationsOutput, error) {
	action := "list-locations"

	response, err := client.request(action)
	if err != nil {
		return nil, err
	}

	var res ListLocationsResult
	if err := json.Unmarshal([]byte(response), &res); err != nil {
		return nil, err
	}

	ll := []Location{}
	for _, l := range res.Response.Locations {
		ll = append(ll, l)
	}

	output := &ListLocationsOutput{
		Locations: ll,
	}

	return output, nil
}
