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

const (
	CONTENT_TYPE_HEADER           = "Content-Type"
	CONTENT_TYPE_APPLICATION_JSON = "application/json"
	RESPONSE_ERROR                = "Response Error"
	RESPONSE_STATUS               = "Response Status"
	RESPONSE_BODY                 = "Response Body"
)

// Global HTTP client with timeout
var client = &http.Client{
	Timeout: 15 * time.Second,
}

type Options struct {
	Header  http.Header
	Query   map[string]string
	Body    interface{}
	Timeout time.Duration // optional per-request timeout
	Retries int           // retry count
}

// Helper: check if error is retryable (network issues)
func isRetryableError(err error) bool {
	var netErr net.Error
	return errors.As(err, &netErr)
}

// Helper: is 5xx → retryable server error
func isRetryableStatus(code int) bool {
	return code >= 500 && code <= 599
}

func doRequest(method, rawURL string, options Options) ([]byte, error) {
	maxRetries := options.Retries
	if maxRetries < 0 {
		maxRetries = 0
	}

	var lastErr error

	for attempt := 0; attempt <= maxRetries; attempt++ {

		// --- TIMEOUT ---
		timeout := options.Timeout
		if timeout == 0 {
			timeout = 10 * time.Second
		}
		ctx, cancel := context.WithTimeout(context.Background(), timeout)
		defer cancel()

		// --- Build URL ---
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

		// --- Body ---
		var bodyReader io.Reader = nil
		if options.Body != nil {
			jsonBytes, err := json.Marshal(options.Body)
			if err != nil {
				return nil, err
			}
			bodyReader = bytes.NewBuffer(jsonBytes)
		}

		// --- Create request ---
		req, reqErr := http.NewRequestWithContext(ctx, method, u.String(), bodyReader)
		if reqErr != nil {
			return nil, reqErr
		}

		// Headers
		for k, values := range options.Header {
			for _, v := range values {
				req.Header.Add(k, v)
			}
		}
		if options.Body != nil {
			req.Header.Set(CONTENT_TYPE_HEADER, CONTENT_TYPE_APPLICATION_JSON)
		}

		// --- Execute ---
		resp, err := client.Do(req)
		if err != nil {
			lastErr = err

			// retry network errors
			if attempt < maxRetries && isRetryableError(err) {
				time.Sleep(time.Duration(attempt+1) * 300 * time.Millisecond)
				continue
			}
			return nil, err
		}

		// Ensure close
		respBody, readErr := io.ReadAll(resp.Body)
		resp.Body.Close()
		if readErr != nil {
			return nil, readErr
		}

		// Retry on 5xx
		if isRetryableStatus(resp.StatusCode) && attempt < maxRetries {
			lastErr = fmt.Errorf("server error: %v", resp.Status)
			time.Sleep(time.Duration(attempt+1) * 300 * time.Millisecond)
			continue
		}

		fmt.Println(RESPONSE_STATUS, ":", resp.Status)
		fmt.Println(RESPONSE_BODY, ":", string(respBody))
		return respBody, nil
	}

	return nil, fmt.Errorf("request failed after retries: %w", lastErr)
}

// Methods
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
