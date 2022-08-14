package raw

import (
	"io"
	"net/http"
)

func (data Data) errorNotFound(rw http.ResponseWriter, req *http.Request) {
	// Write response header
	rw.Header().Set("Content-type", "text/plain")
	rw.WriteHeader(404)

	// Send
	_, e := io.WriteString(rw, "404 Not Found")
	if e != nil {
		data.Log.HttpError(req, e)
	}
}

func (data Data) errorInternal(rw http.ResponseWriter, req *http.Request, err error) {
	// Write to log
	data.Log.HttpError(req, err)

	// Write response header
	rw.Header().Set("Content-type", "text/plain")
	rw.WriteHeader(500)

	// Send
	_, e := io.WriteString(rw, "500 Internal Server Error")
	if e != nil {
		data.Log.HttpError(req, e)
	}
}
