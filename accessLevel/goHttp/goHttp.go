package gohttp

type HttpClient struct {
	Password string
}

func (c *HttpClient) Get()  {}
func (c *HttpClient) Post() {}

type httpClient struct {
	Password string
}

func New() ClientInterface {
	client := httpClient{Password: "hoge"}
	return &client
}

type ClientInterface interface {
	Get()
	Post()
}

func (c *httpClient) Get()  {}
func (c *httpClient) Post() {}
