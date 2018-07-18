package atlantic

import (
	"encoding/json"
	"fmt"
	"strings"
)

// ListInstancesResult represents the result from listing instances.
type ListInstancesResult struct {
	Timestamp int `json:"Timestamp"`
	Response  struct {
		ListInstances map[string]ListInstance `json:"instancesSet"`
		RequestID     string                  `json:"requestid"`
	} `json:"list-instancesresponse"`
}

// ListInstance represents a listed instance.
type ListInstance struct {
	ID               string `json:"InstanceId"`
	CUID             string `json:"cu_id"`
	RatePerHour      string `json:"rate_per_hr"`
	CPUCount         string `json:"vm_cpu_req"`
	CreatedDate      string `json:"vm_created_date"`
	Description      string `json:"vm_description"`
	DiskSize         string `json:"vm_disk_req"`
	Image            string `json:"vm_image"`
	ImageDisplayName string `json:"vm_image_display_name"`
	IPAddress        string `json:"vm_ip_address"`
	Name             string `json:"vm_name"`
	NetworkCount     string `json:"vm_network_req"`
	OSArchitecture   string `json:"vm_os_architecture"`
	PlanName         string `json:"vm_plan_name"`
	RAMSize          string `json:"vm_ram_req"`
	Status           string `json:"vm_status"`
}

// ListInstancesOutput represents the output from listing instances.
type ListInstancesOutput struct {
	ListInstances []ListInstance
}

// TerminateInstanceResult represents the result from terminating instances.
type TerminateInstanceResult struct {
	Timestamp int `json:"Timestamp"`
	Response  struct {
		TerminateInstances map[string]TerminateInstance `json:"instancesSet"`
		RequestID          string                       `json:"requestid"`
	} `json:"terminate-instanceresponse"`
}

// TerminateInstance represents a terminated instance.
type TerminateInstance struct {
	ID      string `json:"InstanceId"`
	Message string `json:"message"`
	Result  string `json:"result"`
}

// TerminateInstanceInput represents the input for terminating instances.
type TerminateInstanceInput struct {
	InstanceID []string
}

// TerminateInstanceOutput represents the output from terminating instances.
type TerminateInstanceOutput struct {
	TerminateInstances []TerminateInstance
}

// RunInstanceResult represents the result from running instances.
type RunInstanceResult struct {
	Timestamp int `json:"Timestamp"`
	Response  struct {
		RunInstances map[string]RunInstance `json:"instancesSet"`
		RequestID    string                 `json:"requestid"`
	} `json:"run-instanceresponse"`
}

// RunInstance represents a ran instance.
type RunInstance struct {
	ID        string `json:"instanceid"`
	IPAddress string `json:"ip_address"`
	Password  string `json:"password"`
	Username  string `json:"username"`
}

// RunInstanceInput represents the input for running instances.
type RunInstanceInput struct {
	ServerName   string
	ImageID      string
	PlanName     string
	Location     string
	EnableBackup bool
	CloneImage   string
	Qty          int
	Term         string
	KeyID        string
}

// RunInstanceOutput represents the output from running instances.
type RunInstanceOutput struct {
	RunInstances []RunInstance
}

// DescribeInstanceResult represents the result from describing an instance.
type DescribeInstanceResult struct {
	Timestamp int `json:"Timestamp"`
	Response  struct {
		DescribeInstances map[string]DescribeInstance `json:"instanceSet"`
		RequestID         string                      `json:"requestid"`
	} `json:"describe-instanceresponse"`
}

