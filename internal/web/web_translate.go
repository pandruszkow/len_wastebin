package web

import (
	"bytes"
	"errors"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"path/filepath"
	"strconv"
	"strings"
)

type Locale map[string]string
type Locales map[string]*Locale

func loadLocales(localeDir string) (Locales, error) {
	locales := make(Locales)

	// Get locale files list
	files, err := ioutil.ReadDir(localeDir)
	if err != nil {
		return locales, err
	}

	for _, fileInfo := range files {
		// Check file
		if fileInfo.IsDir() {
			continue
		}

		fileName := fileInfo.Name()
		if strings.HasSuffix(fileName, ".locale") == false {
			continue
		}
		localeCode := fileName[:len(fileName)-7]

		// Read file
		filePath := filepath.Join(localeDir, fileName)
		fileByte, err := readFile(filePath)
		if err != nil {
			return locales, err
		}

		file := bytes.NewBuffer(fileByte).String()

		// Load locale
		locale := make(Locale)
		for num, str := range strings.Split(file, "\n") {
			if str == "" || strings.HasPrefix(str, "//") {
				continue
			}

			key := ""
			val := ""

			afterEqual := false
			for _, char := range []rune(str) {
				if afterEqual == false {
					if char == '=' {
						afterEqual = true

					} else {
						key = key + string(char)
					}

				} else {
					val = val + string(char)
				}
			}

			if afterEqual == false {
				return locales, errors.New("locale: error in line " + strconv.Itoa(num) + " in the file " + filePath)
			}

			locale[key] = val
		}

		locales[localeCode] = &locale
	}

	return locales, nil
}

func (locales Locales) findLocale(req *http.Request) Locale {
	// Get accept language by cookie
	langCookie := getCookie(req, "lang")
	if langCookie != "" {
		locale, ok := locales[langCookie]
		if ok == true {
			return *locale
		}
	}

	// Get user Accepr-Languages list
	acceptLanguage := req.Header.Get("Accept-Language")
	acceptLanguage = strings.Replace(acceptLanguage, " ", "", -1)

	var langs []string
	for _, part := range strings.Split(acceptLanguage, ";") {
		for _, lang := range strings.Split(part, ",") {
			if strings.HasPrefix(lang, "q=") == false {
				langs = append(langs, lang)
			}
		}
	}

	// Search locale
	for _, lang := range langs {
		for localeCode, locale := range locales {
			if localeCode == lang {
				return *locale
			}
		}
	}

	// Load default locale
	locale, ok := locales["en"]
	if ok != true {
		// If en locale not found load first locale
		for _, l := range locales {
			return *l
		}
	}

	return *locale
}

func (locale Locale) translate(s string, a ...interface{}) template.HTML {
	for key, val := range locale {
		if key == s {
			return template.HTML(fmt.Sprintf(val, a...))
		}
	}

	return template.HTML(s)
}
