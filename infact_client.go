package infakt

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

const InfaktEndpoint string = "https://api.infakt.pl/v3"
const AuthHeader string = "X-inFakt-ApiKey"

// NewInFaktClient -
func NewInFaktClient(host *string, token *string) (*InFaktClient, error) {
	c := InFaktClient{
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

func (c *InFaktClient) doRequest(req *http.Request) ([]byte, error) {
	token := c.Token
	req.Header.Set(c.AuthHeader, token)
	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	switch res.StatusCode {
	case http.StatusCreated, http.StatusOK, http.StatusNoContent:
		return body, err

	default:
		return nil, fmt.Errorf("status: %d, body: %s", res.StatusCode, body)
	}

}

func DoRequest(c *InFaktClient, req *http.Request, debug bool) ([]byte, error) {
	if debug {
		body, err := c.doRequest(req)
		return body, err
	}
	return []byte("Debug mode off"), nil
}
