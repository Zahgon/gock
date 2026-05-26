package gock

import (
	"io"
	"net/http"
	"net/url"
)

// MapRequestFunc represents the required function interface for request mappers.
type MapRequestFunc func(*http.Request) *http.Request

// FilterRequestFunc represents the required function interface for request filters.
type FilterRequestFunc func(*http.Request) bool

// Request represents the high-level HTTP request used to store
// request fields used to match intercepted requests.
type Request struct {
	// Mock stores the parent mock reference for the current request mock used for method delegation.
	Mock Mock

	// Response stores the current Response instance for the current matches Request.
	Response *Response

	// Error stores the latest mock request configuration error.
	Error error

	// Counter stores the pending times that the current mock should be active.
	Counter int

	// Persisted stores if the current mock should be always active.
	Persisted bool

	// Options stores options for current Request.
	Options Options

	// URLStruct stores the parsed URL as *url.URL struct.
	URLStruct *url.URL

	// Method stores the Request HTTP method to match.
	Method string

	// CompressionScheme stores the Request Compression scheme to match and use for decompression.
	CompressionScheme string

	// Header stores the HTTP header fields to match.
	Header http.Header

	// Cookies stores the Request HTTP cookies values to match.
	Cookies []*http.Cookie

	// PathParams stores the path parameters to match.
	PathParams map[string]string

	// BodyBuffer stores the body data to match.
	BodyBuffer []byte

	// Mappers stores the request functions mappers used for matching.
	Mappers []MapRequestFunc

	// Filters stores the request functions filters used for matching.
	Filters []FilterRequestFunc
}

// NewRequest creates a new Request instance.
func NewRequest() *Request { _ = "STUB: not implemented"; return nil }

// URL defines the mock URL to match.
func (r *Request) URL(uri string) *Request { _ = "STUB: not implemented"; return nil }

// SetURL defines the url.URL struct to be used for matching.
func (r *Request) SetURL(u *url.URL) *Request { _ = "STUB: not implemented"; return nil }

// Path defines the mock URL path value to match.
func (r *Request) Path(path string) *Request { _ = "STUB: not implemented"; return nil }

// Get specifies the GET method and the given URL path to match.
func (r *Request) Get(path string) *Request { _ = "STUB: not implemented"; return nil }

// Post specifies the POST method and the given URL path to match.
func (r *Request) Post(path string) *Request { _ = "STUB: not implemented"; return nil }

// Put specifies the PUT method and the given URL path to match.
func (r *Request) Put(path string) *Request { _ = "STUB: not implemented"; return nil }

// Delete specifies the DELETE method and the given URL path to match.
func (r *Request) Delete(path string) *Request { _ = "STUB: not implemented"; return nil }

// Patch specifies the PATCH method and the given URL path to match.
func (r *Request) Patch(path string) *Request { _ = "STUB: not implemented"; return nil }

// Head specifies the HEAD method and the given URL path to match.
func (r *Request) Head(path string) *Request { _ = "STUB: not implemented"; return nil }

// method is a DRY shortcut used to declare the expected HTTP method and URL path.
func (r *Request) method(method, path string) *Request { _ = "STUB: not implemented"; return nil }

// Body defines the body data to match based on a io.Reader interface.
func (r *Request) Body(body io.Reader) *Request { _ = "STUB: not implemented"; return nil }

// BodyString defines the body to match based on a given string.
func (r *Request) BodyString(body string) *Request { _ = "STUB: not implemented"; return nil }

// File defines the body to match based on the given file path string.
func (r *Request) File(path string) *Request { _ = "STUB: not implemented"; return nil }

// Compression defines the request compression scheme, and enables automatic body decompression.
// Supports only the "gzip" scheme so far.
func (r *Request) Compression(scheme string) *Request { _ = "STUB: not implemented"; return nil }

// JSON defines the JSON body to match based on a given structure.
func (r *Request) JSON(data interface{}) *Request { _ = "STUB: not implemented"; return nil }

