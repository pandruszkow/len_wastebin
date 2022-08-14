package apiv1

import (
	"encoding/json"
	"net/http"

	"github.com/coolguy1771/wastebin/internal/netshare"
)

type newPasteAnswer struct {
	ID string `json:"id"`
}

// POST /api/v1/new
func (data Data) NewHand(rw http.ResponseWriter, req *http.Request) {
	// Log request
	data.Log.HttpRequest(req)

	// Check method
	if req.Method != "POST" {
		data.writeError(rw, req, netshare.ErrMethodNotAllowed)
		return
	}

	// Get form data and create paste
	req.ParseForm()

	pasteID, err := netshare.PasteAddFromForm(req.PostForm, data.DB, *data.TitleMaxLen, *data.BodyMaxLen, *data.MaxLifeTime, *data.Lexers)
	if err != nil {
		if err == netshare.ErrBadRequest {
			data.writeError(rw, req, netshare.ErrBadRequest)
			return
		}

		data.writeError(rw, req, netshare.ErrInternal)
		return
	}

	// Return response
	rw.Header().Set("Content-Type", "application/json")

	err = json.NewEncoder(rw).Encode(newPasteAnswer{ID: pasteID})
	if err != nil {
		data.Log.HttpError(req, err)
		return
	}
}
