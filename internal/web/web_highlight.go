package web

import (
	"bytes"
	"html/template"

	"github.com/alecthomas/chroma"
	"github.com/alecthomas/chroma/formatters/html"
	"github.com/alecthomas/chroma/lexers"
	"github.com/alecthomas/chroma/styles"
)

func tryHighlight(source string, lexer string) template.HTML {
	// Determine lexer
	l := lexers.Get(lexer)
	if l == nil {
		return template.HTML(source)
	}

	l = chroma.Coalesce(l)

	// Determine formatter
	f := html.New(
		html.Standalone(false),
		html.WithClasses(false),
		html.TabWidth(4),
		html.WithLineNumbers(true),
		html.WrapLongLines(true),
	)

	s := styles.Get("monokai")

	it, err := l.Tokenise(nil, source)
	if err != nil {
		return template.HTML(source)
	}

	// Format
	var buf bytes.Buffer

	err = f.Format(&buf, s, it)
	if err != nil {
		return template.HTML(source)
	}

	return template.HTML(buf.String())
}
