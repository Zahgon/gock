package gock

import (
	"io"
	"net/http"
)

// Responder builds a mock http.Response based on the given Response mock.
func Responder(req *http.Request, mock *Response, res *http.Response) (*http.Response, error) {
	_ = "STUB: not implemented"
	// If error present, reply it
	return nil, nil
}

// Apply response filter

// Define mock status code

// Define headers by merging fields

// Define mock body, if present

// Set raw mock body, if exist

// Apply response mappers

// Sleep to simulate delay, if necessary

// allow escaping from sleep due to request context expiration or cancellation

// cleanly stop the timer

// check if the request context has ended. we could put this up in the delay code above, but putting it here
// has the added benefit of working even when there is no delay (very small timeouts, already-done contexts, etc.)

// cleanly close the response and return the context error

// createResponse creates a new http.Response with default fields.
func createResponse(req *http.Request) *http.Response { _ = "STUB: not implemented"; return nil }

// mergeHeaders copies the mock headers.
func mergeHeaders(res *http.Response, mres *Response) http.Header {
	_ = "STUB: not implemented"
	return *new(http.Header)
}

// createReadCloser creates an io.ReadCloser from a byte slice that is suitable for use as an
// http response body.
func createReadCloser(body []byte) io.ReadCloser {
	_ = "STUB: not implemented"
	return *new(io.ReadCloser)
}
