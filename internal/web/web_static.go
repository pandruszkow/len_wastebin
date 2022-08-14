package web

import (
	"net/http"
)

func (data Data) StyleCSSHand(rw http.ResponseWriter, req *http.Request) {
	data.Log.HttpRequest(req)

	rw.Header().Set("Content-Type", "text/css")
	rw.Write(*data.StyleCSS)
}

func (data Data) MainJSHand(rw http.ResponseWriter, req *http.Request) {
	data.Log.HttpRequest(req)

	rw.Header().Set("Content-Type", "application/javascript")
	rw.Write(*data.MainJS)
}

func (data Data) PasteJSHand(rw http.ResponseWriter, req *http.Request) {
	data.Log.HttpRequest(req)

	rw.Header().Set("Content-Type", "application/javascript")
	rw.Write(*data.PasteJS)
}
