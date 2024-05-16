package httprequest

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"net/http/httptrace"
	"net/http/httputil"
	"net/url"
	"path"
	"strings"

	"project-v/pkg/httprequest/trace"
)

type Request struct {
	*http.Request

	client  *HttpClient
	err     error
	body    io.Reader
	params  string
	mime    string
	cookies []*http.Cookie
	ctx     context.Context
	headers http.Header

	clientTrace *httptrace.ClientTrace
	debugData   *trace.Debug
	l           ILogger
	baseURL     string
}

func (r Request) WithBaseURL(base string) HTTPRequest {
	r.baseURL = base
	return r
}

func (r Request) Debug(l ILogger) HTTPRequest {
	r.clientTrace, r.debugData = trace.TraceHTTP()
	r.l = l
	return r
}

func (r Request) AddHeaders(key string, value string) HTTPRequest {
	r.headers.Add(key, value)
	return r
}

var _ HTTPRequest = (*Request)(nil)

func NewRequestImpl(client *HttpClient) *Request {
	return &Request{client: client, headers: http.Header{}}
}

func (r Request) WithQuery(params *url.Values) HTTPRequest {
	r.params = params.Encode()
	return r
}

func (r Request) WithForm(data *url.Values) HTTPRequest {
	r.mime = "application/x-www-form-urlencoded"
	r.body = strings.NewReader(data.Encode())
	return r
}

func (r Request) WithJson(data interface{}) HTTPRequest {
	bodyBytes, err := json.Marshal(data)
	if err != nil {
		r.err = err
		return r
	}
	r.mime = "application/json"
	r.body = bytes.NewReader(bodyBytes)
	return r
}

func (r Request) WithBody(b io.Reader) HTTPRequest {
	r.body = b
	return r
}

func (r Request) WithOauth(token string) HTTPRequest {
	r.headers.Set("Authorization", "Bearer "+token)
	return r
}

func (r Request) WithBasicAuth(username string, password string) HTTPRequest {
	r.headers.Set("Authorization", "Basic "+basicAuth(username, password))
	return r
}
func (r Request) WithContext(ctx context.Context) HTTPRequest {
	r.ctx = ctx
	return r
}

func (r Request) Get(uri string) HTTPResponse {
	return r.Do("GET", uri)
}

func (r Request) Post(uri string) HTTPResponse {
	return r.Do("POST", uri)
}

func (r Request) Put(uri string) HTTPResponse {
	return r.Do("PUT", uri)
}

func (r Request) Patch(uri string) HTTPResponse {
	return r.Do("PATCH", uri)
}

func (r Request) Delete(uri string) HTTPResponse {
	return r.Do("DELETE", uri)
}

func (r Request) resolveURL(urlPath string) (string, error) {
	if strings.HasPrefix(urlPath, "http") {
		return urlPath, nil
	}
	if len(r.baseURL) == 0 {
		return urlPath, nil
	}

	urlObj, err := url.Parse(r.baseURL)
	if err != nil {
		return "", fmt.Errorf("invalid url: %v", r.baseURL)
	}
	urlObj.Path = path.Join(urlObj.Path, urlPath)
	finalURL := urlObj.String()
	return finalURL, nil
}

func (r Request) Do(method string, uri string) HTTPResponse {
	var err error
	defer closeReader(r.body)
	if r.err != nil {
		return &Response{nil, nil, r.err, nil}
	}

	uri, err = r.resolveURL(uri)
	if err != nil {
		return &Response{nil, nil, r.err, nil}
	}
	(&r).prepareRequest(method, uri)
	if r.err != nil {
		return &Response{nil, nil, r.err, nil}
	}

	if r.clientTrace != nil {
		clientTraceCtx := httptrace.WithClientTrace(
			r.Request.Context(), r.clientTrace,
		)
		r.Request = r.Request.WithContext(clientTraceCtx)

		dump, err := httputil.DumpRequestOut(r.Request, true)
		if err != nil {
			r.l.Debug(err.Error())
		} else {
			r.l.Debug(string(dump))
		}
	}
	resp, err := r.client.Do(r.Request)

	hr := &Response{resp, nil, err, r.debugData}
	hr.readAll()
	return hr
}

func (r *Request) prepareRequest(method string, uri string) {
	var err error
	body, ok := r.body.(io.ReadCloser)
	if !ok && r.body != nil {
		body = io.NopCloser(r.body)
	}

	if len(r.mime) > 0 {
		r.headers.Set("Content-Type", r.mime)
	}

	if r.ctx != nil {
		r.Request, err = http.NewRequestWithContext(r.ctx, method, uri, body)
	} else {
		r.Request, err = http.NewRequest(method, uri, body)
	}
	if err != nil {
		r.err = err
		return
	}

	r.Request.URL.RawQuery = r.params
	r.Request.Header = r.headers
}
