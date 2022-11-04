package infact_test

import (
	"fmt"
	"net/http"
	"os"
	"testing"

	infact "github.com/sequring/infakt-client-go"
	"github.com/stretchr/testify/assert"
)

func TestSomething(t *testing.T) {

	// assert equality
	assert.Equal(t, 123, 123, "they should be equal")
	// assert inequality
	assert.NotEqual(t, 123, 456, "they should not be equal")

}

func TestInfactClient(t *testing.T) {
	var client *infact.Client
	token := os.Getenv("INFAKT_TOKEN")
	var host string = "https://api.infakt.pl/v3"
	client, _ = infact.NewClient(&host, &token)
	if assert.NotNil(t, client) {
		t.Log("Infact client initialized")
	}

	assert.Equal(t, client.InfaktEndpoint, host)
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/clients.json", client.InfaktEndpoint), nil)
	if err != nil {
		t.Fatal("Error new Request", err)
	}

	body, err := client.DoRequest(req, &token)
	if err != nil {
		t.Fatal("Error request", err)
	}

	fmt.Println(string(body[:]))
}
