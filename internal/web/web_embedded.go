package web

import (
	"html/template"
	"net/http"
	"time"

	"github.com/coolguy1771/wastebin/internal/storage"
)

type embTmpl struct {
	ID            string
	CreateTimeStr string
	DeleteTime    int64
	OneUse        bool
	Title         string
	Body          template.HTML

	ErrorNotFound bool
	Translate     func(string, ...interface{}) template.HTML
}

// Pattern: /emb/
func (data Data) EmbeddedHand(rw http.ResponseWriter, req *http.Request) {
	errorNotFound := false

	// Log request
	data.Log.HttpRequest(req)

	// Get paste ID
	pasteID := string([]rune(req.URL.Path)[5:])

	// Read DB
	paste, err := data.DB.PasteGet(pasteID)
	if err != nil {
		if err == storage.ErrNotFoundID {
			errorNotFound = true

		} else {
			data.errorInternal(rw, req, err)
			return
		}
	}

	// Prepare template data
	createTime := time.Unix(paste.CreateTime, 0).UTC()

	tmplData := embTmpl{
		ID:            paste.ID,
		CreateTimeStr: createTime.Format("1 Jan, 2006"),
		DeleteTime:    paste.DeleteTime,
		OneUse:        paste.OneUse,
		Title:         paste.Title,
		Body:          tryHighlight(paste.Body, paste.Syntax),

		ErrorNotFound: errorNotFound,
		Translate:     data.Locales.findLocale(req).translate,
	}

	// Show paste
	err = data.EmbeddedPage.Execute(rw, tmplData)
	if err != nil {
		data.errorInternal(rw, req, err)
		return
	}
}
