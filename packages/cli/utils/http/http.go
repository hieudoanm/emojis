package http

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const CONTENT_TYPE_HEADER = "Content-Type"
const CONTENT_TYPE_APPLICATION_JSON = "application/json"
const RESPONSE_ERROR = "Response Error"
const RESPONSE_STATUS = "Response Status"
const RESPONSE_BODY = "Response Body"

// Options
type Options struct {
	Header http.Header
	Query  map[string]string
	Body   map[string]interface{}
}

// Get ...
func Get(url string, options Options) ([]byte, error) {
	// Make the GET request
	request, requestError := http.NewRequest("GET", url, nil)
	if requestError != nil {
		fmt.Println("Error:", requestError)
		return nil, requestError
	}
	// Set Headers
	request.Header = options.Header
	// Response
	client := http.Client{}
	response, responseError := client.Do(request)
	if responseError != nil {
		return nil, responseError
	}
	defer response.Body.Close() // Ensure response body is closed
	// Read response body
	body, bodyError := io.ReadAll(response.Body)
	if bodyError != nil {
		fmt.Println(RESPONSE_ERROR, ":", bodyError)
		return nil, bodyError
	}
	fmt.Println(RESPONSE_STATUS, ":", response.Status)
	fmt.Println(RESPONSE_BODY, ":", string(body))

	// Return response body
	return body, nil
}

// Post ...
func Post(url string, options Options) ([]byte, error) {
	// JSON payload
	requestData, jsonMarshalError := json.Marshal(options.Body)
	if jsonMarshalError != nil {
		return nil, jsonMarshalError
	}

	// Create request
	request, requestError := http.NewRequest("POST", url, bytes.NewBuffer(requestData))
	if requestError != nil {
		return nil, requestError
	}
	// Set Headers
	request.Header = options.Header
	request.Header.Set(CONTENT_TYPE_HEADER, CONTENT_TYPE_APPLICATION_JSON)
	// Send request
	client := &http.Client{}
	response, responseError := client.Do(request)
	if responseError != nil {
		fmt.Println(RESPONSE_ERROR, ":", responseError)
		return nil, responseError
	}
	defer response.Body.Close()
	// Read response body
	body, bodyError := io.ReadAll(response.Body)
	if bodyError != nil {
		fmt.Println(RESPONSE_ERROR, ":", bodyError)
		return nil, bodyError
	}
	fmt.Println(RESPONSE_STATUS, ":", response.Status)
	fmt.Println(RESPONSE_BODY, ":", string(body))
	// Return response body
	return body, nil
}

// Put ...
func Put(url string, options Options) ([]byte, error) {
	// JSON payload
	requestData, jsonMarshalError := json.Marshal(options.Body)
	if jsonMarshalError != nil {
		return nil, jsonMarshalError
	}

	// Create request
	request, requestError := http.NewRequest("PUT", url, bytes.NewBuffer(requestData))
	if requestError != nil {
		return nil, requestError
	}
	// Set Headers
	request.Header = options.Header
	request.Header.Set(CONTENT_TYPE_HEADER, CONTENT_TYPE_APPLICATION_JSON)
	// Send request
	client := &http.Client{}
	response, responseError := client.Do(request)
	if responseError != nil {
		fmt.Println("Request error:", responseError)
		return nil, responseError
	}
	defer response.Body.Close()
	// Read response body
	body, bodyError := io.ReadAll(response.Body)
	if bodyError != nil {
		fmt.Println(RESPONSE_ERROR, ":", bodyError)
		return nil, bodyError
	}
	fmt.Println(RESPONSE_STATUS, ":", response.Status)
	fmt.Println(RESPONSE_BODY, ":", string(body))
	// Return response body
	return body, nil
}

// Delete ...
func Delete(url string, options Options) ([]byte, error) {
	// JSON payload
	requestData, jsonMarshalError := json.Marshal(options.Body)
	if jsonMarshalError != nil {
		return nil, jsonMarshalError
	}

	// Create request
	request, requestError := http.NewRequest("DELETE", url, bytes.NewBuffer(requestData))
	if requestError != nil {
		return nil, requestError
	}
	// Set Headers
	request.Header = options.Header
	request.Header.Set(CONTENT_TYPE_HEADER, CONTENT_TYPE_APPLICATION_JSON)
	// Send request
	client := &http.Client{}
	response, responseError := client.Do(request)
	if responseError != nil {
		fmt.Println("Request error:", responseError)
		return nil, responseError
	}
	defer response.Body.Close()
	// Read response body
	body, bodyError := io.ReadAll(response.Body)
	if bodyError != nil {
		fmt.Println(RESPONSE_ERROR, ":", bodyError)
		return nil, bodyError
	}
	fmt.Println(RESPONSE_STATUS, ":", response.Status)
	fmt.Println(RESPONSE_BODY, ":", string(body))
	// Return response body
	return body, nil
}
