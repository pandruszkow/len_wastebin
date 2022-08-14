package web

import (
	"html/template"
	"net/http"

	"github.com/coolguy1771/wastebin/internal/netshare"
)

type createTmpl struct {
	TitleMaxLen int
	BodyMaxLen  int
	MaxLifeTime int64
	Lexers      []string

	AuthorDefault      string
	AuthorEmailDefault string
	AuthorURLDefault   string

	Translate func(string, ...interface{}) template.HTML
}

func (data Data) newPaste(rw http.ResponseWriter, req *http.Request) {
	// Read request
	req.ParseForm()

	if req.PostForm.Get("body") != "" {
		// Create paste
		pasteID, err := netshare.PasteAddFromForm(req.PostForm, data.DB, *data.TitleMaxLen, *data.BodyMaxLen, *data.MaxLifeTime, *data.Lexers)
		if err != nil {
			if err == netshare.ErrBadRequest {
				data.errorBadRequest(rw, req)
				return
			}

			data.errorInternal(rw, req, err)
			return
		}

		// Redirect to paste
		writeRedirect(rw, req, "/"+pasteID, 302)
		return
	}

	// Else show create page
	tmplData := createTmpl{
		TitleMaxLen:        *data.TitleMaxLen,
		BodyMaxLen:         *data.BodyMaxLen,
		MaxLifeTime:        *data.MaxLifeTime,
		Lexers:             *data.Lexers,
		AuthorDefault:      getCookie(req, "author"),
		AuthorEmailDefault: getCookie(req, "authorEmail"),
		AuthorURLDefault:   getCookie(req, "authorURL"),
		Translate:          data.Locales.findLocale(req).translate,
	}

	rw.Header().Set("Content-Type", "text/html")

	err := data.Main.Execute(rw, tmplData)
	if err != nil {
		data.errorInternal(rw, req, err)
		return
	}
}
