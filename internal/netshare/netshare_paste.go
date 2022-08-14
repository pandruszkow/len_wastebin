package netshare

import (
	"net/url"
	"strconv"
	"strings"
	"time"

	"git.lcomrade.su/root/lineend"
	"github.com/coolguy1771/wastebin/internal/storage"
)

func PasteAddFromForm(form url.Values, db storage.DB, titleMaxLen int, bodyMaxLen int, maxLifeTime int64, lexerNames []string) (string, error) {
	// Read form
	paste := storage.Paste{
		Title:       form.Get("title"),
		Body:        form.Get("body"),
		Syntax:      form.Get("syntax"),
		DeleteTime:  0,
		OneUse:      false,
		Author:      form.Get("author"),
		AuthorEmail: form.Get("authorEmail"),
		AuthorURL:   form.Get("authorURL"),
	}

	// Remove new line from title
	paste.Title = strings.Replace(paste.Title, "\n", "", -1)
	paste.Title = strings.Replace(paste.Title, "\r", "", -1)

	// Check title
	if len(paste.Title) > titleMaxLen && titleMaxLen >= 0 {
		return "", ErrBadRequest
	}

	// Check paste body
	if paste.Body == "" {
		return "", ErrBadRequest
	}

	if len(paste.Body) > bodyMaxLen && bodyMaxLen > 0 {
		return "", ErrBadRequest
	}

	// Change paste body lines end
	switch form.Get("lineEnd") {
	case "", "LF", "lf":
		paste.Body = lineend.UnknownToUnix(paste.Body)

	case "CRLF", "crlf":
		paste.Body = lineend.UnknownToDos(paste.Body)

	case "CR", "cr":
		paste.Body = lineend.UnknownToOldMac(paste.Body)

	default:
		return "", ErrBadRequest
	}

	// Check syntax
	syntaxOk := false
	for _, name := range lexerNames {
		if name == paste.Syntax {
			syntaxOk = true
			break
		}
	}

	if !syntaxOk {
		return "", ErrBadRequest
	}

	// Get delete time
	expirStr := form.Get("expiration")
	if expirStr != "" {
		// Convert string to int
		expir, err := strconv.ParseInt(expirStr, 10, 64)
		if err != nil {
			return "", ErrBadRequest
		}

		// Check limits
		if maxLifeTime > 0 {
			if expir > maxLifeTime || expir <= 0 {
				return "", ErrBadRequest
			}
		}

		// Save if ok
		if expir > 0 {
			paste.DeleteTime = time.Now().Unix() + expir
		}
	}

	// Get "one use" parameter
	if form.Get("oneUse") == "true" {
		paste.OneUse = true
	}

	// Create paste
	pasteID, err := db.PasteAdd(paste)
	if err != nil {
		return pasteID, err
	}

	return pasteID, nil
}
