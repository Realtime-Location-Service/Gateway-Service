package http

import (
	shttp "net/http"
	"path/filepath"
	"strings"
)

// BuildQueryString builds query string
func BuildQueryString(req *shttp.Request, queries map[string]string) {
	q := req.URL.Query()
	for k, v := range queries {
		q.Add(k, v)
	}
	req.URL.RawQuery = q.Encode()
}

// BuildURL removes service name from path
// then joins with base url
// path => ping/v1/locations?q=whatever
// returns http://baseurl.com/v1/locations?q=whatever
func BuildURL(baseURL, path, prefix string) string {
	s := strings.Split(path, prefix)
	if len(s) > 1 {
		s[0] = prefix
	}
	return baseURL + filepath.Join("/", strings.Join(s, "/"))
}
