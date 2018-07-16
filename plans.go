package atlantic

import (
	"encoding/json"
	"fmt"
	"strings"
)

// DescribePlanResult represents the result from describing a plan.
type DescribePlanResult struct {
	Timestamp int                  `json:"Timestamp"`
	Response  DescribePlanResponse `json:"describe-planresponse"`
}

// DescribePlanResponse represents the response from describing a plan.
type DescribePlanResponse struct {
	Plans     map[string]Plan `json:"plans"`
	RequestID string          `json:"requestid"`
}

// Plan struct represents a plan.
type Plan struct {
	CentOSCapable    string `json:"centos_capable"`
	CPanelCapable    string `json:"cpanel_capable"`
	DisplayDisk      string `json:"display_disk"`
	DisplayRAM       string `json:"display_ram"`
	FreeTransfer     string `json:"free_transfer"`
	NumCPU           string `json:"num_cpu"`
	OSType           string `json:"ostype"`
	Locked           string `json:"plan_locked"`
	Name             string `json:"plan_name"`
	Type             string `json:"plan_type"`
	Platform         string `json:"platform"`
	RatePerHour      string `json:"rate_per_hr"`
	RatePerHour1Year string `json:"rate_per_hr_1y"`
	RatePerHour3Year string `json:"rate_per_hr_3y"`
	WindowsCapable   string `json:"windows_capable"`
}

// DescribePlanInput represents the input for describing a plan.
type DescribePlanInput struct {
	PlanName string
	Platform string
}

// DescribePlanOutput represents the output from describing a plan.
type DescribePlanOutput struct {
	Plans []Plan
}

// DescribePlan returns the description of all, or a specific, server plans.
func (client *Client) DescribePlan(input *DescribePlanInput) (*DescribePlanOutput, error) {
	var actionBuilder strings.Builder

	fmt.Fprintf(&actionBuilder, "describe-plan")

	if input.PlanName != "" {
		fmt.Fprintf(&actionBuilder, "&planName=%s", input.PlanName)
	}

	if input.Platform != "" {
		fmt.Fprintf(&actionBuilder, "&platform=%s", input.Platform)
	}

	action := actionBuilder.String()

	response, err := client.request(action)
	if err != nil {
		return nil, err
	}

	var res DescribePlanResult
	if err := json.Unmarshal([]byte(response), &res); err != nil {
		return nil, err
	}

	pp := []Plan{}
	for _, p := range res.Response.Plans {
		pp = append(pp, p)
	}

	output := &DescribePlanOutput{
		Plans: pp,
	}

	return output, nil
}
