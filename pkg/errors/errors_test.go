package errors_test

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"

	customerrors "project-v/pkg/errors"
	"project-v/pkg/l"
)

var ll = l.New()

func TestErrorInterface(t *testing.T) {
	err := errors.New("ABC")
	if _, ok := err.(customerrors.IError); !ok {
		t.Error("Must satisfy errors interface")
	}
}

func TestError(t *testing.T) {
	t.Run(
		"Simple", func(t *testing.T) {
			err := customerrors.Error(customerrors.NotFound, "Foo")

			v := AssertJSONError(t, err)
			assert.Equal(
				t, &ErrorJSON{
					Code: "NotFound",
					Msg:  "Foo",
					Logs: json.RawMessage([]byte(`[]`)),
				}, v,
			)
		},
	)

	t.Run(
		"Simple with trace", func(t *testing.T) {
			err := customerrors.ErrorTrace(customerrors.NotFound, "Foo")

			err2 := customerrors.ErrorTrace(customerrors.NotFound, "Foo 2", err)

			err3 := customerrors.ErrorTrace(
				customerrors.WrongPassword, "Foo 3", err2,
			)

			// fmt.Println(err3.(*APIError).ErrorStack())

			_ = AssertJSONError(t, err3)
			fmt.Printf("%#x", err3)

			v := AssertJSONError(t, err)
			assert.Equal(
				t, &ErrorJSON{
					Code:  "Unauthenticated",
					Msg:   "WRONG_PASSWORD",
					Logs:  json.RawMessage([]byte(`[]`)),
					Stack: v.Stack,
				}, v,
			)

		},
	)
}

type ErrorJSON struct {
	Code  string          `json:"code"`
	XCode string          `json:"xcode"`
	Err   string          `json:"err"`
	Msg   string          `json:"msg"`
	Orig  string          `json:"orig"`
	Logs  json.RawMessage `json:"logs"`
	Stack string          `json:"stack"`
}

func AssertJSONError(t *testing.T, err *customerrors.APIError) *ErrorJSON {
	data, jsonErr := err.MarshalJSON()
	fmt.Printf("--> %s\n", data)
	ll.Error("AssertJSONError", l.Error(err))
	assert.NoError(t, jsonErr)

	data, jsonErr = json.Marshal(err)
	assert.NoError(t, jsonErr)

	var v ErrorJSON
	jsonErr = json.Unmarshal(data, &v)
	if jsonErr != nil {
		t.Errorf("Got error while decoding JSON: %v", jsonErr)
	}
	return &v
}
