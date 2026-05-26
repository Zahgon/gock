package gock

import "net/http"

// MatchersHeader exposes an slice of HTTP header specific mock matchers.
var MatchersHeader = []MatchFunc{
	MatchMethod,
	MatchScheme,
	MatchHost,
	MatchPath,
	MatchHeaders,
	MatchQueryParams,
	MatchPathParams,
}

// MatchersBody exposes an slice of HTTP body specific built-in mock matchers.
var MatchersBody = []MatchFunc{
	MatchBody,
}

// Matchers stores all the built-in mock matchers.
var Matchers = append(MatchersHeader, MatchersBody...)

// DefaultMatcher stores the default Matcher instance used to match mocks.
var DefaultMatcher = NewMatcher()

// MatchFunc represents the required function
// interface implemented by matchers.
type MatchFunc func(*http.Request, *Request) (bool, error)

// Matcher represents the required interface implemented by mock matchers.
type Matcher interface {
	// Get returns a slice of registered function matchers.
	Get() []MatchFunc

	// Add adds a new matcher function.
	Add(MatchFunc)

	// Set sets the matchers functions stack.
	Set([]MatchFunc)

	// Flush flushes the current matchers function stack.
	Flush()

	// Match matches the given http.Request with a mock Request.
	Match(*http.Request, *Request) (bool, error)
}

// MockMatcher implements a mock matcher
type MockMatcher struct {
	Matchers []MatchFunc
}

// NewMatcher creates a new mock matcher
// using the default matcher functions.
func NewMatcher() *MockMatcher { _ = "STUB: not implemented"; return nil }

// NewBasicMatcher creates a new matcher with header only mock matchers.
func NewBasicMatcher() *MockMatcher { _ = "STUB: not implemented"; return nil }

// NewEmptyMatcher creates a new empty matcher without default matchers.
func NewEmptyMatcher() *MockMatcher { _ = "STUB: not implemented"; return nil }

// Get returns a slice of registered function matchers.
func (m *MockMatcher) Get() []MatchFunc { _ = "STUB: not implemented"; return nil }

// Add adds a new function matcher.
func (m *MockMatcher) Add(fn MatchFunc) { _ = "STUB: not implemented"; return }

// Set sets a new stack of matchers functions.
func (m *MockMatcher) Set(stack []MatchFunc) {
	_ = "STUB: not implemented"

	// Flush flushes the current matcher
	return
}

func (m *MockMatcher) Flush() { _ = "STUB: not implemented"; return }

// Clone returns a separate MockMatcher instance that has a copy of the same MatcherFuncs
func (m *MockMatcher) Clone() *MockMatcher { _ = "STUB: not implemented"; return nil }

// Match matches the given http.Request with a mock request
// returning true in case that the request matches, otherwise false.
func (m *MockMatcher) Match(req *http.Request, ereq *Request) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// MatchMock is a helper function that matches the given http.Request
// in the list of registered mocks, returning it if matches or error if it fails.
func MatchMock(req *http.Request) (Mock, error) { _ = "STUB: not implemented"; return *new(Mock), nil }