// DescribeInstance represents a described instance.
type DescribeInstance struct {
	ID                          string `json:"InstanceId"`
	ClonedFrom                  string `json:"cloned_from"`
	CUID                        string `json:"cu_id"`
	DisallowDeletion            string `json:"disallow_deletion"`
	RatePerHour                 string `json:"rate_per_hr"`
	Removed                     string `json:"removed"`
	ReprovisioningProcessedDate string `json:"reprovisioning_processed_date"`
	ResetpwdProcessedDate       string `json:"resetpwd_processed_date"`
	VMCPUReq                    string `json:"vm_cpu_req"`
	VMCreatedDate               string `json:"vm_created_date"`
	VMDescription               string `json:"vm_description"`
	VMDiskReq                   string `json:"vm_disk_req"`
	VMID                        string `json:"vm_id"`
	VMImage                     string `json:"vm_image"`
	VMImageDisplayName          string `json:"vm_image_display_name"`
	VMIPAddress                 string `json:"vm_ip_address"`
	VMIPGateway                 string `json:"vm_ip_gateway"`
	VMIPSubnet                  string `json:"vm_ip_subnet"`
	VMNetworkReq                string `json:"vm_network_req"`
	VMOSArchitecture            string `json:"vm_os_architecture"`
	VMPlanName                  string `json:"vm_plan_name"`
	VMRAMReq                    string `json:"vm_ram_req"`
	VMRemovedDate               string `json:"vm_removed_date"`
	VMStatus                    string `json:"vm_status"`
	VMUsername                  string `json:"vm_username"`
}

// DescribeInstanceInput represents the input for describing an instance.
type DescribeInstanceInput struct {
	InstanceID string
}

// DescribeInstanceOutput represents the output from describing an instance.
type DescribeInstanceOutput struct {
	DescribeInstance DescribeInstance
}

// RebootInstanceResult represents the result from rebooting an instance.
type RebootInstanceResult struct {
	Timestamp int `json:"Timestamp"`
	Response  struct {
		RebootInstance RebootInstance `json:"return"`
		RequestID      string         `json:"requestid"`
	} `json:"reboot-instanceresponse"`
}

// RebootInstance represents a rebooted instance.
type RebootInstance struct {
	Message string `json:"message"`
	Value   string `json:"value"`
}

// RebootInstanceInput represents the input for rebooting an instance.
type RebootInstanceInput struct {
	InstanceID string
	RebootType string
}

// RebootInstanceOutput represents the output from rebooting an instance.
type RebootInstanceOutput struct {
	RebootInstance RebootInstance
}

// ShutdownInstanceResult represents the result from shutting down instances.
type ShutdownInstanceResult struct {
	Timestamp int `json:"Timestamp"`
	Response  struct {
		ShutdownInstances map[string]ShutdownInstance `json:"instancesSet"`
		RequestID         string                      `json:"requestid"`
	} `json:"shutdown-instanceresponse"`
}

// ShutdownInstance represents a shut down instance.
type ShutdownInstance struct {
	ID      string `json:"InstanceID"`
	Message string `json:"Message"`
	Value   string `json:"value"`
}

// ShutdownInstanceInput represents the input for shutting down instances.
type ShutdownInstanceInput struct {
	InstanceID   []string
	ShutdownType string
}

// ShutdownInstanceOutput represents the output from shutting down instances.
type ShutdownInstanceOutput struct {
	ShutdownInstances []ShutdownInstance
}

// PowerOnInstanceResult represents the result from powering on instances.
type PowerOnInstanceResult struct {
	Timestamp int `json:"Timestamp"`
	Response  struct {
		PowerOnInstances map[string]PowerOnInstance `json:"instancesSet"`
		RequestID        string                     `json:"requestid"`
	} `json:"power-on-instanceresponse"`
}

// PowerOnInstance represents a powered on instance.
type PowerOnInstance struct {
	ID      string `json:"InstanceID"`
	Message string `json:"Message"`
	Value   string `json:"value"`
}

// PowerOnInstanceInput represents the input for powering on instances.
type PowerOnInstanceInput struct {
	InstanceID []string
}

// PowerOnInstanceOutput represents the output from powering on instances.
type PowerOnInstanceOutput struct {
	PowerOnInstances []PowerOnInstance
}

// ResizeInstanceResult represents the result from resizing an instance.
type ResizeInstanceResult struct {
	Timestamp int `json:"Timestamp"`
	Response  struct {
		ResizeInstances map[string]ResizeInstance `json:"return"`
		RequestID       string                    `json:"requestid"`
	} `json:"resize-instanceresponse"`
}

