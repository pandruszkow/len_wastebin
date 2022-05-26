// Copyright (C) 2021-2022 Leonid Maslakov.

// This file is part of Lenpaste.

// Lenpaste is free software: you can redistribute it
// and/or modify it under the terms of the
// GNU Affero Public License as published by the
// Free Software Foundation, either version 3 of the License,
// or (at your option) any later version.

// Lenpaste is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of MERCHANTABILITY
// or FITNESS FOR A PARTICULAR PURPOSE.
// See the GNU Affero Public License for more details.

// You should have received a copy of the GNU Affero Public License along with Lenpaste.
// If not, see <https://www.gnu.org/licenses/>.

package web

import (
	"git.lcomrade.su/root/lenpaste/internal/netshare"
	"net/http"
)

type createTmpl struct {
	Lexers []string
}

func (data Data) newPaste(rw http.ResponseWriter, req *http.Request) {
	// Read request
	req.ParseForm()

	if req.PostForm.Get("body") != "" {
		// Create paste
		paste, err := netshare.PasteAddFromForm(data.DB, data.Lexers, req.PostForm)
		if err != nil {
			if err == netshare.ErrBadRequest {
				data.errorBadRequest(rw, req)
				return
			}

			data.errorInternal(rw, req, err)
			return
		}

		// Redirect to paste
		writeRedirect(rw, req, "/"+paste.ID, 302)
		return
	}

	// Else show create page
	tmplData := createTmpl{
		Lexers: data.Lexers,
	}

	rw.Header().Set("Content-Type", "text/html")

	err := data.Main.Execute(rw, tmplData)
	if err != nil {
		data.errorInternal(rw, req, err)
		return
	}
}
