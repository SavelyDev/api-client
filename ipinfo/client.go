package ipinfo

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

type Client struct {
	client    *http.Client
	API_TOKEN string
}

func NewClient(timeout time.Duration, token string) (*Client, error) {
	if timeout == 0 {
		return nil, errors.New("timeout can't be zero")
	}

	return &Client{
		client: &http.Client{
			CheckRedirect: func(req *http.Request, via []*http.Request) error {
				fmt.Println(req.Response.Status)
				fmt.Println("REDIRECT")
				fmt.Println(req.Response.Location())
				return nil
			},
			Transport: &loggingRoundTripper{
				logger: os.Stdout,
				next:   http.DefaultTransport,
			},
			Timeout: timeout,
		},
		API_TOKEN: token,
	}, nil
}

func (c *Client) GetIPInfo(addr string) (IPInfo, error) {
	res, err := c.client.Get(fmt.Sprintf("https://ipinfo.io/%s?token=%s", addr, c.API_TOKEN))
	if err != nil {
		return IPInfo{}, err
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return IPInfo{}, err
	}
	defer res.Body.Close()

	var result IPInfo
	if err = json.Unmarshal(body, &result); err != nil {
		return IPInfo{}, err
	}

	return result, nil
}
