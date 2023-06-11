package gocally

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHttpCall_SuccessRequestWithNotHandledResponseByResponseMethod(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Authorization", "Bearer some-enc-token")

		response := map[string]interface{}{
			"payload": "Request Dummy Payload",
		}
		json.NewEncoder(w).Encode(response)
	}))

	defer server.Close()

	c, err := SetRequest().
		WithUrl(server.URL).
		Get().
		Response()

	defer c.Body.Close()

	assert.NotNil(t, c)
	assert.Nil(t, err)

	var response map[string]interface{}

	err = json.NewDecoder(c.Body).Decode(&response)

	assert.Equal(t, "200 OK", c.Status)
	assert.Equal(t, 200, c.StatusCode)
	assert.Equal(t, "Request Dummy Payload", response["payload"])
}

func TestHttpCall_SuccessRequestWithHandledResponseByPayloadMethod(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Authorization", "Bearer some-enc-token")

		response := map[string]interface{}{
			"status":      "200 OK",
			"status_code": 200,
			"payload":     "Request Dummy Payload",
		}
		json.NewEncoder(w).Encode(response)
	}))

	defer server.Close()

	c, err := SetRequest().
		WithUrl(server.URL).
		Get().
		Payload()

	assert.NotNil(t, c)
	assert.Nil(t, err)
	assert.Equal(t, "200 OK", c["status"])
	assert.Equal(t, 200, c["status_code"])
	assert.Equal(t, "Request Dummy Payload", c["payload"].(map[string]interface{})["payload"])
}

type CustomResponse struct {
	Username    string `json:"username"`
	IsSatisfied bool   `json:"is_satisfied"`
}

func TestHttpCall_SuccessRequestWithCustomStructForResponseOfEntityMethod(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Authorization", "Bearer some-enc-token")

		response := map[string]interface{}{
			"username":     "gocally_user",
			"is_satisfied": true,
		}

		json.NewEncoder(w).Encode(response)
	}))

	defer server.Close()

	var response CustomResponse

	c, err := SetRequest().
		WithUrl(server.URL).
		Get().
		Entity(&response)

	assert.NotNil(t, c)
	assert.Nil(t, err)
	assert.Equal(t, "200 OK", c["status"])
	assert.Equal(t, 200, c["status_code"])
	assert.Equal(t, "gocally_user", response.Username)
	assert.True(t, response.IsSatisfied)
}

func TestHttpCall_FailedRequestWithServerError(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Simulate an error response
		w.WriteHeader(http.StatusInternalServerError)
	}))

	defer server.Close()

	// Make a request to the server
	response, err := http.Get(server.URL)
	assert.Nil(t, err)
	assert.NotNil(t, response)
	assert.Equal(t, 500, response.StatusCode)
	assert.Equal(t, "500 Internal Server Error", response.Status)
}

func TestHttpCall_NoWithUrlFailure(t *testing.T) {
	c, err := SetRequest().Get().Response()

	assert.Error(t, err, "error: no URL is set")
	assert.Nil(t, c.Request)
}

func TestHttpCall_SetAuthorization(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		response := map[string]interface{}{
			"status":      "200 OK",
			"status_code": 200,
			"payload":     "Request Dummy Payload",
		}
		json.NewEncoder(w).Encode(response)
	}))

	defer server.Close()

	c, err := SetRequest().
		WithUrl(server.URL).
		SetAuthorization("Bearer some-dummy-token").
		Get().
		Response()

	assert.Nil(t, err)
	assert.NotNil(t, c)
	assert.Equal(t, "Bearer some-dummy-token", c.Request.Header.Get("Authorization"))
}

func TestHttpCall_SetHeaderAndBulkHeaders(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		response := map[string]interface{}{
			"status":      "success",
			"status_code": 200,
			"payload":     "Request Dummy Payload",
		}
		json.NewEncoder(w).Encode(response)
	}))

	defer server.Close()

	c := SetRequest().
		WithUrl(server.URL).
		SetHeader("Header", "Value").
		SetHeaderBulk(map[string]string{
			"Header1": "Value1",
			"Header2": "Value2",
		})

	assert.NotNil(t, c)
	_, _ = c.Get().Payload()
	assert.Equal(t, "Value", c.requestItems.header.payload["Header"])
	assert.Equal(t, "Value1", c.requestItems.header.payload["Header1"])
	assert.Equal(t, "Value2", c.requestItems.header.payload["Header2"])
}

func TestHttpCall_SetDisableJsonHeaders(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		response := map[string]interface{}{
			"status":      "success",
			"status_code": 200,
			"payload":     "Request Dummy Payload",
		}
		json.NewEncoder(w).Encode(response)
	}))

	defer server.Close()

	c := SetRequest().
		WithUrl(server.URL).
		DisableJsonHeaders()

	assert.NotNil(t, c)

	response, err := c.Get().Response()
	assert.Nil(t, err)
	assert.Empty(t, response.Request.Header.Get("Accept"))
	assert.Empty(t, response.Request.Header.Get("Content-Type"))
}

