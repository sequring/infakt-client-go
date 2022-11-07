package infakt_test

import (
	"crypto/rand"
	"fmt"
	"math/big"

	"reflect"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	infakt "github.com/sequring/infakt-client-go"

	"github.com/stretchr/testify/assert"
)

const id int = 20274698

func TestClient_GetCountAllClient(t *testing.T) {
	client := GetInfactClient()
	count, err := client.GetCountAllClient()
	if err != nil {
		t.Fatal("Error get count clients ", err)
	}
	assert.Equal(t, count, 3)
}

func TestClient_GetAllClient(t *testing.T) {
	client := GetInfactClient()
	var clients []infakt.Client
	clients, err := client.GetAllClient(0, 0)
	if err != nil {
		t.Fatal("err get all clients:", err)
	}
	assert.Equal(t, clients[0].ID, id)
}

func TestClient_GetClient(t *testing.T) {
	c := GetInfactClient()
	client, err := c.GetClient(id)
	if err != nil {
		t.Fatal("err get client by id:", err)
	}
	assert.Equal(t, client.ID, id)
}

func TestClient_NewClient(t *testing.T) {
	c := GetInfactClient()
	client := c.NewClient()
	assert.Equal(t, fmt.Sprint(reflect.TypeOf(client)), "infakt.Client")
}

func CreateTestClient() *infakt.Client {
	fake := gofakeit.NewCrypto()
	c := GetInfactClient()
	client := c.NewClient()
	client.CompanyName = fake.Company()
	client.City = fake.City()
	client.Street = fake.Street()
	client.StreetNumber = fake.StreetNumber()
	client.PostalCode = fake.Zip()
	client.Country = fake.Country()
	client.SameForwardAddress = true
	return &client
}

func TestClient_CreateClient(t *testing.T) {
	//c := GetInfactClient()
	fmt.Println(rand.Int(rand.Reader, big.NewInt(1000)))
	client := CreateTestClient()
	fmt.Println(client)
}
