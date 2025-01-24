package http_client

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/url"
	"slices"
)

var ErrResponseIsNotJson = errors.New("Error response is not a json")

type HttpClient struct {
	client *http.Client

	cookies []*http.Cookie
}

type ClientOpts func(*http.Client) *http.Client

func New(opts ...ClientOpts) *HttpClient {
	client := &http.Client{}

	for _, opt := range opts {
		client = opt(client)
	}

	return &HttpClient{
		client:  client,
		cookies: make([]*http.Cookie, 0),
	}
}

func (c *HttpClient) updateCookies(cookies []*http.Cookie) {
	cookieNames := make([]string, len(c.cookies))
	for i, cookie := range c.cookies {
		cookieNames[i] = cookie.Name
	}

	for _, cookie := range cookies {
		if !slices.Contains(cookieNames, cookie.Name) {
			c.cookies = append(c.cookies, cookie)
		} else {
			for _, ck := range c.cookies {
				if ck.Name == cookie.Name {
					cookie.Value = ck.Value
				}
			}
		}
	}
}

func (c *HttpClient) Get(u *url.URL, data interface{}) error {
	req, err := http.NewRequest(http.MethodGet, u.String(), nil)
	if err != nil {
		return err
	}

	for _, cookie := range c.cookies {
		req.AddCookie(cookie)
	}

	res, err := c.client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	contentType := res.Header.Get("Content-Type")

	if contentType == "application/json" {
		if err := json.NewDecoder(res.Body).Decode(&data); err != nil {
			return err
		}
	}

	c.updateCookies(res.Cookies())

	return nil
}
