package quote

import (
	"encoding/json"
	"github.com/eneskzlcn/quote-reverser-client/internal/config"
	"io"
	"net/http"
)

type Client struct {
	config config.Client
}

func NewClient(config config.Client) *Client {
	return &Client{config: config}
}

func (c *Client) ConsumeQuotes() (Quotes, error) {
	resp, err := http.Get(c.config.Url)
	if err != nil {
		return nil, err
	}
	var quotes Quotes
	if err = c.parseBody(resp.Body, &quotes); err != nil {
		return nil, err
	}
	return quotes, nil
}

func (c *Client) parseBody(body io.Reader, target interface{}) error {
	bytes, err := io.ReadAll(body)
	if err != nil {
		return err
	}
	return json.Unmarshal(bytes, target)
}
