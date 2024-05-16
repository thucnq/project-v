package httprequest

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestNewRestClientGet(t *testing.T) {
	headerKey := "X-Test"
	headerVal := "value1"

	s := httptest.NewServer(
		http.HandlerFunc(
			func(w http.ResponseWriter, r *http.Request) {
				if r.Method != http.MethodGet {
					w.WriteHeader(http.StatusMethodNotAllowed)
					return
				}
				if val := r.Header.Get(headerKey); len(val) > 0 {
					w.Header().Set(headerKey, val)
				}
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusOK)
				w.Write([]byte(fmt.Sprintf(`{"q": "%v"}`, r.URL.RawQuery)))
			},
		),
	)
	defer s.Close()

	httpClient := NewClient()
	rc := NewRestClient(httpClient)

	q := &url.Values{}
	q.Set("key1", "val1")

	data := make(map[string]string)
	err := rc.NewRequest().
		AddHeaders(headerKey, headerVal).
		WithQuery(q).
		Get(s.URL).
		MustHaveStatus(http.StatusOK).
		MustHaveHeader(headerKey, headerVal).
		Json(&data).Error()
	if err != nil {
		t.Fatal(err)
	}

	val, ok := data["q"]
	if !ok {
		t.Fatalf("got data without `q` field, want `q` field")
	}

	if val != q.Encode() {
		t.Fatalf("got %v, want %v", val, q.Encode())
	}
}

func TestNewRestClientPost(t *testing.T) {
	headerKey := "X-Test"
	headerVal := "value1"
	mockUser := "user"
	mockPassword := "user"
	headerAuthVal := basicAuth(mockUser, mockPassword)
	bodyKey := "data"
	bodyVal := "hello"

	s := httptest.NewServer(
		http.HandlerFunc(
			func(w http.ResponseWriter, r *http.Request) {
				if r.Method != http.MethodPost {
					w.WriteHeader(http.StatusMethodNotAllowed)
					return
				}
				if r.Header.Get("Authorization") != "Basic "+headerAuthVal {
					w.WriteHeader(http.StatusUnauthorized)
					return
				}
				if val := r.Header.Get(headerKey); len(val) > 0 {
					w.Header().Set(headerKey, val)
				}
				w.Header().Set("Content-Type", "application/json")
				bb, err := ioutil.ReadAll(r.Body)
				if err != nil {
					w.WriteHeader(http.StatusInternalServerError)
					return
				}
				w.WriteHeader(http.StatusOK)
				w.Write(bb)
			},
		),
	)
	defer s.Close()

	httpClient := NewClient()
	rc := NewRestClient(httpClient)

	data := make(map[string]string)
	err := rc.NewRequest().
		AddHeaders(headerKey, headerVal).
		WithBasicAuth(mockUser, mockPassword).
		WithJson(map[string]string{bodyKey: bodyVal}).
		Post(s.URL).
		MustHaveStatus(http.StatusOK).
		MustHaveHeader(headerKey, headerVal).
		Json(&data).Error()
	if err != nil {
		t.Fatal(err)
	}

	val, ok := data[bodyKey]
	if !ok {
		t.Fatalf("got data without `%v` field, want `%v` field", bodyKey, bodyKey)
	}

	if val != bodyVal {
		t.Fatalf("got %v, want %v", val, bodyVal)
	}
}