// ResizeInstance represents a resized instance.
type ResizeInstance struct {
	ID      string `json:"instanceid"`
	Message string `json:"Message"`
	Value   string `json:"value"`
	Status  string `json:"vm_status"`
}

// ResizeInstanceInput represents the input for resizing an instance.
type ResizeInstanceInput struct {
	InstanceID string
	PlanName   string
}

// ResizeInstanceOutput represents the output from resizing an instance.
type ResizeInstanceOutput struct {
	ResizeInstance ResizeInstance
}

// ReprovisionInstanceResult represents the result from reprovisioning an instance.
type ReprovisionInstanceResult struct {
	Timestamp int `json:"Timestamp"`
	Response  struct {
		ReprovisionInstances map[string]ReprovisionInstance `json:"return"`
		RequestID            string                         `json:"requestid"`
	} `json:"reprovision-instanceresponse"`
}

// ReprovisionInstance represents a reprovisioned instance.
type ReprovisionInstance struct {
	Info struct {
		ID      string `json:"instanceid"`
		Message string `json:"Message"`
		Value   string `json:"value"`
		Status  string `json:"vm_status"`
	} `json:"1instance"`
	Item struct {
		Username string `json:"username"`
		Password string `json:"password"`
	} `json:"1item"`
}

// ReprovisionInstanceInput represents the input for reprovisioning an instance.
type ReprovisionInstanceInput struct {
	InstanceID string
	PlanName   string
	ImageID    string
}

// ReprovisionInstanceOutput represents the output from reprovisioning an instance.
type ReprovisionInstanceOutput struct {
	ReprovisionInstance ReprovisionInstance
}

// RunInstance creates one or more new instances.
func (client *Client) RunInstance(input *RunInstanceInput) (*RunInstanceOutput, error) {
	var actionBuilder strings.Builder

	fmt.Fprintf(&actionBuilder, "run-instance&servername=%s&imageid=%s&planname=%s&vm_location=%s", input.ServerName, input.ImageID, input.PlanName, input.Location)

	if input.EnableBackup {
		fmt.Fprintf(&actionBuilder, "&enablebackup=Y")
	} else {
		fmt.Fprintf(&actionBuilder, "&enablebackup=N")
	}

	if input.Qty < 1 {
		input.Qty = 1
	}
	fmt.Fprintf(&actionBuilder, "&serverqty=%d", input.Qty)

	if input.CloneImage != "" {
		fmt.Fprintf(&actionBuilder, "&cloneimage=%s", input.CloneImage)
	}

	if input.Term != "" {
		fmt.Fprintf(&actionBuilder, "&term=%s", input.Term)
	}

	if input.KeyID != "" {
		fmt.Fprintf(&actionBuilder, "&key_id=%s", input.KeyID)
	}

	action := actionBuilder.String()

	response, err := client.request(action)
	if err != nil {
		return nil, err
	}

	var res RunInstanceResult
	err = json.Unmarshal([]byte(response), &res)
	if err != nil {
		return nil, err
	}

	ii := []RunInstance{}
	for _, i := range res.Response.RunInstances {
		ii = append(ii, i)
	}

	output := &RunInstanceOutput{
		RunInstances: ii,
	}

	return output, nil
}

// ListInstances retrieves all active instances.
func (client *Client) ListInstances() (*ListInstancesOutput, error) {
	action := "list-instances"

	response, err := client.request(action)
	if err != nil {
		return nil, err
	}

	var res ListInstancesResult
	err = json.Unmarshal([]byte(response), &res)
	if err != nil {
		return nil, err
	}

	ii := []ListInstance{}
	for _, i := range res.Response.ListInstances {
		ii = append(ii, i)
	}

	output := &ListInstancesOutput{
		ListInstances: ii,
	}

	return output, nil
}

