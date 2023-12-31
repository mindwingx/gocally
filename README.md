# GO Cally

GO Cally is a minimal HTTP/REST Caller for Golang, designed to be easy to use and fast.

### Features

- Simple
- Easy to use
- Fast Forward

### Installation

Install GO Cally using the following command:

```
go get github.com/mindwingx/gocally
```

### Usage

Instantiate the HTTP call object with the following code:

```go
request := gocally.SetRequest()
```
Prepare URL:
```go
response, error := request.WithUrl("https://dummy-url.io/api")
```
Set header using the following code:

```go
response, error := request.
                    WithUrl("https://dummy-url.io/api").
                    SetHeader("Header", "Value")
```
or for multiple headers:
```go
response, error := request.
                    WithUrl("https://dummy-url.io/api").
                    SetHeaderBulk(map[string]string{
                        "Header1": "Value1",
                        "Header2": "Value2",
                    })
```
Set authorization using the following code:
```go
response, error := request.
                    WithUrl("https://dummy-url.io/api").
	            SetAuthorization("Bearer dummy-token")
```
By default, the following headers are set:
```go
req.Header.Set("Accept", "application/json")
req.Header.Set("Content-Type", "application/json")
```
To disable setting these headers, use the following code:
```go
response, error := request.
                    WithUrl("https://dummy-url.io/api").
                    DisableJsonHeaders()
```

Set query parameters, single or bulk, using the following code:
```go
response, error := request.
                    WithUrl("https://dummy-url.io/api").
                    SetQueryParam("sort", "desc")
```
or
```go
response, error := request.
                    WithUrl("https://dummy-url.io/api").
                    SetQueryParamBulk(map[string]string{
                        "page":  "1",
                        "limit": "10",
                   })
```

Set the request body using type of `any`, commonly a `map[string]interface{}`:

```go
response, error := request.
                    WithUrl("https://dummy-url.io/api").
                    SetBody(map[string]interface{
						"field": "value"
					})
```
Set the request `--form` for multipart requests using the following code:

```go
response, error := request.
                    WithUrl("https://dummy-url.io/api").
                    SetForm("form-key", "value")
```
or for multiple form values:

```go
response, error := request.
                    WithUrl("https://dummy-url.io/api").
                    SetFormBulk(map[string]string{
                        "form-key1": "value1",
                        "form-key2": "value2",
                    })
```
Set the request timeout in seconds (default value is `10`) using the following code:

```go
response, error := request.
                    WithUrl("https://dummy-url.io/api").
                    SetRequestTimeout(30)
```

Set the HTTP methods:

`GET`
```go
response, error := request.
                    WithUrl("https://dummy-url.io/api").
                    Get()
```
`POST`
```go
response, error := request.
                    WithUrl("https://dummy-url.io/api").
                    Post()
```
`PUT`
```go
response, error := request.
                    WithUrl("https://dummy-url.io/api").
                    Put()
```
`DELETE`
```go
response, error := request.
                    WithUrl("https://dummy-url.io/api").
                    Delete()
```
##### Response is available with three methods :

`Response()` : this method returns exact Http Response with no changes. 

`Payload()` : this method returns `status`,`status_code` and `payload`(as `map[string]interface{}`). 

`Entity(&DummyStruct)` : this method receives compatible `struct` with response and returns `status`,`status_code`(as `map[string]interface{}`). 

### Example

- request

```go
request := SetRequest().
        SetRequestTimeout(30).
        WithUrl("https://dummy-url.io/api").
        SetHeader("header","header-value").
        SetQueryParam("order","desc").
        Get()
```

- response methods

```go
response, error := request.Response()
```

```
response, error := request.Payload()
```

```
type DummyResponse struct {
	Username    string `json:"username"`
	IsSatisfied bool   `json:"is_satisfied"`
}

response, error := request.Entity(&DummyResponse)
```

the `DummyResponse` fields will be filled base on response body.

### Contributing

Contributions are welcome! If you find any issues or have suggestions for improvement, please submit an issue or a pull
request on the GitHub repository.

### License

The GO Cally package is open-source software licensed under the MIT license.

### Credits

The GO Cally package is developed and maintained by Milad Roudgarian.
