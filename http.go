package opensea

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func (c *Client) apiRequest(URI string) []byte {
	URL := fmt.Sprintf("%s/%s", c.URL, URI)
	req, err := http.NewRequest("GET", URL, nil)
	if err != nil {
		log.Fatalf("[ERROR] %s", err)
	}
	req.Header.Add("X-API-KEY", c.APIKey)

	if c.LogLevel == "verbose" {
		log.Printf("[VERBOSE] GET %s", URL)
	}

	client := &http.Client{}
	resp, _ := client.Do(req)

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("[ERROR] %s", err)
	}
	return body
}
