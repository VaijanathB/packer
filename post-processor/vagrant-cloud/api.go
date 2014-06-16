package vagrantcloud

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"
)

type Box struct {
	Tag string `json:"tag"`
}

type VagrantCloudClient struct {
	// The http client for communicating
	client *http.Client

	// The base URL of the API
	BaseURL string

	// Access token
	AccessToken string
}

func (v VagrantCloudClient) New(baseUrl string, token string) *VagrantCloudClient {
	c := &VagrantCloudClient{
		client: &http.Client{
			Transport: &http.Transport{
				Proxy: http.ProxyFromEnvironment,
			},
		},
		BaseURL:     baseUrl,
		AccessToken: token,
	}
	return c
}

func decodeBody(resp *http.Response, out interface{}) error {
	defer resp.Body.Close()
	dec := json.NewDecoder(resp.Body)
	return dec.Decode(out)
}

// encodeBody is used to encode a request body
func encodeBody(obj interface{}) (io.Reader, error) {
	buf := bytes.NewBuffer(nil)
	enc := json.NewEncoder(buf)
	if err := enc.Encode(obj); err != nil {
		return nil, err
	}
	return buf, nil
}

func (v VagrantCloudClient) Box(tag string) (*Box, error) {
	resp, err := v.Get(tag)

	if err != nil {
		return nil, fmt.Errorf("Error retrieving box: %s", err)
	}

	box := &Box{}

	if err = decodeBody(resp, box); err != nil {
		return nil, fmt.Errorf("Error parsing box response: %s", err)
	}

	return box, nil
}

func (v VagrantCloudClient) Get(path string) (*http.Response, error) {
	params := url.Values{}
	params.Set("access_token", v.AccessToken)
	reqUrl := fmt.Sprintf("%s/%s%s", v.BaseURL, path, params.Encode())

	// Scrub API key for logs
	scrubbedUrl := strings.Replace(reqUrl, v.AccessToken, "ACCESS_TOKEN", -1)
	log.Printf("Post-Processor Vagrant Cloud API GET: %s", scrubbedUrl)

	req, err := http.NewRequest("GET", reqUrl, nil)
	resp, err := v.client.Do(req)

	log.Printf("Post-Processor Vagrant Cloud API Response: \n\n%s", resp)

	return resp, err
}

func (v VagrantCloudClient) Post(path string, body map[string]interface{}) (map[string]interface{}, error) {

	// Scrub API key for logs
	scrubbedUrl := strings.Replace(path, v.AccessToken, "ACCESS_TOKEN", -1)
	log.Printf("Post-Processor Vagrant Cloud API POST: %s. \n\n Body: %s", scrubbedUrl, body)
	return nil, nil
}
