package atlantic

import (
	"encoding/json"
	"fmt"
	"strings"
)

// DescribeImageResult represents the result from describing an image.
type DescribeImageResult struct {
	Timestamp int                   `json:"Timestamp"`
	Response  DescribeImageResponse `json:"describe-imageresponse"`
}

// DescribeImageResponse represents the response from describing an image.
type DescribeImageResponse struct {
	Images    map[string]Image `json:"imagesset"`
	RequestID string           `json:"requestid"`
}

// Image struct represents an image.
type Image struct {
	Architecture string `json:"architecture"`
	DisplayName  string `json:"displayname"`
	Type         string `json:"image_type"`
	ID           string `json:"imageid"`
	OSType       string `json:"ostype"`
	Owner        string `json:"owner"`
	Platform     string `json:"platform"`
	Version      string `json:"version"`
}

// DescribeImageInput represents the input for describing an image.
type DescribeImageInput struct {
	ImageID string
}

// DescribeImageOutput represents the output from describing an image.
type DescribeImageOutput struct {
	Images []Image
}

// DescribeImage returns the description of a specific, or all, cloud images
func (client *Client) DescribeImage(input *DescribeImageInput) (*DescribeImageOutput, error) {
	var actionBuilder strings.Builder

	fmt.Fprintf(&actionBuilder, "describe-image")

	if input.ImageID != "" {
		fmt.Fprintf(&actionBuilder, "&imageid=%s", input.ImageID)
	}

	action := actionBuilder.String()

	response, err := client.request(action)
	if err != nil {
		return nil, err
	}

	var res DescribeImageResult
	if err := json.Unmarshal([]byte(response), &res); err != nil {
		return nil, err
	}

	ii := []Image{}
	for _, i := range res.Response.Images {
		ii = append(ii, i)
	}

	output := &DescribeImageOutput{
		Images: ii,
	}

	return output, nil
}