// DescribeInstance retrieves the details of a specific instance.
func (client *Client) DescribeInstance(input *DescribeInstanceInput) (*DescribeInstanceOutput, error) {
	if input.InstanceID == "" {
		return nil, fmt.Errorf("atlantic: Instance ID must be provided")
	}

	action := fmt.Sprintf("describe-instance&instanceid=%s", input.InstanceID)

	response, err := client.request(action)
	if err != nil {
		return nil, err
	}

	var res DescribeInstanceResult
	err = json.Unmarshal([]byte(response), &res)
	if err != nil {
		return nil, err
	}

	i := res.Response.DescribeInstances["item"]

	output := &DescribeInstanceOutput{
		DescribeInstance: i,
	}

	return output, nil
}

// RebootInstance reboots a specific instance.
func (client *Client) RebootInstance(input *RebootInstanceInput) (*RebootInstanceOutput, error) {
	var actionBuilder strings.Builder

	if input.InstanceID == "" {
		return nil, fmt.Errorf("atlantic: Instance ID must be provided")
	}

	fmt.Fprintf(&actionBuilder, "reboot-instance&instanceid=%s", input.InstanceID)

	if input.RebootType != "" {
		fmt.Fprintf(&actionBuilder, "&reboottype=%s", input.RebootType)
	}

	action := actionBuilder.String()

	response, err := client.request(action)
	if err != nil {
		return nil, err
	}

	var res RebootInstanceResult
	err = json.Unmarshal([]byte(response), &res)
	if err != nil {
		return nil, err
	}

	i := res.Response.RebootInstance

	output := &RebootInstanceOutput{
		RebootInstance: i,
	}

	return output, nil
}

// ShutdownInstance shuts down one or more instances.
func (client *Client) ShutdownInstance(input *ShutdownInstanceInput) (*ShutdownInstanceOutput, error) {
	var actionBuilder strings.Builder
	var instancesBuilder strings.Builder
	var instances string

	fmt.Fprintf(&actionBuilder, "shutdown-instance")

	if len(input.InstanceID) == 0 {
		return nil, fmt.Errorf("atlantic: Instance ID must be provided")
	} else if len(input.InstanceID) == 1 {
		fmt.Fprintf(&instancesBuilder, "instanceid=%s", input.InstanceID[0])
		instances = instancesBuilder.String()
	} else {
		for i := 0; i < len(input.InstanceID); i++ {
			fmt.Fprintf(&instancesBuilder, "instanceid_%d=%s&", i+1, input.InstanceID[i])
		}
		// remove the trailing '&'
		instances = strings.TrimRight(instancesBuilder.String(), "&")
	}

	fmt.Fprintf(&actionBuilder, "%s", instances)

	if input.ShutdownType != "" {
		fmt.Fprintf(&actionBuilder, "&shutdowntype=%s", input.ShutdownType)
	}

	action := actionBuilder.String()

	response, err := client.request(action)
	if err != nil {
		return nil, err
	}

	var res ShutdownInstanceResult
	err = json.Unmarshal([]byte(response), &res)
	if err != nil {
		return nil, err
	}

	ii := []ShutdownInstance{}
	for _, i := range res.Response.ShutdownInstances {
		ii = append(ii, i)
	}

	output := &ShutdownInstanceOutput{
		ShutdownInstances: ii,
	}

	return output, nil
}

// PowerOnInstance power's on one or more instances.
func (client *Client) PowerOnInstance(input *PowerOnInstanceInput) (*PowerOnInstanceOutput, error) {
	var instancesBuilder strings.Builder
	var instances string

	if len(input.InstanceID) == 0 {
		return nil, fmt.Errorf("atlantic: Instance ID must be provided")
	} else if len(input.InstanceID) == 1 {
		fmt.Fprintf(&instancesBuilder, "instanceid=%s", input.InstanceID[0])
		instances = instancesBuilder.String()
	} else {
		for i := 0; i < len(input.InstanceID); i++ {
			fmt.Fprintf(&instancesBuilder, "instanceid_%d=%s", i+1, input.InstanceID[i])
		}
		// remove the trailing '&'
		instances = strings.TrimRight(instancesBuilder.String(), "&")
	}

	action := fmt.Sprintf("power-on-instance&%s", instances)

	response, err := client.request(action)
	if err != nil {
		return nil, err
	}

	var res PowerOnInstanceResult
	err = json.Unmarshal([]byte(response), &res)
	if err != nil {
		return nil, err
	}

	ii := []PowerOnInstance{}
	for _, i := range res.Response.PowerOnInstances {
		ii = append(ii, i)
	}

	output := &PowerOnInstanceOutput{
		PowerOnInstances: ii,
	}

	return output, nil
}

