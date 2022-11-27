package infakt

type infakt interface {
	GetCountAllClient() (int, error)
	GetAllClient(int, int) ([]Client, error)
	GetClient(id int) (Client, error)
	NewClient() Client
	CreateClient(Client) error
	UpdateClient(Client) error
	DeleteClient(Client) error
}
