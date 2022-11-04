package infact

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

const InfaktEndpoint string = "https://api.infakt.pl/v3"
const AuthHeader string = "X-inFakt-ApiKey"

// Client -
type Client struct {
	InfaktEndpoint string
	HTTPClient     *http.Client
	Token          string
	AuthHeader     string
}

// NewClient -
func NewClient(host, token *string) (*Client, error) {
	c := Client{
		HTTPClient:     &http.Client{Timeout: 10 * time.Second},
		InfaktEndpoint: InfaktEndpoint,
		AuthHeader:     AuthHeader,
	}

	if host != nil {
		c.InfaktEndpoint = *host
	}

	if token != nil {
		c.Token = *token
	}

	return &c, nil
}

func (c *Client) DoRequest(req *http.Request, authToken *string) ([]byte, error) {
	token := c.Token

	if authToken != nil {
		token = *authToken
	}

	req.Header.Set(c.AuthHeader, token)

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("status: %d, body: %s", res.StatusCode, body)
	}

	return body, err
}
