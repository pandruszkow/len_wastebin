package web

import (
	"html/template"
	"net/http"
)

type aboutTmpl struct {
	Version     string
	TitleMaxLen int
	BodyMaxLen  int

	ServerAbout template.HTML
	ServerRules template.HTML

	AdminName string
	AdminMail string

	Translate func(string, ...interface{}) template.HTML
}

type aboutMinTmp struct {
	Translate func(string, ...interface{}) template.HTML
}

// Pattern: /about
func (data Data) AboutHand(rw http.ResponseWriter, req *http.Request) {
	// Log request
	data.Log.HttpRequest(req)

	// Prepare data
	dataTmpl := aboutTmpl{
		Version:     *data.Version,
		TitleMaxLen: *data.TitleMaxLen,
		BodyMaxLen:  *data.BodyMaxLen,
		ServerAbout: template.HTML(*data.ServerAbout),
		ServerRules: template.HTML(*data.ServerRules),
		AdminName:   *data.AdminName,
		AdminMail:   *data.AdminMail,
		Translate:   data.Locales.findLocale(req).translate,
	}

	// Show page
	rw.Header().Set("Content-Type", "text/html")

	err := data.About.Execute(rw, dataTmpl)
	if err != nil {
		data.errorInternal(rw, req, err)
	}
}

// Pattern: /about/license
func (data Data) LicenseHand(rw http.ResponseWriter, req *http.Request) {
	// Log request
	data.Log.HttpRequest(req)

	// Show page
	rw.Header().Set("Content-Type", "text/html")

	err := data.License.Execute(rw, aboutMinTmp{Translate: data.Locales.findLocale(req).translate})
	if err != nil {
		data.errorInternal(rw, req, err)
	}
}

// Pattern: /about/source_code
func (data Data) SourceCodePageHand(rw http.ResponseWriter, req *http.Request) {
	// Log request
	data.Log.HttpRequest(req)

	// Show page
	rw.Header().Set("Content-Type", "text/html")

	err := data.SourceCodePage.Execute(rw, aboutMinTmp{Translate: data.Locales.findLocale(req).translate})
	if err != nil {
		data.errorInternal(rw, req, err)
	}
}
