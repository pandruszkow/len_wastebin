package raw

import (
	"io"
	"net/http"

	"github.com/coolguy1771/wastebin/internal/storage"
)

// Pattern: /raw/
func (data Data) RawHand(rw http.ResponseWriter, req *http.Request) {
	// Log request
	data.Log.HttpRequest(req)

	// Read DB
	pasteID := string([]rune(req.URL.Path)[5:])

	paste, err := data.DB.PasteGet(pasteID)
	if err != nil {
		if err == storage.ErrNotFoundID {
			data.errorNotFound(rw, req)
			return

		} else {
			data.errorInternal(rw, req, err)
			return
		}
	}

	// If "one use" paste
	if paste.OneUse == true {
		// Delete paste
		err = data.DB.PasteDelete(pasteID)
		if err != nil {
			data.errorInternal(rw, req, err)
			return
		}
	}

	// Write result
	rw.Header().Set("Content-Type", "text/plain; charset=utf-8")

	_, err = io.WriteString(rw, paste.Body)
	if err != nil {
		data.errorInternal(rw, req, err)
		return
	}
}
