package gock

import (
	"sync"
)

// storeMutex is used interally for store synchronization.
var storeMutex = sync.RWMutex{}

// mocks is internally used to store registered mocks.
var mocks = []Mock{}

// Register registers a new mock in the current mocks stack.
func Register(mock Mock) { _ = "STUB: not implemented"; return }

// Make ops thread safe

// Expose mock in request/response for delegation

// Registers the mock in the global store

// GetAll returns the current stack of registered mocks.
func GetAll() []Mock { _ = "STUB: not implemented"; return nil }

// Exists checks if the given Mock is already registered.
func Exists(m Mock) bool { _ = "STUB: not implemented"; return false }

// Remove removes a registered mock by reference.
func Remove(m Mock) { _ = "STUB: not implemented"; return }

// Flush flushes the current stack of registered mocks.
func Flush() { _ = "STUB: not implemented"; return }

// Pending returns an slice of pending mocks.
func Pending() []Mock { _ = "STUB: not implemented"; return nil }

// IsDone returns true if all the registered mocks has been triggered successfully.
func IsDone() bool { _ = "STUB: not implemented"; return false }

// IsPending returns true if there are pending mocks.
func IsPending() bool { _ = "STUB: not implemented"; return false }

// Clean cleans the mocks store removing disabled or obsolete mocks.
func Clean() { _ = "STUB: not implemented"; return }
