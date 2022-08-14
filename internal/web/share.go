package web

import (
	"io/ioutil"
	"net/http"
	"os"
)

func readFile(path string) ([]byte, error) {
	// Open file
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Read file
	fileByte, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}

	return fileByte, nil
}

func getCookie(req *http.Request, name string) string {
	cookie, err := req.Cookie(name)
	if err != nil {
		return ""
	}

	return cookie.Value
}
