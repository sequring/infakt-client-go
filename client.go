package infakt

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func (c *InFaktClient) GetCountAllClient() (int, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/clients.json%s", c.InfaktEndpoint, ""), nil)
	if err != nil {
		log.Fatal("[client|GetAllClient] Error new Request", err)
	}

	body, err := c.doRequest(req)
	var res ClientRes
	err = json.Unmarshal(body, &res)
	if err != nil {
		return 0, err
	}
	counter := res.MetaInfo.Count
 	return counter, nil
}

func (c *InFaktClient) GetAllClient(offset int, limit int) ([]Client, error) {
	//var clients []Client
	if limit == 0 {
		limit = 10
	}
	pager := fmt.Sprintf("?offset=%d&limit=%d", offset, limit)
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/clients.json%s", c.InfaktEndpoint, pager), nil)
	if err != nil {
		log.Fatal("[client|GetAllClient] Error new Request", err)
	}

	body, err := c.doRequest(req)
	if err != nil {
		log.Fatal("[client|GetAllClient] Error request", err)
	}
	var res ClientRes
	err = json.Unmarshal(body, &res)
	if err != nil {
		return nil, err
	}
	return res.Clients, nil
}
