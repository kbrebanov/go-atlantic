package atlantic

import (
	"encoding/json"
	"fmt"
	"strings"
)

// ListPublicIPsResult represents the result from listing public IP's.
type ListPublicIPsResult struct {
	Timestamp int `json:"Timestamp"`
	Response  struct {
		PublicIPs map[string]PublicIP `json:"KeysSet"`
		RequestID string              `json:"requestid"`
	} `json:"list-public-ipsresponse"`
}

// ListPublicIPsInput represents the input for listing public IP's.
type ListPublicIPsInput struct {
	Location  string
	IPAddress string
}

// ListPublicIPsOutput represents the output from listing public IP's.
type ListPublicIPsOutput struct {
	PublicIPs []PublicIP
}

// PublicIP represents a public IP.
type PublicIP struct {
	InstanceID string `json:"instanceid"`
	Address    string `json:"ip_address"`
	Gateway    string `json:"ip_gateway"`
	Location   string `json:"ip_location"`
	Subnet     string `json:"ip_subnet"`
}

// ReservePublicIPResult represents the result from reserving a public IP.
type ReservePublicIPResult struct {
	Timestamp int `json:"Timestamp"`
	Response  struct {
		ReservePublicIPs map[string]ReservePublicIP `json:"reserve-ip"`
		RequestID        string                     `json:"requestid"`
	} `json:"reserve-public-ipresponse"`
}

// ReservePublicIP represents a reserved public IP.
type ReservePublicIP struct {
	Address  string `json:"ip_address"`
	DNS1     string `json:"ip_dns1"`
	DNS2     string `json:"ip_dns2"`
	Gateway  string `json:"ip_gateway"`
	Location string `json:"ip_location"`
	Subnet   string `json:"ip_subnet"`
	Message  string `json:"message"`
	Result   string `json:"result"`
}

// ReservePublicIPInput represents the input for reserving a public IP.
type ReservePublicIPInput struct {
	Location string
	Qty      int
}

// ReservePublicIPOutput represents the output from reserving a public IP.
type ReservePublicIPOutput struct {
	ReservePublicIPs []ReservePublicIP
}

// ReleasePublicIPResult represents the result from releasing public IP's.
type ReleasePublicIPResult struct {
	Timestamp int `json:"Timestamp"`
	Response  struct {
		ReleasePublicIPs map[string]ReleasePublicIP `json:"release-ip"`
		RequestID        string                     `json:"requestid"`
	} `json:"release-public-ipresponse"`
}

// ReleasePublicIP represents a released public IP.
type ReleasePublicIP struct {
	Address string `json:"ip_address"`
	Message string `json:"message"`
	Result  string `json:"result"`
}

// ReleasePublicIPInput represents the input for releasing public IP's.
type ReleasePublicIPInput struct {
	IPAddress []string
}

// ReleasePublicIPOutput represents the output from releasing public IP's.
type ReleasePublicIPOutput struct {
	ReleasePublicIPs []ReleasePublicIP
}

// AssignPublicIPResult represents the result from assigning public IP's.
type AssignPublicIPResult struct {
	Timestamp int `json:"Timestamp"`
	Response  struct {
		AssignPublicIPs map[string]AssignPublicIP `json:"assign-ip"`
		RequestID       string                    `json:"requestid"`
	} `json:"assign-public-ipresponse"`
}

// AssignPublicIP represents an assigned public IP.
type AssignPublicIP struct {
	InstanceID int    `json:"instanceid"`
	Address    string `json:"ip_address"`
	Message    string `json:"message"`
	Result     string `json:"result"`
}

// AssignPublicIPInput represents the input for assigning public IP's.
type AssignPublicIPInput struct {
	IPAddress  []string
	InstanceID string
}

// AssignPublicIPOutput represents the output from assigning public IP's.
type AssignPublicIPOutput struct {
	AssignPublicIPs []AssignPublicIP
}

// UnassignPublicIPResult represents the result from unassigning public IP's.
type UnassignPublicIPResult struct {
	Timestamp int `json:"Timestamp"`
	Response  struct {
		UnassignPublicIPs map[string]UnassignPublicIP `json:"unassign-ip"`
		RequestID         string                      `json:"requestid"`
	} `json:"unassign-public-ipresponse"`
}

// UnassignPublicIP represents an unassigned public IP.
type UnassignPublicIP struct {
	Address string `json:"ip_address"`
	Message string `json:"message"`
	Result  string `json:"result"`
}

// UnassignPublicIPInput represents the input for unassigning public IP's.
type UnassignPublicIPInput struct {
	IPAddress []string
}

// UnassignPublicIPOutput represents the output from unassigning public IP's.
type UnassignPublicIPOutput struct {
	UnassignPublicIPs []UnassignPublicIP
}

// ListPublicIPs returns the details of the additional public IP addresses reserved on the account.
func (client *Client) ListPublicIPs(input *ListPublicIPsInput) (*ListPublicIPsOutput, error) {
	var actionBuilder strings.Builder

	fmt.Fprintf(&actionBuilder, "list-public-ips")

	if input.Location != "" {
		fmt.Fprintf(&actionBuilder, "&location=%s", input.Location)
	}

	if input.IPAddress != "" {
		fmt.Fprintf(&actionBuilder, "&ip_address=%s", input.IPAddress)
	}

	action := actionBuilder.String()

	response, err := client.request(action)
	if err != nil {
		return nil, err
	}

	var res ListPublicIPsResult
	if err := json.Unmarshal([]byte(response), &res); err != nil {
		return nil, err
	}

	ips := []PublicIP{}
	for _, ip := range res.Response.PublicIPs {
		ips = append(ips, ip)
	}

	output := &ListPublicIPsOutput{
		PublicIPs: ips,
	}

	return output, nil
}

