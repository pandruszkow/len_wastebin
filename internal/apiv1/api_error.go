package apiv1

import (
	"encoding/json"
	"net/http"

	"github.com/coolguy1771/wastebin/internal/netshare"
	"github.com/coolguy1771/wastebin/internal/storage"
)

type errorType struct {
	Code  int    `json:"code"`
	Error string `json:"error"`
}

func (data Data) writeError(rw http.ResponseWriter, req *http.Request, e error) {
	var resp errorType

	if e == netshare.ErrBadRequest {
		resp.Code = 400
		resp.Error = "Bad Request"

	} else if e == storage.ErrNotFoundID {
		resp.Code = 404
		resp.Error = "Could not find ID"

	} else if e == netshare.ErrNotFound {
		resp.Code = 404
		resp.Error = "Not Found"

	} else if e == netshare.ErrMethodNotAllowed {
		resp.Code = 405
		resp.Error = "Method Not Allowed"

	} else {
		resp.Code = 500
		resp.Error = "Internal Server Error"
		data.Log.HttpError(req, e)
	}

	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(resp.Code)

	err := json.NewEncoder(rw).Encode(resp)
	if err != nil {
		data.Log.HttpError(req, err)
		return
	}
}
