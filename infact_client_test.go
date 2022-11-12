package infakt_test

import (
	"fmt"
	"net/http"
	"os"
	"testing"

	infakt "github.com/sequring/infakt-client-go"
	"github.com/stretchr/testify/assert"
)

const DebugRequest bool = true

var host string = "https://api.infakt.pl/v3"
var token string

func GetInfactClient() *infakt.InFaktClient {
	var client *infakt.InFaktClient
	token = os.Getenv("INFAKT_TOKEN")
	client, _ = infakt.NewInFaktClient(nil, &token)
	return client
}
func TestSomething(t *testing.T) {

	// assert equality
	assert.Equal(t, 123, 123, "they should be equal")
	// assert inequality
	assert.NotEqual(t, 123, 456, "they should not be equal")

}

func TestInfactClient(t *testing.T) {
	client := GetInfactClient()
	assert.NotNil(t, client, "they should be not null")
	assert.Equal(t, client.InfaktEndpoint, host)
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/clients.json", client.InfaktEndpoint), nil)
	assert.Nil(t, err, "Request should be not null ")
	assert.NotNil(t, req)
	body, err := infakt.DoRequest(client, req, DebugRequest)
	assert.Nil(t, err, "DoRequest should be not null ")
	assert.NotNil(t, body)
}
