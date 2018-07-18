package atlantic

import (
	"encoding/json"
	"fmt"
	"strings"
)

// ListSSHKeysResult represents the result from listing SSH keys.
type ListSSHKeysResult struct {
	Timestamp int `json:"Timestamp"`
	Response  struct {
		SSHKeys   map[string]SSHKey `json:"KeysSet"`
		RequestID string            `json:"requestid"`
	} `json:"list-sshkeysresponse"`
}

// SSHKey represents an SSH key.
type SSHKey struct {
	ID        string `json:"key_id"`
	Name      string `json:"key_name"`
	PublicKey string `json:"public_key"`
}

// ListSSHKeysOutput represents the output from listing SSH keys.
type ListSSHKeysOutput struct {
	Keys []SSHKey
}

// AddSSHKeyResult represents the result from adding an SSH key.
type AddSSHKeyResult struct {
	Timestamp int `json:"Timestamp"`
	Response  struct {
		AddSSHKey AddSSHKey `json:"result"`
		RequestID string    `json:"requestid"`
	} `json:"add-sshkeyresponse"`
}

// AddSSHKey represents an added SSH key.
type AddSSHKey struct {
	ID      string `json:"key_id"`
	Message string `json:"message"`
	Result  string `json:"result"`
}

// AddSSHKeyInput represents the input for adding an SSH key.
type AddSSHKeyInput struct {
	KeyName   string
	PublicKey string
}

// AddSSHKeyOutput represents the output from adding an SSH key.
type AddSSHKeyOutput struct {
	ID      string
	Message string
}

// DeleteSSHKeyResult represents the result from deleting SSH keys.
type DeleteSSHKeyResult struct {
	Timestamp int `json:"Timestamp"`
	Response  struct {
		DeleteSSHKeys map[string]DeleteSSHKey `json:"delete-sshkey"`
		RequestID     string                  `json:"requestid"`
	} `json:"delete-sshkeyresponse"`
}

// DeleteSSHKey represents a deleted SSH key.
type DeleteSSHKey struct {
	ID      string `json:"key_id"`
	Message string `json:"message"`
	Result  string `json:"result"`
}

// DeleteSSHKeyInput represents the input for deleting SSH keys.
type DeleteSSHKeyInput struct {
	KeyIDs []string
}

// DeleteSSHKeyOutput represents the output from deleting SSH keys.
type DeleteSSHKeyOutput struct {
	Keys []DeleteSSHKey
}

// ListSSHKeys returns the details of all SSH keys that have been added to the account.
func (client *Client) ListSSHKeys() (*ListSSHKeysOutput, error) {
	action := "list-sshkeys"

	response, err := client.request(action)
	if err != nil {
		return nil, err
	}

	var res ListSSHKeysResult
	if err := json.Unmarshal([]byte(response), &res); err != nil {
		return nil, err
	}

	kk := []SSHKey{}
	for _, k := range res.Response.SSHKeys {
		kk = append(kk, k)
	}

	output := &ListSSHKeysOutput{
		Keys: kk,
	}

	return output, nil
}

// GetSSHKeyID returns the key ID associated to the given key name.
func (client *Client) GetSSHKeyID(keyName string) (string, error) {
	sshKeys, err := client.ListSSHKeys()

	if err != nil {
		return "", err
	}

	for _, k := range sshKeys.Keys {
		if k.Name == keyName {
			return k.ID, nil
		}
	}

	return "", fmt.Errorf("atlantic: ssh key not found")
}

// AddSSHKey adds an SSH key to the account.
func (client *Client) AddSSHKey(input *AddSSHKeyInput) (*AddSSHKeyOutput, error) {
	if input.KeyName == "" {
		return nil, fmt.Errorf("atlantic: Key name must be provided")
	}

	if input.PublicKey == "" {
		return nil, fmt.Errorf("atlantic: Public key must be provided")
	}

	action := fmt.Sprintf("add-sshkey&key_name=%s&public_key=%s", input.KeyName, input.PublicKey)

	response, err := client.request(action)
	if err != nil {
		return nil, err
	}

	var res AddSSHKeyResult
	if err := json.Unmarshal([]byte(response), &res); err != nil {
		return nil, err
	}

	output := &AddSSHKeyOutput{
		ID:      res.Response.AddSSHKey.ID,
		Message: res.Response.AddSSHKey.Message,
	}

	return output, nil
}

// DeleteSSHKey deletes one or more SSH keys from the account.
func (client *Client) DeleteSSHKey(input *DeleteSSHKeyInput) (*DeleteSSHKeyOutput, error) {
	if len(input.KeyIDs) == 0 {
		return nil, fmt.Errorf("atlantic: SSH key ID must be provided")
	}

	var actionBuilder strings.Builder

	fmt.Fprintf(&actionBuilder, "delete-sshkey")

	if len(input.KeyIDs) == 1 {
		fmt.Fprintf(&actionBuilder, "&key_id=%s", input.KeyIDs[0])
	} else {
		fmt.Fprintf(&actionBuilder, "&key_id=%s", input.KeyIDs[0])
		for _, keyID := range input.KeyIDs[1:] {
			fmt.Fprintf(&actionBuilder, ",%s", keyID)
		}
	}

	action := actionBuilder.String()

	response, err := client.request(action)
	if err != nil {
		return nil, err
	}

	var res DeleteSSHKeyResult
	if err := json.Unmarshal([]byte(response), &res); err != nil {
		return nil, err
	}

	kk := []DeleteSSHKey{}
	for _, k := range res.Response.DeleteSSHKeys {
		kk = append(kk, k)
	}

	output := &DeleteSSHKeyOutput{
		Keys: kk,
	}

	return output, nil
}
