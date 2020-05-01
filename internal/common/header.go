package common

import "net/http"

// Header messages to be used in http.Header for client calls
const (
	HeaderAcceptMimeKey     = "Accept"
	HeaderAcceptMimeVal     = "application/json"
	HeaderAcceptEncodingKey = "Accept-Encoding"
	HeaderAcceptEncodingVal = "gzip, deflate, sdch"
)

// Default header accept for client calls
var (
	HeaderAccept http.Header = map[string][]string{
		HeaderAcceptMimeKey:     {HeaderAcceptMimeVal},
		HeaderAcceptEncodingKey: {HeaderAcceptEncodingVal},
	}
)
