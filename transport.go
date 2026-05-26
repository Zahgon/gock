package gock

import (
	"errors"
	"net/http"
	"sync"
)

// var mutex *sync.Mutex = &sync.Mutex{}

var (
	// DefaultTransport stores the default mock transport used by gock.
	DefaultTransport = NewTransport()

	// NativeTransport stores the native net/http default transport
	// in order to restore it when needed.
	NativeTransport = http.DefaultTransport
)

var (
	// ErrCannotMatch store the error returned in case of no matches.
	ErrCannotMatch = errors.New("gock: cannot match any request")
)

// Transport implements http.RoundTripper, which fulfills single http requests issued by
// an http.Client.
//
// gock's Transport encapsulates a given or default http.Transport for further
// delegation, if needed.
type Transport struct {
	// mutex is used to make transport thread-safe of concurrent uses across goroutines.
	mutex sync.Mutex

	// Transport encapsulates the original http.RoundTripper transport interface for delegation.
	Transport http.RoundTripper
}

// NewTransport creates a new *Transport with no responders.
func NewTransport() *Transport { _ = "STUB: not implemented"; return nil }

// RoundTrip receives HTTP requests and routes them to the appropriate responder.  It is required to
// implement the http.RoundTripper interface.  You will not interact with this directly, instead
// the *http.Client you are using will call it for you.
func (m *Transport) RoundTrip(req *http.Request) (*http.Response, error) {
	_ = "STUB: not implemented"
	// Just act as a proxy if not intercepting
	return nil, nil
}

// Match mock for the incoming http.Request

// Invoke the observer with the intercepted http.Request and matched mock

// Verify if should use real networking

// Ensure me unlock the mutex before building the response

// Perform real networking via original transport

// In no mock matched, continue with the response

// CancelRequest is a no-op function.
func (m *Transport) CancelRequest(req *http.Request) { _ = "STUB: not implemented"; return }

func shouldUseNetwork(req *http.Request, mock Mock) bool { _ = "STUB: not implemented"; return false }
