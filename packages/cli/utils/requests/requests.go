package requests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

const (
	CONTENT_TYPE_HEADER           = "Content-Type"
	CONTENT_TYPE_APPLICATION_JSON = "application/json"
	RESPONSE_ERROR                = "Response Error"
	RESPONSE_STATUS               = "Response Status"
	RESPONSE_BODY                 = "Response Body"
)

var client = &http.Client{}

// Options lets you pass headers, query params and JSON body.
type Options struct {
	Header http.Header
	Query  map[string]string
	Body   interface{} // accept any struct/map
}

// Core reusable request helper
func doRequest(method, rawURL string, options Options) ([]byte, error) {

	// --- Build URL with query params ---
	u, urlErr := url.Parse(rawURL)
	if urlErr != nil {
		return nil, urlErr
	}

	if options.Query != nil {
		q := u.Query()
		for k, v := range options.Query {
			q.Set(k, v)
		}
		u.RawQuery = q.Encode()
	}

	// --- Prepare body ---
	var bodyReader io.Reader = nil
	if options.Body != nil {
		jsonBytes, err := json.Marshal(options.Body)
		if err != nil {
			return nil, err
		}
		bodyReader = bytes.NewBuffer(jsonBytes)
	}

	// --- Create request ---
	req, reqErr := http.NewRequest(method, u.String(), bodyReader)
	if reqErr != nil {
		return nil, reqErr
	}

	// --- Set headers ---
	for key, values := range options.Header {
		for _, v := range values {
			req.Header.Add(key, v)
		}
	}

	if options.Body != nil {
		req.Header.Set(CONTENT_TYPE_HEADER, CONTENT_TYPE_APPLICATION_JSON)
	}

	// --- Execute ---
	resp, respErr := client.Do(req)
	if respErr != nil {
		return nil, respErr
	}
	defer resp.Body.Close()

	// --- Read output ---
	respBody, readErr := io.ReadAll(resp.Body)
	if readErr != nil {
		return nil, readErr
	}

	// Optional logging
	fmt.Println(RESPONSE_STATUS, ":", resp.Status)
	fmt.Println(RESPONSE_BODY, ":", string(respBody))

	// --- Return response ---
	return respBody, nil
}

func Get(url string, options Options) ([]byte, error) { return doRequest(http.MethodGet, url, options) }
func Post(url string, options Options) ([]byte, error) {
	return doRequest(http.MethodPost, url, options)
}
func Put(url string, options Options) ([]byte, error) { return doRequest(http.MethodPut, url, options) }
func Patch(url string, options Options) ([]byte, error) {
	return doRequest(http.MethodPatch, url, options)
}
func Delete(url string, options Options) ([]byte, error) {
	return doRequest(http.MethodDelete, url, options)
}
