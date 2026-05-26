package gock

import (
	"net/http"
	"sync"
)

// Mock represents the required interface that must
// be implemented by HTTP mock instances.
type Mock interface {
	// Disable disables the current mock manually.
	Disable()

	// Done returns true if the current mock is disabled.
	Done() bool

	// Request returns the mock Request instance.
	Request() *Request

	// Response returns the mock Response instance.
	Response() *Response

	// Match matches the given http.Request with the current mock.
	Match(*http.Request) (bool, error)

	// AddMatcher adds a new matcher function.
	AddMatcher(MatchFunc)

	// SetMatcher uses a new matcher implementation.
	SetMatcher(Matcher)
}

// Mocker implements a Mock capable interface providing
// a default mock configuration used internally to store mocks.
type Mocker struct {
	// disabler stores a disabler for thread safety checking current mock is disabled
	disabler *disabler

	// mutex stores the mock mutex for thread safety.
	mutex sync.Mutex

	// matcher stores a Matcher capable instance to match the given http.Request.
	matcher Matcher

	// request stores the mock Request to match.
	request *Request

	// response stores the mock Response to use in case of match.
	response *Response
}

type disabler struct {
	// disabled stores if the current mock is disabled.
	disabled bool

	// mutex stores the disabler mutex for thread safety.
	mutex sync.RWMutex
}

func (d *disabler) isDisabled() bool { _ = "STUB: not implemented"; return false }

func (d *disabler) Disable() { _ = "STUB: not implemented"; return }

// NewMock creates a new HTTP mock based on the given request and response instances.
// It's mostly used internally.
func NewMock(req *Request, res *Response) *Mocker { _ = "STUB: not implemented"; return nil }

// Disable disables the current mock manually.
func (m *Mocker) Disable() { _ = "STUB: not implemented"; return }

// Done returns true in case that the current mock
// instance is disabled and therefore must be removed.
func (m *Mocker) Done() bool {
	_ = "STUB: not implemented"
	// prevent deadlock with m.mutex
	return false
}

// Request returns the Request instance
// configured for the current HTTP mock.
func (m *Mocker) Request() *Request {
	_ = "STUB: not implemented"

	// Response returns the Response instance
	// configured for the current HTTP mock.
	return nil
}

func (m *Mocker) Response() *Response {
	_ = "STUB: not implemented"

	// Match matches the given http.Request with the current Request
	// mock expectation, returning true if matches.
	return nil
}

func (m *Mocker) Match(req *http.Request) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// Filter

// Map

// Match

// SetMatcher sets a new matcher implementation
// for the current mock expectation.
func (m *Mocker) SetMatcher(matcher Matcher) { _ = "STUB: not implemented"; return }

// AddMatcher adds a new matcher function
// for the current mock expectation.
func (m *Mocker) AddMatcher(fn MatchFunc) {
	_ = "STUB: not implemented"

	// decrement decrements the current mock Request counter.
	return
}

func (m *Mocker) decrement() { _ = "STUB: not implemented"; return }
