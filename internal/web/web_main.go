package web

import (
	"net/http"
)

// Pattern: /
func (data Data) MainHand(rw http.ResponseWriter, req *http.Request) {
	// Log request
	data.Log.HttpRequest(req)

	// If main page
	if req.URL.Path == "/" {
		data.newPaste(rw, req)

		// Else show paste
	} else {
		data.getPaste(rw, req)
	}
}
