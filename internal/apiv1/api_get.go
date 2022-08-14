package apiv1

import (
	"encoding/json"
	"net/http"

	"github.com/coolguy1771/wastebin/internal/netshare"
	"github.com/coolguy1771/wastebin/internal/storage"
)

// GET /api/v1/get
func (data Data) GetHand(rw http.ResponseWriter, req *http.Request) {
	// Log request
	data.Log.HttpRequest(req)

	// Check method
	if req.Method != "GET" {
		data.writeError(rw, req, netshare.ErrMethodNotAllowed)
		return
	}

	// Get paste ID
	req.ParseForm()

	pasteID := req.Form.Get("id")

	// Check paste id
	if pasteID == "" {
		data.writeError(rw, req, netshare.ErrBadRequest)
		return
	}

	// Get paste
	paste, err := data.DB.PasteGet(pasteID)
	if err != nil {
		data.writeError(rw, req, err)
		return
	}

	// If "one use" paste
	if paste.OneUse == true {
		if req.Form.Get("openOneUse") == "true" {
			// Delete paste
			err = data.DB.PasteDelete(pasteID)
			if err != nil {
				data.writeError(rw, req, err)
				return
			}

		} else {
			// Remove secret data
			paste = storage.Paste{
				ID:     paste.ID,
				OneUse: true,
			}
		}
	}

	// Return response
	rw.Header().Set("Content-Type", "application/json")

	err = json.NewEncoder(rw).Encode(paste)
	if err != nil {
		data.Log.HttpError(req, err)
		return
	}
}