// XML defines the XML body to match based on a given structure.
func (r *Request) XML(data interface{}) *Request { _ = "STUB: not implemented"; return nil }

// MatchType defines the request Content-Type MIME header field.
// Supports custom MIME types and type aliases. E.g: json, xml, form, text...
func (r *Request) MatchType(kind string) *Request { _ = "STUB: not implemented"; return nil }

// BasicAuth defines a username and password for HTTP Basic Authentication
func (r *Request) BasicAuth(username, password string) *Request {
	_ = "STUB: not implemented"
	return nil
}

// MatchHeader defines a new key and value header to match.
func (r *Request) MatchHeader(key, value string) *Request { _ = "STUB: not implemented"; return nil }

// HeaderPresent defines that a header field must be present in the request.
func (r *Request) HeaderPresent(key string) *Request { _ = "STUB: not implemented"; return nil }

// MatchHeaders defines a map of key-value headers to match.
func (r *Request) MatchHeaders(headers map[string]string) *Request {
	_ = "STUB: not implemented"
	return nil
}

// MatchParam defines a new key and value URL query param to match.
func (r *Request) MatchParam(key, value string) *Request { _ = "STUB: not implemented"; return nil }

// MatchParams defines a map of URL query param key-value to match.
func (r *Request) MatchParams(params map[string]string) *Request {
	_ = "STUB: not implemented"
	return nil
}

// ParamPresent matches if the given query param key is present in the URL.
func (r *Request) ParamPresent(key string) *Request { _ = "STUB: not implemented"; return nil }

// PathParam matches if a given path parameter key is present in the URL.
//
// The value is representative of the restful resource the key defines, e.g.
//
//	// /users/123/name
//	r.PathParam("users", "123")
//
// would match.
func (r *Request) PathParam(key, val string) *Request { _ = "STUB: not implemented"; return nil }

// Persist defines the current HTTP mock as persistent and won't be removed after intercepting it.
func (r *Request) Persist() *Request { _ = "STUB: not implemented"; return nil }

// WithOptions sets the options for the request.
func (r *Request) WithOptions(options Options) *Request { _ = "STUB: not implemented"; return nil }

// Times defines the number of times that the current HTTP mock should remain active.
func (r *Request) Times(num int) *Request { _ = "STUB: not implemented"; return nil }

// AddMatcher adds a new matcher function to match the request.
func (r *Request) AddMatcher(fn MatchFunc) *Request { _ = "STUB: not implemented"; return nil }

// SetMatcher sets a new matcher function to match the request.
func (r *Request) SetMatcher(matcher Matcher) *Request { _ = "STUB: not implemented"; return nil }

// Map adds a new request mapper function to map http.Request before the matching process.
func (r *Request) Map(fn MapRequestFunc) *Request { _ = "STUB: not implemented"; return nil }

// Filter filters a new request filter function to filter http.Request before the matching process.
func (r *Request) Filter(fn FilterRequestFunc) *Request { _ = "STUB: not implemented"; return nil }

// EnableNetworking enables the use real networking for the current mock.
func (r *Request) EnableNetworking() *Request { _ = "STUB: not implemented"; return nil }

// Reply defines the Response status code and returns the mock Response DSL.
func (r *Request) Reply(status int) *Response { _ = "STUB: not implemented"; return nil }

// ReplyError defines the Response simulated error.
func (r *Request) ReplyError(err error) *Response { _ = "STUB: not implemented"; return nil }

// ReplyFunc allows the developer to define the mock response via a custom function.
func (r *Request) ReplyFunc(replier func(*Response)) *Response {
	_ = "STUB: not implemented"
	return nil
}

// See 2 (end of page 4) https://www.ietf.org/rfc/rfc2617.txt
// "To receive authorization, the client sends the userid and password,
// separated by a single colon (":") character, within a base64
// encoded string in the credentials."
// It is not meant to be urlencoded.
func basicAuth(username, password string) string { _ = "STUB: not implemented"; return "" }
