package httprequest

import (
	"context"
	"io"
	"net/http"
	"net/url"
)

// HTTPRequest...
type HTTPRequest interface {
	WithBaseURL(base string) HTTPRequest
	// WithContext
	WithContext(context context.Context) HTTPRequest
	// WithBasicAuth sets the request's Authorization to use Basic Auth
	WithBasicAuth(username string, password string) HTTPRequest
	// WithOauth sets the request's Authorization to use Bearer Auth
	WithOauth(token string) HTTPRequest
	// AddHeaders sets the key, value pair to the header
	AddHeaders(key string, value string) HTTPRequest
	// WithQuery sets the URL-encoded query string
	WithQuery(params *url.Values) HTTPRequest
	// WithJson sets the body request in a json format
	WithJson(data interface{}) HTTPRequest
	WithBody(b io.Reader) HTTPRequest
	// WithJson sets the body request in a form-url-encoded format
	WithForm(data *url.Values) HTTPRequest
	Debug(ILogger) HTTPRequest
	// Get make a request with GET method
	Get(url string) HTTPResponse
	// Post
	Post(url string) HTTPResponse
	// Put
	Put(url string) HTTPResponse
	// Patch
	Patch(url string) HTTPResponse
	// Delete
	Delete(url string) HTTPResponse
}

// HTTPResponse ...
type HTTPResponse interface {
	// MustHaveHeader check the header response must have key, value pair
	MustHaveHeader(key, val string) HTTPResponse
	// MustHaveStatus
	MustHaveStatus(status int) HTTPResponse

	Error() error
	// Content
	Content(data *[]byte) HTTPResponse
	// Text
	Text(data *string) HTTPResponse
	// Json will unmarshal the body response into data
	Json(data interface{}) HTTPResponse
	Headers(header http.Header) HTTPResponse
	GetResponse() *http.Response
	PrintDebug(ILogger) HTTPResponse
	GetResponseAs(data interface{}) HTTPResponse
	GetResponseStatusCodeAs(httpStatusCode *int) HTTPResponse
}

// ILogger the interface provide the methods to send log data by type
type ILogger interface {
	Debug(mgs string)
}

type RestClient interface {
	NewRequest() HTTPRequest
}

type restClient struct {
	httpClient *HttpClient
}

func (r restClient) NewRequest() HTTPRequest {
	return NewRequestImpl(r.httpClient)
}

var _ RestClient = (*restClient)(nil)

func NewRestClient(httpClient *HttpClient) *restClient {
	return &restClient{
		httpClient: httpClient,
	}
}