func TestNewRestClientPatch(t *testing.T) {
	headerKey := "X-Test"
	headerVal := "value1"
	headerToken := "mock jwt token"
	bodyKey := "data"
	bodyVal := "hello"

	s := httptest.NewServer(
		http.HandlerFunc(
			func(w http.ResponseWriter, r *http.Request) {
				if r.Method != http.MethodPatch {
					w.WriteHeader(http.StatusMethodNotAllowed)
					return
				}
				if r.Header.Get("Authorization") != "Bearer "+headerToken {
					w.WriteHeader(http.StatusUnauthorized)
					return
				}
				if val := r.Header.Get(headerKey); len(val) > 0 {
					w.Header().Set(headerKey, val)
				}
				w.Header().Set("Content-Type", "application/json")
				bb, err := ioutil.ReadAll(r.Body)
				if err != nil {
					w.WriteHeader(http.StatusInternalServerError)
					return
				}
				w.WriteHeader(http.StatusOK)
				w.Write(bb)
			},
		),
	)
	defer s.Close()

	httpClient := NewClient()
	rc := NewRestClient(httpClient)

	data := make(map[string]string)
	err := rc.NewRequest().
		AddHeaders(headerKey, headerVal).
		WithOauth(headerToken).
		WithJson(map[string]string{bodyKey: bodyVal}).
		Patch(s.URL).
		MustHaveStatus(http.StatusOK).
		MustHaveHeader(headerKey, headerVal).
		Json(&data).Error()
	if err != nil {
		t.Fatal(err)
	}

	val, ok := data[bodyKey]
	if !ok {
		t.Fatalf("got data without `%v` field, want `%v` field", bodyKey, bodyKey)
	}

	if val != bodyVal {
		t.Fatalf("got %v, want %v", val, bodyVal)
	}
}

func TestNewRestClientPut(t *testing.T) {
	headerKey := "X-Test"
	headerVal := "value1"
	headerToken := "mock jwt token"
	bodyKey := "data"
	bodyVal := "hello"

	s := httptest.NewServer(
		http.HandlerFunc(
			func(w http.ResponseWriter, r *http.Request) {
				if r.Method != http.MethodPut {
					w.WriteHeader(http.StatusMethodNotAllowed)
					return
				}
				if r.Header.Get("Authorization") != "Bearer "+headerToken {
					w.WriteHeader(http.StatusUnauthorized)
					return
				}
				if val := r.Header.Get(headerKey); len(val) > 0 {
					w.Header().Set(headerKey, val)
				}
				w.Header().Set("Content-Type", "application/json")

				r.ParseForm()
				w.WriteHeader(http.StatusOK)
				w.Write([]byte(fmt.Sprintf(`{"%v": "%v"}`, bodyKey, r.PostForm.Encode())))

			},
		),
	)
	defer s.Close()

	httpClient := NewClient()
	rc := NewRestClient(httpClient)
	q := &url.Values{}
	q.Set(bodyKey, bodyVal)
	data := make(map[string]string)
	err := rc.NewRequest().
		AddHeaders(headerKey, headerVal).
		WithOauth(headerToken).
		WithForm(q).
		Put(s.URL).
		MustHaveStatus(http.StatusOK).
		MustHaveHeader(headerKey, headerVal).
		Json(&data).Error()
	if err != nil {
		t.Fatal(err)
	}

	val, ok := data[bodyKey]
	if !ok {
		t.Fatalf("got data without `%v` field, want `%v` field", bodyKey, bodyKey)
	}

	if val != q.Encode() {
		t.Fatalf("got %v, want %v", val, q.Encode())
	}
}

func TestNewRestClientDelete(t *testing.T) {
	headerKey := "X-Test"
	headerVal := "value1"
	bodyKey := "data"
	bodyVal := "hello"

	s := httptest.NewServer(
		http.HandlerFunc(
			func(w http.ResponseWriter, r *http.Request) {
				if r.Method != http.MethodDelete {
					w.WriteHeader(http.StatusMethodNotAllowed)
					return
				}
				if val := r.Header.Get(headerKey); len(val) > 0 {
					w.Header().Set(headerKey, val)
				}
				w.Header().Set("Content-Type", "application/json")

				w.WriteHeader(http.StatusOK)

			},
		),
	)
	defer s.Close()
	ctx := context.TODO()

	httpClient := NewClient()
	rc := NewRestClient(httpClient)
	q := &url.Values{}
	q.Set(bodyKey, bodyVal)
	err := rc.NewRequest().
		WithContext(ctx).
		AddHeaders(headerKey, headerVal).
		Delete(s.URL).
		MustHaveStatus(http.StatusOK).
		MustHaveHeader(headerKey, headerVal).
		Error()
	if err != nil {
		t.Fatal(err)
	}
}
