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
	client, _ = infakt.NewInFaktClient(&host, &token)
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
	if assert.NotNil(t, client) {
		t.Log("Infakt client initialized")
	}

	assert.Equal(t, client.InfaktEndpoint, host)
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/clients.json", client.InfaktEndpoint), nil)
	if err != nil {
		t.Fatal("Error new Request", err)
	}

	body, err := infakt.DoRequest(client, req, host, &token, DebugRequest)
	if err != nil {
		t.Fatal("Error request", err)
	}

	assert.NotNil(t, body)
}