func TestHttpCall_SetQueryParamAndBulkQueryParams(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Return a JSON response
		response := map[string]interface{}{
			"status":      "success",
			"status_code": 200,
			"payload":     "Request Dummy Payload",
		}
		json.NewEncoder(w).Encode(response)
	}))

	defer server.Close()

	c := SetRequest().
		WithUrl(server.URL).
		SetQueryParam("sort", "desc").
		SetQueryParamBulk(map[string]string{
			"page":  "1",
			"limit": "10",
		})

	assert.NotNil(t, c)
	response, err := c.Get().Response()
	assert.Nil(t, err)
	assert.Equal(t, "sort=desc&page=1&limit=10", response.Request.URL.RawQuery)
}

func TestHttpCall_SetBody(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var body map[string]interface{}
		err := json.NewDecoder(r.Body).Decode(&body)

		if err != nil {
			t.Errorf("Failed to read request body: %v", err)
			return
		}

		assert.Equal(t, body, map[string]interface{}{"key": "value"})

		defer r.Body.Close()

		response := map[string]interface{}{
			"status":      "200 OK",
			"status_code": 200,
			"payload":     "Request Dummy Payload",
		}
		err = json.NewEncoder(w).Encode(response)
		if err != nil {
			t.Errorf("Failed to encode response: %v", err)
		}
	}))
	defer server.Close()

	c := SetRequest().
		WithUrl(server.URL).
		SetBody(map[string]interface{}{
			"key": "value",
		})

	response, err := c.Post().Response()

	assert.Nil(t, err)
	assert.Equal(t, response.StatusCode, 200)
}

func TestHttpCall_SetFormAndFormBulk(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		/*r.ParseForm()
		  formValue1 := r.Form.Get("form-key")
		  assert.Equal(t, formValue1, "value")

		  formValue1 := r.Form.Get("form-key1")
		  assert.Equal(t, formValue1, "value1")

		  formValue2 := r.Form.Get("form-key2")
		  assert.Equal(t, formValue2, "value2")*/

		defer r.Body.Close()

		response := map[string]interface{}{
			"status":      "200 OK",
			"status_code": 200,
			"payload":     nil,
		}

		err := json.NewEncoder(w).Encode(response)

		if err != nil {
			t.Errorf("Failed to encode response: %v", err)
		}
	}))

	defer server.Close()

	c := SetRequest().
		WithUrl(server.URL).
		SetForm("form-key", "value").
		SetFormBulk(map[string]string{
			"form-key1": "value1",
			"form-key2": "value2",
		})

	response, err := c.Post().Response()

	assert.Nil(t, err)
	assert.Equal(t, response.Request.PostFormValue("form-key"), "value")
	assert.Equal(t, response.Request.PostFormValue("form-key1"), "value1")
	assert.Equal(t, response.Request.PostFormValue("form-key2"), "value2")
}

func TestHttpCall_GetRequestMethod(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, r.Method, http.MethodGet)

		response := map[string]interface{}{
			"status":      "200 OK",
			"status_code": 200,
			"payload":     nil,
		}

		err := json.NewEncoder(w).Encode(response)

		if err != nil {
			t.Errorf("Failed to encode response: %v", err)
		}
	}))

	defer server.Close()

	c := SetRequest().WithUrl(server.URL)
	response, err := c.Get().Response()

	assert.Nil(t, err)
	assert.Equal(t, response.StatusCode, 200)
}

func TestHttpCall_PostRequestMethod(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, r.Method, http.MethodPost)

		response := map[string]interface{}{
			"status":      "200 OK",
			"status_code": 200,
			"payload":     nil,
		}

		err := json.NewEncoder(w).Encode(response)

		if err != nil {
			t.Errorf("Failed to encode response: %v", err)
		}
	}))

	defer server.Close()

	c := SetRequest().WithUrl(server.URL)
	response, err := c.Post().Response()

	assert.Nil(t, err)
	assert.Equal(t, response.StatusCode, 200)
}

func TestHttpCall_PutRequestMethod(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, r.Method, http.MethodPut)

		response := map[string]interface{}{
			"status":      "200 OK",
			"status_code": 200,
			"payload":     nil,
		}

		err := json.NewEncoder(w).Encode(response)

		if err != nil {
			t.Errorf("Failed to encode response: %v", err)
		}
	}))

	defer server.Close()

	c := SetRequest().WithUrl(server.URL)
	response, err := c.Put().Response()

	assert.Nil(t, err)
	assert.Equal(t, response.StatusCode, 200)
}

func TestHttpCall_DeleteRequestMethod(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, r.Method, http.MethodDelete)

		response := map[string]interface{}{
			"status":      "200 OK",
			"status_code": 200,
			"payload":     nil,
		}

		err := json.NewEncoder(w).Encode(response)

		if err != nil {
			t.Errorf("Failed to encode response: %v", err)
		}
	}))

	defer server.Close()

	c := SetRequest().WithUrl(server.URL)
	response, err := c.Delete().Response()

	assert.Nil(t, err)
	assert.Equal(t, response.StatusCode, 200)
}
