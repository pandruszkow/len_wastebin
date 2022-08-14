package web

import (
	"html/template"
	"path/filepath"

	chromaLexers "github.com/alecthomas/chroma/lexers"
	"github.com/coolguy1771/wastebin/internal/config"
	"github.com/coolguy1771/wastebin/internal/logger"
	"github.com/coolguy1771/wastebin/internal/storage"
)

type Data struct {
	DB  storage.DB
	Log logger.Config

	Lexers  *[]string
	Locales *Locales

	StyleCSS       *[]byte
	ErrorPage      *template.Template
	Main           *template.Template
	MainJS         *[]byte
	PastePage      *template.Template
	PasteJS        *[]byte
	PasteContinue  *template.Template
	Settings       *template.Template
	About          *template.Template
	License        *template.Template
	SourceCodePage *template.Template

	Docs        *template.Template
	DocsApiV1   *template.Template
	DocsApiLibs *template.Template

	EmbeddedPage     *template.Template
	EmbeddedHelpPage *template.Template

	Version *string

	TitleMaxLen *int
	BodyMaxLen  *int
	MaxLifeTime *int64

	ServerAbout *string
	ServerRules *string

	AdminName *string
	AdminMail *string

	RobotsDisallow *bool
}

func Load(cfg config.Config, webDir string) (Data, error) {
	var data Data
	var err error

	// Setup base info
	data.DB = cfg.DB
	data.Log = cfg.Log

	data.Version = &cfg.Version

	data.TitleMaxLen = &cfg.TitleMaxLen
	data.BodyMaxLen = &cfg.BodyMaxLen
	data.MaxLifeTime = &cfg.MaxLifeTime

	data.ServerAbout = &cfg.ServerAbout
	data.ServerRules = &cfg.ServerRules

	data.AdminName = &cfg.AdminName
	data.AdminMail = &cfg.AdminMail

	data.RobotsDisallow = &cfg.RobotsDisallow

	// Get Chroma lexers
	lexers := chromaLexers.Names(false)
	data.Lexers = &lexers

	// Load locales
	locales, err := loadLocales(filepath.Join(webDir, "locale"))
	if err != nil {
		return data, err
	}
	data.Locales = &locales

	// style.css file
	styleCSS, err := readFile(filepath.Join(webDir, "style.css"))
	if err != nil {
		return data, err
	}
	data.StyleCSS = &styleCSS

	// main.tmpl
	data.Main, err = template.ParseFiles(
		filepath.Join(webDir, "base.tmpl"),
		filepath.Join(webDir, "main.tmpl"),
	)
	if err != nil {
		return data, err
	}

	// main.js
	mainJS, err := readFile(filepath.Join(webDir, "main.js"))
	if err != nil {
		return data, err
	}
	data.MainJS = &mainJS

	// paste.tmpl
	data.PastePage, err = template.ParseFiles(
		filepath.Join(webDir, "base.tmpl"),
		filepath.Join(webDir, "paste.tmpl"),
	)
	if err != nil {
		return data, err
	}

	// paste.js
	pasteJS, err := readFile(filepath.Join(webDir, "paste.js"))
	if err != nil {
		return data, err
	}
	data.PasteJS = &pasteJS

	// paste_continue.tmpl
	data.PasteContinue, err = template.ParseFiles(
		filepath.Join(webDir, "base.tmpl"),
		filepath.Join(webDir, "paste_continue.tmpl"),
	)
	if err != nil {
		return data, err
	}

	// settings.tmpl
	data.Settings, err = template.ParseFiles(
		filepath.Join(webDir, "base.tmpl"),
		filepath.Join(webDir, "settings.tmpl"),
	)
	if err != nil {
		return data, err
	}

	// about.tmpl
	data.About, err = template.ParseFiles(
		filepath.Join(webDir, "base.tmpl"),
		filepath.Join(webDir, "about.tmpl"),
	)
	if err != nil {
		return data, err
	}

	// license.tmpl
	data.License, err = template.ParseFiles(
		filepath.Join(webDir, "base.tmpl"),
		filepath.Join(webDir, "license.tmpl"),
	)
	if err != nil {
		return data, err
	}

	// source_code.tmpl
	data.SourceCodePage, err = template.ParseFiles(
		filepath.Join(webDir, "base.tmpl"),
		filepath.Join(webDir, "source_code.tmpl"),
	)
	if err != nil {
		return data, err
	}

	// docs.tmpl
	data.Docs, err = template.ParseFiles(
		filepath.Join(webDir, "base.tmpl"),
		filepath.Join(webDir, "docs.tmpl"),
	)
	if err != nil {
		return data, err
	}

	// docs_apiv1.tmpl
	data.DocsApiV1, err = template.ParseFiles(
		filepath.Join(webDir, "base.tmpl"),
		filepath.Join(webDir, "docs_apiv1.tmpl"),
	)
	if err != nil {
		return data, err
	}

	// docs_api_libs.tmpl
	data.DocsApiLibs, err = template.ParseFiles(
		filepath.Join(webDir, "base.tmpl"),
		filepath.Join(webDir, "docs_api_libs.tmpl"),
	)
	if err != nil {
		return data, err
	}

	// error.tmpl
	data.ErrorPage, err = template.ParseFiles(
		filepath.Join(webDir, "base.tmpl"),
		filepath.Join(webDir, "error.tmpl"),
	)
	if err != nil {
		return data, err
	}

	// emb.tmpl
	data.EmbeddedPage, err = template.ParseFiles(
		filepath.Join(webDir, "emb.tmpl"),
	)
	if err != nil {
		return data, err
	}

	// emb_help.tmpl
	data.EmbeddedHelpPage, err = template.ParseFiles(
		filepath.Join(webDir, "base.tmpl"),
		filepath.Join(webDir, "emb_help.tmpl"),
	)
	if err != nil {
		return data, err
	}

	return data, nil
}
