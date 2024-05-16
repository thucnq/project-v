package httprequest

import (
	"encoding/base64"
	"io"
	"net/url"
	"path"
	"strings"
)

func closeReader(r io.Reader) bool {
	rc, ok := r.(io.ReadCloser)
	if ok {
		rc.Close()
	}

	return ok
}

func basicAuth(username, password string) string {
	auth := username + ":" + password
	return base64.StdEncoding.EncodeToString([]byte(auth))
}

func ResolveURL(baseURL, urlPath string) (string, error) {
	if strings.HasPrefix(urlPath, "http") {
		return urlPath, nil
	}
	urlObj, err := url.Parse(baseURL)
	if err != nil {
		return "", err
	}
	urlObj.Path = path.Join(urlObj.Path, urlPath)
	finalURL := urlObj.String()
	return finalURL, nil
}
