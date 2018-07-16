package atlantic

import "encoding/json"

// ListPrivateNetworksResult represents the result from listing private networks.
type ListPrivateNetworksResult struct {
	Timestamp int                         `json:"Timestamp"`
	Response  ListPrivateNetworksResponse `json:"list-private-networksresponse"`
}

// ListPrivateNetworksResponse represents the response from listing private networks.
type ListPrivateNetworksResponse struct {
	PrivateNetworks map[string]PrivateNetwork `json:"KeysSet"`
	RequestID       string                    `json:"requestid"`
}

// PrivateNetwork represents a private network.
type PrivateNetwork struct {
	IPRange string `json:"ip_range"`
	Network string `json:"network"`
	Prefix  string `json:"prefix"`
}

// ListPrivateNetworksOutput represents the output from listing private networks.
type ListPrivateNetworksOutput struct {
	PrivateNetworks []PrivateNetwork
}

// ListPrivateNetworks returns all private network ranges assigned to the account.
func (client *Client) ListPrivateNetworks() (*ListPrivateNetworksOutput, error) {
	action := "list-private-networks"

	response, err := client.request(action)
	if err != nil {
		return nil, err
	}

	var res ListPrivateNetworksResult
	if err := json.Unmarshal([]byte(response), &res); err != nil {
		return nil, err
	}

	pns := []PrivateNetwork{}
	for _, pn := range res.Response.PrivateNetworks {
		pns = append(pns, pn)
	}

	output := &ListPrivateNetworksOutput{
		PrivateNetworks: pns,
	}

	return output, nil
}
