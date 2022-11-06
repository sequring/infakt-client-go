package infakt_test

import (
	"fmt"
	infakt "github.com/sequring/infakt-client-go"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestClientModel(t *testing.T) {
	client := GetInfactClient()
	if assert.NotNil(t, client) {
		t.Log("Infakt client initialized")
	}

	count, err := client.GetCountAllClient()
	if err != nil {
		t.Fatal("Error get count clients ", err)
	}

	assert.Equal(t, count, 3)

	var clients []infakt.Client
	clients,err = client.GetAllClient(0,0)
	if err != nil {
		t.Fatal("err get all clients:", err)
	}
	fmt.Println("clients:" , clients)
}
