package netshare

import (
	"net/http"
)

func GetHost(req *http.Request) string {
	// Read header
	xHost := req.Header.Get("X-Forwarded-Host")

	// Check
	if xHost != "" {
		return xHost
	}

	return req.Host
}

func GetProtocol(header http.Header) string {
	// Read header
	xProto := header.Get("X-Forwarded-Proto")

	// Check
	if xProto != "" {
		return xProto
	}

	return "http"
}
