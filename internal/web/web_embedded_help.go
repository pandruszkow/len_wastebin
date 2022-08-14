package web

import (
	"html/template"
	"net/http"

	"github.com/coolguy1771/wastebin/internal/netshare"
	"github.com/coolguy1771/wastebin/internal/storage"
)

type embHelpTmpl struct {
	ID         string
	DeleteTime int64
	OneUse     bool

	Protocol string
	Host     string

	Translate func(string, ...interface{}) template.HTML
	Highlight func(string, string) template.HTML
}

// Pattern: /emb_help/
func (data Data) EmbeddedHelpHand(rw http.ResponseWriter, req *http.Request) {
	// Log request
	data.Log.HttpRequest(req)

	// Get paste ID
	pasteID := string([]rune(req.URL.Path)[10:])

	// Read DB
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

	// Show paste
	tmplData := embHelpTmpl{
		ID:         paste.ID,
		DeleteTime: paste.DeleteTime,
		OneUse:     paste.OneUse,
		Protocol:   netshare.GetProtocol(req.Header),
		Host:       netshare.GetHost(req),
		Translate:  data.Locales.findLocale(req).translate,
		Highlight:  tryHighlight,
	}

	err = data.EmbeddedHelpPage.Execute(rw, tmplData)
	if err != nil {
		data.errorInternal(rw, req, err)
		return
	}
}
