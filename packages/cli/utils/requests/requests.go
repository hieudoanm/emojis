package requests

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net"
	"net/http"
	"net/url"
	"time"
)

// ------------------------
// Header Keys
// ------------------------
const (
	HEADER_CONTENT_TYPE     = "Content-Type"
	HEADER_ACCEPT           = "Accept"
	HEADER_AUTHORIZATION    = "Authorization"
	HEADER_USER_AGENT       = "User-Agent"
	HEADER_CACHE_CONTROL    = "Cache-Control"
	HEADER_CONTENT_LENGTH   = "Content-Length"
	HEADER_CONTENT_ENCODING = "Content-Encoding"
	HEADER_ACCEPT_ENCODING  = "Accept-Encoding"
	HEADER_ACCEPT_LANGUAGE  = "Accept-Language"
)

// ------------------------
// MIME / Content-Type Values
// ------------------------
const (
	CONTENT_TYPE_JSON            = "application/json"
	CONTENT_TYPE_XML             = "application/xml"
	CONTENT_TYPE_FORM_URLENCODED = "application/x-www-form-urlencoded"
	CONTENT_TYPE_TEXT_PLAIN      = "text/plain"
	CONTENT_TYPE_TEXT_HTML       = "text/html"
	CONTENT_TYPE_MULTIPART_FORM  = "multipart/form-data"
)

// ------------------------
// Log Labels
// ------------------------
const (
	LOG_RESPONSE_STATUS = "Response Status"
	LOG_RESPONSE_BODY   = "Response Body"
)

var client = &http.Client{Timeout: 15 * time.Second}

type Options struct {
	Header  http.Header
	Query   map[string]string
	Body    interface{}
	Timeout time.Duration
	Retries int
}

// ------------------------
// Small helper functions
// ------------------------

func buildURL(rawURL string, query map[string]string) (*url.URL, error) {
	u, err := url.Parse(rawURL)
	if err != nil {
		return nil, err
	}
	if query != nil {
		q := u.Query()
		for k, v := range query {
			q.Set(k, v)
		}
		u.RawQuery = q.Encode()
	}
	return u, nil
}

func buildBody(body interface{}) (io.Reader, error) {
	if body == nil {
		return nil, nil
	}
	jsonBytes, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	return bytes.NewBuffer(jsonBytes), nil
}

func createRequest(ctx context.Context, method string, u *url.URL, body io.Reader, opt Options) (*http.Request, error) {
	req, err := http.NewRequestWithContext(ctx, method, u.String(), body)
	if err != nil {
		return nil, err
	}
	for k, values := range opt.Header {
		for _, v := range values {
			req.Header.Add(k, v)
		}
	}
	if opt.Body != nil {
		req.Header.Set(HEADER_CONTENT_TYPE, CONTENT_TYPE_JSON)
	}
	return req, nil
}

func shouldRetry(err error, status int, attempt, maxRetries int) bool {
	if attempt >= maxRetries {
		return false
	}
	if err != nil {
		var netErr net.Error
		return errors.As(err, &netErr)
	}
	return status >= 500 && status <= 599
}

func backoff(attempt int) {
	time.Sleep(time.Duration(attempt+1) * 300 * time.Millisecond)
}

// ------------------------
// Main request function
// ------------------------
func handleResponse(resp *http.Response) ([]byte, int, error) {
	if resp == nil {
		return nil, 0, errors.New("nil response")
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, resp.StatusCode, err
	}

	return body, resp.StatusCode, nil
}

func attemptRequest(
	method string,
	rawURL string,
	options Options,
	ctx context.Context,
) (*http.Response, error) {

	urlObj, err := buildURL(rawURL, options.Query)
	if err != nil {
		return nil, err
	}

	body, err := buildBody(options.Body)
	if err != nil {
		return nil, err
	}

	req, err := createRequest(ctx, method, urlObj, body, options)
	if err != nil {
		return nil, err
	}

	return client.Do(req)
}

func doRequest(method, rawURL string, options Options) ([]byte, error) {
	maxRetries := options.Retries
	timeout := options.Timeout
	if timeout == 0 {
		timeout = 10 * time.Second
	}

	var lastErr error

	for attempt := 0; attempt <= maxRetries; attempt++ {

		ctx, cancel := context.WithTimeout(context.Background(), timeout)

		resp, err := attemptRequest(method, rawURL, options, ctx)
		cancel()

		if err != nil {
			lastErr = err
			if shouldRetry(err, 0, attempt, maxRetries) {
				backoff(attempt)
				continue
			}
			return nil, err
		}

		body, status, readErr := handleResponse(resp)
		if readErr != nil {
			return nil, readErr
		}

		if shouldRetry(nil, status, attempt, maxRetries) {
			lastErr = fmt.Errorf("server error: %v", resp.Status)
			backoff(attempt)
			continue
		}

		fmt.Println(LOG_RESPONSE_STATUS, ":", resp.Status)
		fmt.Println(LOG_RESPONSE_BODY, ":", string(body))

		return body, nil
	}

	return nil, fmt.Errorf("request failed after retries: %w", lastErr)
}

// Public methods
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