// ReservePublicIP reserves one or more public IP address in specified location.
func (client *Client) ReservePublicIP(input *ReservePublicIPInput) (*ReservePublicIPOutput, error) {
	if input.Location == "" {
		return nil, fmt.Errorf("atlantic: Location must be provided")
	}

	var actionBuilder strings.Builder

	fmt.Fprintf(&actionBuilder, "reserve-public-ip&location=%s", input.Location)

	if input.Qty < 1 {
		input.Qty = 1
	}
	fmt.Fprintf(&actionBuilder, "&qty=%d", input.Qty)

	action := actionBuilder.String()

	response, err := client.request(action)
	if err != nil {
		return nil, err
	}

	var res ReservePublicIPResult
	if err := json.Unmarshal([]byte(response), &res); err != nil {
		return nil, err
	}

	ips := []ReservePublicIP{}
	for _, ip := range res.Response.ReservePublicIPs {
		ips = append(ips, ip)
	}

	output := &ReservePublicIPOutput{
		ReservePublicIPs: ips,
	}

	return output, nil
}

// ReleasePublicIP releases one or more additional public IP addresses from account.
func (client *Client) ReleasePublicIP(input *ReleasePublicIPInput) (*ReleasePublicIPOutput, error) {
	if len(input.IPAddress) == 0 {
		return nil, fmt.Errorf("atlantic: IP Address must be provided")
	}

	var actionBuilder strings.Builder

	fmt.Fprintf(&actionBuilder, "release-public-ip")

	if len(input.IPAddress) == 1 {
		fmt.Fprintf(&actionBuilder, "&ip_address=%s", input.IPAddress[0])
	} else {
		fmt.Fprintf(&actionBuilder, "&ip_address=%s", input.IPAddress[0])
		for _, ip := range input.IPAddress[1:] {
			fmt.Fprintf(&actionBuilder, ",%s", ip)
		}
	}

	action := actionBuilder.String()

	response, err := client.request(action)
	if err != nil {
		return nil, err
	}

	var res ReleasePublicIPResult
	if err := json.Unmarshal([]byte(response), &res); err != nil {
		return nil, err
	}

	ips := []ReleasePublicIP{}
	for _, ip := range res.Response.ReleasePublicIPs {
		ips = append(ips, ip)
	}

	output := &ReleasePublicIPOutput{
		ReleasePublicIPs: ips,
	}

	return output, nil
}

// AssignPublicIP assigns one or more public IP addresses to a server.
func (client *Client) AssignPublicIP(input *AssignPublicIPInput) (*AssignPublicIPOutput, error) {
	if input.InstanceID == "" {
		return nil, fmt.Errorf("atlantic: Instance ID must be provided")
	}

	if len(input.IPAddress) == 0 {
		return nil, fmt.Errorf("atlantic: IP address must be provided")
	}

	var actionBuilder strings.Builder

	fmt.Fprintf(&actionBuilder, "assign-public-ip&instanceid=%s", input.InstanceID)

	if len(input.IPAddress) == 1 {
		fmt.Fprintf(&actionBuilder, "&ip_address=%s", input.IPAddress[0])
	} else {
		fmt.Fprintf(&actionBuilder, "&ip_address=%s", input.IPAddress[0])
		for _, ip := range input.IPAddress[1:] {
			fmt.Fprintf(&actionBuilder, ",%s", ip)
		}
	}

	action := actionBuilder.String()

	response, err := client.request(action)
	if err != nil {
		return nil, err
	}

	var res AssignPublicIPResult
	if err := json.Unmarshal([]byte(response), &res); err != nil {
		return nil, err
	}

	ips := []AssignPublicIP{}
	for _, ip := range res.Response.AssignPublicIPs {
		ips = append(ips, ip)
	}

	output := &AssignPublicIPOutput{
		AssignPublicIPs: ips,
	}

	return output, nil
}

// UnassignPublicIP unassigns one or more public IP addresses from server.
func (client *Client) UnassignPublicIP(input *UnassignPublicIPInput) (*UnassignPublicIPOutput, error) {
	if len(input.IPAddress) == 0 {
		return nil, fmt.Errorf("atlantic: IP address must be provided")
	}

	var actionBuilder strings.Builder

	fmt.Fprintf(&actionBuilder, "unassign-public-ip")

	if len(input.IPAddress) == 1 {
		fmt.Fprintf(&actionBuilder, "&ip_address=%s", input.IPAddress[0])
	} else {
		fmt.Fprintf(&actionBuilder, "&ip_address=%s", input.IPAddress[0])
		for _, ip := range input.IPAddress[1:] {
			fmt.Fprintf(&actionBuilder, ",%s", ip)
		}
	}

	action := actionBuilder.String()

	response, err := client.request(action)
	if err != nil {
		return nil, err
	}

	var res UnassignPublicIPResult
	if err := json.Unmarshal([]byte(response), &res); err != nil {
		return nil, err
	}

	ips := []UnassignPublicIP{}
	for _, ip := range res.Response.UnassignPublicIPs {
		ips = append(ips, ip)
	}

	output := &UnassignPublicIPOutput{
		UnassignPublicIPs: ips,
	}

	return output, nil
}