// ResizeInstance resizes an instance to a larger plan.
func (client *Client) ResizeInstance(input *ResizeInstanceInput) (*ResizeInstanceOutput, error) {
	if input.InstanceID == "" {
		return nil, fmt.Errorf("atlantic: Instance ID must be provided")
	}

	if input.PlanName == "" {
		return nil, fmt.Errorf("atlantic: Plan name must be provided")
	}

	action := fmt.Sprintf("resize-instance&instanceid=%s&planname=%s", input.InstanceID, input.PlanName)

	response, err := client.request(action)
	if err != nil {
		return nil, err
	}

	var res ResizeInstanceResult
	err = json.Unmarshal([]byte(response), &res)
	if err != nil {
		return nil, err
	}

	i := res.Response.ResizeInstances["1instance"]

	output := &ResizeInstanceOutput{
		ResizeInstance: i,
	}

	return output, nil
}

// ReprovisionInstance reprovisions (rebuilds) an instance with the same or different specifications.
func (client *Client) ReprovisionInstance(input *ReprovisionInstanceInput) (*ReprovisionInstanceOutput, error) {
	if input.InstanceID == "" {
		return nil, fmt.Errorf("atlantic: Instance ID must be provided")
	}

	if input.PlanName == "" {
		return nil, fmt.Errorf("atlantic: Plan name must be provided")
	}

	if input.ImageID == "" {
		return nil, fmt.Errorf("atlantic: Image ID must be provided")
	}

	action := fmt.Sprintf("reprovision-instance&instanceid=%s&planname=%s&imageid=%s", input.InstanceID, input.PlanName, input.ImageID)

	response, err := client.request(action)
	if err != nil {
		return nil, err
	}

	var res ReprovisionInstanceResult
	err = json.Unmarshal([]byte(response), &res)
	if err != nil {
		return nil, err
	}

	i := res.Response.ReprovisionInstances["return"]

	output := &ReprovisionInstanceOutput{
		ReprovisionInstance: i,
	}

	return output, nil
}

// TerminateInstance removes one or more instances.
func (client *Client) TerminateInstance(input *TerminateInstanceInput) (*TerminateInstanceOutput, error) {
	var instancesBuilder strings.Builder
	var instances string

	if len(input.InstanceID) == 0 {
		return nil, fmt.Errorf("atlantic: Instance ID must be provided")
	} else if len(input.InstanceID) == 1 {
		fmt.Fprintf(&instancesBuilder, "instanceid=%s", input.InstanceID[0])
		instances = instancesBuilder.String()
	} else {
		for i := 0; i < len(input.InstanceID); i++ {
			fmt.Fprintf(&instancesBuilder, "instanceid_%d=%s&", i+1, input.InstanceID[i])
		}
		// remove the trailing '&'
		instances = strings.TrimRight(instancesBuilder.String(), "&")
	}

	action := fmt.Sprintf("terminate-instance&%s", instances)

	response, err := client.request(action)
	if err != nil {
		return nil, err
	}

	var res TerminateInstanceResult
	err = json.Unmarshal([]byte(response), &res)
	if err != nil {
		return nil, err
	}

	ii := []TerminateInstance{}
	for _, i := range res.Response.TerminateInstances {
		ii = append(ii, i)
	}

	output := &TerminateInstanceOutput{
		TerminateInstances: ii,
	}

	return output, nil
}
