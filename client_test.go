package infakt_test

import (
	"fmt"

	"reflect"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	infakt "github.com/sequring/infakt-client-go"

	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
)

const id int = 20274698
const res_all_client = `{"metainfo":{"count":3,"total_count":3,"next":"https://api.infakt.pl/api/v3/clients.json?offset=10&limit=10","previous":"https://api.infakt.pl/api/v3/clients.json?offset=0&limit=10"},"entities":[{"id":20274698,"company_name":"Ewa Ząbkiewicz Magic Optic","street":"ul. Łukowska","street_number":"9","flat_number":"149","city":"Warszawa","country":"PL","postal_code":"04-133","nip":"1132365035","phone_number":"","web_site":"","email":"","note":"","receiver":"","mailing_company_name":"","mailing_street":"","mailing_city":"","mailing_postal_code":"","days_to_payment":"","payment_method":"","invoice_note":"","same_forward_address":true,"first_name":"Ewa","last_name":"Ząbkiewicz","business_activity_kind":"self_employed"},{"id":20274719,"company_name":"LAURA KLUB SPORTOWY CHYLICE","street":"ul. Dworska","street_number":"5","flat_number":"","city":"Chylice","country":"PL","postal_code":"05-510","nip":"1230934563","phone_number":"","web_site":"","email":"","note":"","receiver":"","mailing_company_name":"","mailing_street":"","mailing_city":"","mailing_postal_code":"","days_to_payment":"","payment_method":"","invoice_note":"","same_forward_address":true,"first_name":null,"last_name":null,"business_activity_kind":"other_business"},{"id":20274724,"company_name":"Q5 Maciej Kaszyński","street":"ul. Adama Branickiego","street_number":"11","flat_number":"154","city":"Warszawa","country":"PL","postal_code":"02-972","nip":"1132302833","phone_number":"","web_site":"","email":"","note":"","receiver":"","mailing_company_name":"","mailing_street":"","mailing_city":"","mailing_postal_code":"","days_to_payment":"","payment_method":"","invoice_note":"","same_forward_address":true,"first_name":"Maciej","last_name":"Kaszyński","business_activity_kind":"self_employed"}]}`
const res_client = `{"id":20274698,"company_name":"Ewa Ząbkiewicz Magic Optic","street":"ul. Łukowska","street_number":"9","flat_number":"149","city":"Warszawa","country":"PL","postal_code":"04-133","nip":"1132365035","phone_number":"","web_site":"","email":"","note":"","receiver":"","mailing_company_name":"","mailing_street":"","mailing_city":"","mailing_postal_code":"","days_to_payment":"","payment_method":"","invoice_note":"","same_forward_address":true,"first_name":"Ewa","last_name":"Ząbkiewicz","business_activity_kind":"self_employed"}`

func TestClient_GetCountAllClient(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	httpmock.RegisterResponder(
		"GET",
		"https://api.infakt.pl/v3/clients.json",
		httpmock.NewStringResponder(200, res_all_client),
	)
	client := GetInfactClient()
	count, err := client.GetCountAllClient()
	assert.Nil(t, err, "Err after get client by id should be null ")
	assert.NotNil(t, count, "Client after get count client by id should be not null ")
	assert.Equal(t, count, 3)
}

func TestClient_GetAllClient(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	httpmock.RegisterResponder(
		"GET",
		"https://api.infakt.pl/v3/clients.json",
		httpmock.NewStringResponder(200, res_all_client),
	)
	client := GetInfactClient()
	var clients []infakt.Client
	clients, err := client.GetAllClient(0, 0)
	assert.Nil(t, err, "Err after get all clients by id should be null ")
	assert.NotNil(t, clients, "Client after get all client by id should be not null ")
	assert.Equal(t, clients[0].ID, id)
}

func TestClient_GetClient(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	httpmock.RegisterResponder(
		"GET",
		fmt.Sprintf("https://api.infakt.pl/v3/clients/%d.json", id),
		httpmock.NewStringResponder(200, res_client),
	)
	c := GetInfactClient()
	client, err := c.GetClient(id)
	assert.Nil(t, err, "Err after get client by id should be null ")
	assert.NotNil(t, client, "Client after get client by id should be not null ")
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

func TestClient_UpdateClient(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	httpmock.RegisterResponder(
		"PUT",
		fmt.Sprintf("https://api.infakt.pl/v3/clients/%d.json", id),
		httpmock.NewStringResponder(200, "ok"),
	)
	httpmock.RegisterResponder(
		"GET",
		fmt.Sprintf("https://api.infakt.pl/v3/clients/%d.json", id),
		httpmock.NewStringResponder(200, res_client),
	)
	c := GetInfactClient()
	client, err := c.GetClient(id)
	assert.Nil(t, err, "Err after get client by id should be null ")
	assert.NotNil(t, client, "Client after get client by id should be not null ")
	client.City = "San Diego"
	err = c.UpdateClient(client)
	assert.Nil(t, err, "Err after update client should be null ")
}

func TestClient_CreateClient(t *testing.T) {
	newClient := CreateTestClient()
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	httpmock.RegisterResponder(
		"POST",
		"https://api.infakt.pl/v3/clients.json",
		httpmock.NewStringResponder(201, `{"id":20652124,"company_name":"Avvo","street":null,"street_number":null,"flat_number":null,"city":null,"country":"US","postal_code":null,"nip":null,"phone_number":null,"web_site":"","email":null,"note":"","receiver":"","mailing_company_name":"","mailing_street":"","mailing_city":"","mailing_postal_code":"","days_to_payment":"","payment_method":"","invoice_note":"","same_forward_address":true,"first_name":null,"last_name":null,"business_activity_kind":"other_business"}`),
	)
	c := GetInfactClient()
	err := c.CreateClient(*newClient)
	assert.Nil(t, err, "Err after create client should be null ")
}

func TestClient_DeleteClient(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	httpmock.RegisterResponder(
		"DELETE",
		fmt.Sprintf("https://api.infakt.pl/v3/clients/%d.json", id),
		httpmock.NewStringResponder(204, ""),
	)
	httpmock.RegisterResponder(
		"GET",
		fmt.Sprintf("https://api.infakt.pl/v3/clients/%d.json", id),
		httpmock.NewStringResponder(200, res_client),
	)
	c := GetInfactClient()
	client, err := c.GetClient(id)
	assert.Nil(t, err, "Err after get client by id should be null ")
	assert.NotNil(t, client, "Client after get client by id should be not null ")
	err = c.DeleteClient(client)
	assert.Nil(t, err, "Err after delete client by client should be null ")
}
