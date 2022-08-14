package web

import (
	"html/template"
	"net/http"
)

type docsTmpl struct {
	Highlight func(string, string) template.HTML
	Translate func(string, ...interface{}) template.HTML
}

// Pattern: /docs
func (data Data) DocsHand(rw http.ResponseWriter, req *http.Request) {
	data.Log.HttpRequest(req)

	rw.Header().Set("Content-Type", "text/html")
	err := data.Docs.Execute(rw, docsTmpl{Translate: data.Locales.findLocale(req).translate})
	if err != nil {
		data.errorInternal(rw, req, err)
		return
	}
}

// Pattern: /docs/apiv1
func (data Data) DocsApiV1Hand(rw http.ResponseWriter, req *http.Request) {
	data.Log.HttpRequest(req)

	rw.Header().Set("Content-Type", "text/html")
	err := data.DocsApiV1.Execute(rw, docsTmpl{
		Translate: data.Locales.findLocale(req).translate,
		Highlight: tryHighlight,
	})
	if err != nil {
		data.errorInternal(rw, req, err)
		return
	}
}

// Pattern: /docs/api_libs
func (data Data) DocsApiLibsHand(rw http.ResponseWriter, req *http.Request) {
	data.Log.HttpRequest(req)

	rw.Header().Set("Content-Type", "text/html")
	err := data.DocsApiLibs.Execute(rw, docsTmpl{Translate: data.Locales.findLocale(req).translate})
	if err != nil {
		data.errorInternal(rw, req, err)
		return
	}
}
