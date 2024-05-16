
## http V2

### How to integrate?

1. Include the module in the project.
2. Create a new instance of the http client.


```go

func main() {
  httpClient := httprequest.NewClient()
  
  rc := httprequest.NewRestClient(httpClient)
  
  q := &url.Values{}
  q.Set("q", "test")

  headerKey := "X-Test"
  headerVal := "value1"
  bodyKey := "data"
  bodyVal := "hello"
  
  var data string
  err := rc.NewRequest().
    AddHeaders(headerKey, headerVal).
    WithQuery(q).
    Get("https://webhook.site/2e98de81-9354-4263-88d8-42709f11bbfc").
    MustHaveStatus(http.StatusOK).
    Text(&data). // response as text
    Error()
  if err != nil {
    // todo: handle error    
    panic(err)
  }
  
  // todo: handle data   
  println(data)
}
```

Make post request with query params and body request in json format

```go
data := make(map[string]string)
err := rc.NewRequest().
    AddHeaders(headerKey, headerVal).
    WithJson(map[string]string{bodyKey: bodyVal}).
    WithQuery(q).
    Post(s.URL).
    MustHaveStatus(http.StatusOK). 
    MustHaveHeader(headerKey, headerVal). // validate response header
    Json(&data). // response as struct. data can be instance of struct. the body response will be unmarshal to this struct
    Error()
```

Make post request with form url encoded

```go
data := make(map[string]string)
err := rc.NewRequest().
    AddHeaders(headerKey, headerVal).
    WithForm(q).
    Post(s.URL).
    MustHaveStatus(http.StatusOK). 
    MustHaveHeader(headerKey, headerVal).
    Json(&data). 
    Error()
```


Make patch request

```go
var data []byte
err := rc.NewRequest().
    AddHeaders(headerKey, headerVal).
    WithJson(map[string]string{bodyKey: bodyVal}).
    Patch(s.URL).
    MustHaveStatus(http.StatusOK).
    Content(&data). // response as byte array
    Error()
```

Make put request

```go
err := rc.NewRequest().
    AddHeaders(headerKey, headerVal).
    WithJson(map[string]string{bodyKey: bodyVal}).
    Patch(s.URL).
    MustHaveStatus(http.StatusOK). 
    Json(&data).
    Error()
```

Make delete request

```go
err := rc.NewRequest().
    AddHeaders(headerKey, headerVal).
    WithQuery(q).
    Delete(s.URL).
    MustHaveStatus(http.StatusOK). 
    Error()
```



