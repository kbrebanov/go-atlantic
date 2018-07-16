package atlantic

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	uuid "github.com/satori/go.uuid"
)

// Client represents an Atlantic API client.
type Client struct {
	Version    string
	EndPoint   string
	Format     string
	AccessKey  string
	PrivateKey string
}

// ClientError represents an Atlantic API client error.
type ClientError struct {
	Error ErrAtlantic `json:"error"`
}

// ErrAtlantic represents an Atlantic API error.
type ErrAtlantic struct {
	Code      string `json:"code"`
	Message   string `json:"message"`
	Timestamp int    `json:"time"`
}

func (e ErrAtlantic) Error() string {
	return fmt.Sprintf("atlantic: %s (%s)", e.Message, e.Code)
}

// NewClient returns a new Atlantic API client.
func NewClient(accesskey string, privatekey string) *Client {
	return &Client{
		Version:    "2010-12-30",
		EndPoint:   "https://cloudapi.atlantic.net/",
		Format:     "json",
		AccessKey:  accesskey,
		PrivateKey: privatekey,
	}
}

// generateSignature returns a signature required when sending a client request.
func (client *Client) generateSignature(timeSinceEpoch int64, randomUUID string) string {
	key := []byte(client.PrivateKey)
	stringToSign := fmt.Sprintf("%d%s", timeSinceEpoch, randomUUID)

	m := hmac.New(sha256.New, key)
	m.Write([]byte(stringToSign))

	return base64.StdEncoding.EncodeToString(m.Sum(nil))
}

// request sends a request to Atlantic's API.
func (client *Client) request(action string) (string, error) {
	randomUUID := uuid.NewV4().String()
	timeSinceEpoch := time.Now().Unix()
	signature := client.generateSignature(timeSinceEpoch, randomUUID)
	form := url.Values{}
	form.Add("Format", client.Format)
	form.Add("Version", client.Version)
	form.Add("ACSAccessKeyId", client.AccessKey)
	form.Add("Timestamp", strconv.FormatInt(timeSinceEpoch, 10))
	form.Add("Rndguid", randomUUID)
	form.Add("Signature", signature)
	form.Add("Action", action)

	encodedForm := form.Encode()

	request, err := http.NewRequest("POST", client.EndPoint, strings.NewReader(encodedForm))
	if err != nil {
		return "", err
	}

	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()

	jsonData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", err
	}

	if bytes.HasPrefix(jsonData, []byte("{\"error\":")) {
		var ce ClientError
		if err := json.Unmarshal([]byte(jsonData), &ce); err != nil {
			return "", err
		}
		return "", ce.Error
	}

	var prettyJSONData bytes.Buffer
	if err := json.Indent(&prettyJSONData, jsonData, "", "  "); err != nil {
		return "", err
	}

	return string(prettyJSONData.Bytes()), nil
}
