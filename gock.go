package gock

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"sync"
)

// mutex is used interally for locking thread-sensitive functions.
var mutex = &sync.Mutex{}

// config global singleton store.
var config = struct {
	Networking        bool
	NetworkingFilters []FilterRequestFunc
	Observer          ObserverFunc
}{}

// ObserverFunc is implemented by users to inspect the outgoing intercepted HTTP traffic
type ObserverFunc func(*http.Request, Mock)

// DumpRequest is a default implementation of ObserverFunc that dumps
// the HTTP/1.x wire representation of the http request
var DumpRequest ObserverFunc = func(request *http.Request, mock Mock) {
	bytes, _ := httputil.DumpRequestOut(request, true)
	fmt.Println(string(bytes))
	fmt.Printf("\nMatches: %v\n---\n", mock != nil)
}

// track unmatched requests so they can be tested for
var unmatchedRequests = []*http.Request{}

// New creates and registers a new HTTP mock with
// default settings and returns the Request DSL for HTTP mock
// definition and set up.
func New(uri string) *Request { _ = "STUB: not implemented"; return nil }

// Create the new mock expectation

// Intercepting returns true if gock is currently able to intercept.
func Intercepting() bool { _ = "STUB: not implemented"; return false }

// Intercept enables HTTP traffic interception via http.DefaultTransport.
// If you are using a custom HTTP transport, you have to use `gock.Transport()`
func Intercept() { _ = "STUB: not implemented"; return }

// InterceptClient allows the developer to intercept HTTP traffic using
// a custom http.Client who uses a non default http.Transport/http.RoundTripper implementation.
func InterceptClient(cli *http.Client) { _ = "STUB: not implemented"; return }

// if transport already intercepted, just ignore it

// RestoreClient allows the developer to disable and restore the
// original transport in the given http.Client.
func RestoreClient(cli *http.Client) { _ = "STUB: not implemented"; return }

// Disable disables HTTP traffic interception by gock.
func Disable() { _ = "STUB: not implemented"; return }

// Off disables the default HTTP interceptors and removes
// all the registered mocks, even if they have not been intercepted yet.
func Off() {
	_ = "STUB: not implemented"

	// OffAll is like `Off()`, but it also removes the unmatched requests registry.
	return
}

func OffAll() { _ = "STUB: not implemented"; return }

// Observe provides a hook to support inspection of the request and matched mock
func Observe(fn ObserverFunc) { _ = "STUB: not implemented"; return }

// EnableNetworking enables real HTTP networking
func EnableNetworking() { _ = "STUB: not implemented"; return }

// DisableNetworking disables real HTTP networking
func DisableNetworking() { _ = "STUB: not implemented"; return }

// NetworkingFilter determines if an http.Request should be triggered or not.
func NetworkingFilter(fn FilterRequestFunc) { _ = "STUB: not implemented"; return }

// DisableNetworkingFilters disables registered networking filters.
func DisableNetworkingFilters() { _ = "STUB: not implemented"; return }

// GetUnmatchedRequests returns all requests that have been received but haven't matched any mock
func GetUnmatchedRequests() []*http.Request { _ = "STUB: not implemented"; return nil }

// HasUnmatchedRequest returns true if gock has received any requests that didn't match a mock
func HasUnmatchedRequest() bool { _ = "STUB: not implemented"; return false }

// CleanUnmatchedRequest cleans the unmatched requests internal registry.
func CleanUnmatchedRequest() { _ = "STUB: not implemented"; return }

func trackUnmatchedRequest(req *http.Request) { _ = "STUB: not implemented"; return }

func normalizeURI(uri string) string { _ = "STUB: not implemented"; return "" }
