// This example shows how to enable the networking for a request to a local server
// and mock a second request to a remote server.
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"

	"github.com/h2non/gock"
)

// Starts a local HTTP server in background
func startHTTPServer() *httptest.Server { _ = "STUB: not implemented"; return nil }

// MUST NOT get original body since the networking
// wasn't enabled for this request

func main() {
	defer gock.Disable()
	defer gock.DisableNetworking()

	srv := startHTTPServer()
	defer srv.Close()

	// Register our local server
	gock.New(srv.URL).
		EnableNetworking()

	gock.New("http://httpbin.org").
		Get("/nope").
		Reply(201).
		SetHeader("Server", "gock")

	res, err := http.Get(srv.URL)
	if err != nil {
		fmt.Printf("Error from request to localhost: %s", err)
		return
	}

	// The response status comes from the mock
	fmt.Printf("Status: %d\n", res.StatusCode)
	// The server header comes from mock as well
	fmt.Printf("Server header: %s\n", res.Header.Get("Server"))
	// MUST get original response since the networking was enabled for this request
	body, _ := ioutil.ReadAll(res.Body)
	fmt.Printf("Body From Local Server: %s", string(body))
}
