package http_client

import "net/url"

type UrlOpts func(*url.URL) *url.URL

func WithPath(path string) UrlOpts {
	return func(u *url.URL) *url.URL {
		u.Path = path
		return u
	}
}

func WithQueries(queries map[string]string) UrlOpts {
	return func(u *url.URL) *url.URL {
		q := u.Query()

		for key, value := range queries {
			q.Set(key, value)
		}

		u.RawQuery = q.Encode()
		return u
	}
}

func GetUrl(host string, opts ...UrlOpts) *url.URL {
	u := &url.URL{}
	u.Scheme = "https"
	u.Host = host

	for _, opt := range opts {
		u = opt(u)
	}

	return u
}
