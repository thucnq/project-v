package httprequest

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"project-v/pkg/httprequest/trace"

	jsoniter "github.com/json-iterator/go"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

type Response struct {
	*http.Response
	content   []byte
	err       error
	debugData *trace.Debug
}

func (resp *Response) GetResponseAs(data interface{}) HTTPResponse {
	if resp.err != nil || resp.Response == nil {
		return resp
	}
	if resp.content == nil {
		resp.content = resp.readAll()
	}
	if err := json.Unmarshal(resp.content, &data); err != nil {
		resp.err = err
	}
	return resp
}

func (resp *Response) GetResponseStatusCodeAs(httpStatusCode *int) HTTPResponse {
	if resp.err != nil {
		return resp
	}
	*httpStatusCode = resp.StatusCode
	return resp
}

func (resp *Response) PrintDebug(logger ILogger) HTTPResponse {
	if resp.debugData != nil {
		bb, err := json.Marshal(resp.debugData)
		if err != nil {
			return resp
		}
		logger.Debug(string(bb))
	}
	return resp
}

var _ HTTPResponse = (*Response)(nil)

func (resp *Response) MustHaveHeader(key, val string) HTTPResponse {
	if resp.err != nil {
		return resp
	}
	if resp.Response == nil {
		resp.err = fmt.Errorf("response nil")
		return resp
	}
	if resp.Header.Get(key) != val {
		resp.err = fmt.Errorf("header %v must be %v", key, val)
	}
	return resp
}

func (resp *Response) MustHaveStatus(status int) HTTPResponse {
	if resp.err != nil {
		return resp
	}
	if resp.Response == nil {
		resp.err = fmt.Errorf("response nil")
		return resp
	}
	if resp.StatusCode != status {
		resp.err = fmt.Errorf(
			"status must be %v, got %v. res: %s", status, resp.StatusCode,
			resp.content,
		)
	}
	return resp
}

func (resp *Response) Error() error {
	return resp.err
}

func (resp *Response) Content(data *[]byte) HTTPResponse {
	if resp.err != nil || resp.Response == nil {
		return resp
	}
	if resp.content == nil {
		*data = resp.readAll()
	}
	return resp
}

func (resp *Response) Text(data *string) HTTPResponse {
	if resp.err != nil || resp.Response == nil {
		return resp
	}
	if resp.content == nil {
		resp.content = resp.readAll()
	}
	*data = string(resp.content)
	return resp
}

func (resp *Response) Json(data interface{}) HTTPResponse {
	if resp.err != nil || resp.Response == nil {
		return resp
	}
	contentType := resp.Header.Get("Content-Type")
	if strings.Contains(contentType, "application/json") {
		if resp.content == nil {
			resp.content = resp.readAll()
		}
		if err := json.Unmarshal(resp.content, &data); err != nil {
			resp.err = err
		}
	}
	return resp
}

func (resp *Response) Headers(header http.Header) HTTPResponse {
	header = resp.Header
	return resp
}
func (resp *Response) GetResponse() *http.Response {
	return resp.Response
}

func (resp *Response) readAll() (content []byte) {
	var err error
	if resp.Response == nil {
		return
	}
	content, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		resp.err = err
		return
	}
	resp.content = content
	defer func() {
		_ = resp.Body.Close()
	}()
	return
}
