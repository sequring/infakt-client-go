package infakt

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func (c *InFaktClient) GetCountAllClient() (int, error) {
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/clients.json%s", c.InfaktEndpoint, ""), nil)
	if err != nil {
		return 0, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return 0, err
	}
	var res ClientRes
	err = json.Unmarshal(body, &res)
	if err != nil {
		return 0, err
	}
	counter := res.MetaInfo.Count
	return counter, nil
}

// GET /v3/clients.json
func (c *InFaktClient) GetAllClient(offset int, limit int) ([]Client, error) {
	//var clients []Client
	if limit == 0 {
		limit = 10
	}
	pager := fmt.Sprintf("?offset=%d&limit=%d", offset, limit)
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/clients.json%s", c.InfaktEndpoint, pager), nil)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}
	var res ClientRes
	err = json.Unmarshal(body, &res)
	if err != nil {
		return nil, err
	}
	return res.Clients, nil
}

// GET /v3/clients/{id}.json
func (c *InFaktClient) GetClient(id int) (Client, error) {

	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/clients/%d.json", c.InfaktEndpoint, id), nil)
	if err != nil {
		log.Fatal("[client|GetClient] Error new Request", err)
	}
	body, err := c.doRequest(req)
	if err != nil {
		log.Fatal("[client|GetAllClient] Error request", err)
	}
	var res Client
	err = json.Unmarshal(body, &res)
	if err != nil {
		return res, err
	}
	return res, nil
}

func (c *InFaktClient) NewClient() Client {
	return Client{}
}

// POST /v3/clients.json
func (c *InFaktClient) CreateClient(client Client) error {
	clientReq := NewClientReq{Client: client}
	newClient, err := json.MarshalIndent(clientReq, "", " ")
	if err != nil {
		return err
	}
	bodyReader := bytes.NewReader(newClient)
	req, _ := http.NewRequest(http.MethodPost, fmt.Sprintf("%s/clients.json", c.InfaktEndpoint), bodyReader)
	if err != nil {
		return err
	}
	_, err = c.doRequest(req)

	if err != nil {
		return err
	}
	return nil
}

// PUT /v3/clients/{id}.json
func (c *InFaktClient) UpdateClient(client Client) error {
	clientReq := NewClientReq{Client: client}
	newClient, err := json.Marshal(clientReq)
	if err != nil {
		return err
	}
	bodyReader := bytes.NewReader(newClient)
	req, err := http.NewRequest(http.MethodPut, fmt.Sprintf("%s/clients/%d.json", c.InfaktEndpoint, client.ID), bodyReader)
	if err != nil {
		return err
	}
	_, err = c.doRequest(req)

	if err != nil {
		return err
	}
	return nil
}

// DELETE /v3/clients/{id}.json
func (c *InFaktClient) DeleteClient(client Client) error {
	req, err := http.NewRequest(http.MethodDelete, fmt.Sprintf("%s/clients/%d.json", c.InfaktEndpoint, client.ID), nil)
	if err != nil {
		return err
	}
	_, err = c.doRequest(req)

	if err != nil {
		return err
	}
	return nil
}
