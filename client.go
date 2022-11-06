package infakt

import (
	"fmt"
)

func (c *InFaktClient) GetAllClient(offset int, limit int) ([]Client, error) {
	var clients []Client
	if limit == 0 {
		limit = 10
	}
	fmt.Println("offset:", offset, " ", "limit:", limit)
	return clients, nil
}
