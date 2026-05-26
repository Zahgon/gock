package gock

import (
	"io"
	"net/http"
)

// EOL represents the end of line character.
const EOL = 0xa

// BodyTypes stores the supported MIME body types for matching.
// Currently only text-based types.
var BodyTypes = []string{
	"text/html",
	"text/plain",
	"application/json",
	"application/xml",
	"multipart/form-data",
	"application/x-www-form-urlencoded",
}

// BodyTypeAliases stores a generic MIME type by alias.
var BodyTypeAliases = map[string]string{
	"html": "text/html",
	"text": "text/plain",
	"json": "application/json",
	"xml":  "application/xml",
	"form": "multipart/form-data",
	"url":  "application/x-www-form-urlencoded",
}

// CompressionSchemes stores the supported Content-Encoding types for decompression.
var CompressionSchemes = []string{
	"gzip",
}

// MatchMethod matches the HTTP method of the given request.
func MatchMethod(req *http.Request, ereq *Request) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// MatchScheme matches the request URL protocol scheme.
func MatchScheme(req *http.Request, ereq *Request) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// MatchHost matches the HTTP host header field of the given request.
func MatchHost(req *http.Request, ereq *Request) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// MatchPath matches the HTTP URL path of the given request.
func MatchPath(req *http.Request, ereq *Request) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// MatchHeaders matches the headers fields of the given request.
func MatchHeaders(req *http.Request, ereq *Request) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// Some values may contain reserved regex params e.g. "()", try matching with these escaped.

// MatchQueryParams matches the URL query params fields of the given request.
func MatchQueryParams(req *http.Request, ereq *Request) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// MatchPathParams matches the URL path parameters of the given request.
func MatchPathParams(req *http.Request, ereq *Request) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// MatchBody tries to match the request body.
// TODO: not too smart now, needs several improvements.
func MatchBody(req *http.Request, ereq *Request) (bool, error) {
	_ = "STUB: not implemented"
	// If match body is empty, just continue
	return false, nil
}

// Only can match certain MIME body types

// Can only match certain compression schemes

// Create a reader for the body depending on compression type

// Read the whole request body

// Restore body reader stream

// If empty, ignore the match

// Match body by atomic string comparison

// Match request body by regexp

// todo - add conditional do only perform the conversion of body bytes
// representation of JSON to a map and then compare them for equality.

// Check if the key + value pairs match

// Ensure that both byte bodies that that should be JSON can be converted to maps.

func supportedType(req *http.Request, ereq *Request) bool { _ = "STUB: not implemented"; return false }

func supportedCompressionScheme(req *http.Request) bool { _ = "STUB: not implemented"; return false }

func castToString(buf []byte) string { _ = "STUB: not implemented"; return "" }

func compressionReader(r io.ReadCloser, scheme string) (io.ReadCloser, error) {
	_ = "STUB: not implemented"
	return *new(io.ReadCloser), nil
}
