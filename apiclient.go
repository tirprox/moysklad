package moysklad

type moyskladClient struct {
	login, password string

}

func NewClient(login, password string) *moyskladClient {
	return &moyskladClient{login, password}
}