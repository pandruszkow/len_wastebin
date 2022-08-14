package logger

import (
	"fmt"
	"net/http"
	"time"
)

type Config struct {
	TimeFormat string
}

func (config Config) Info(msg string) {
	now := time.Now()
	fmt.Println(now.Format(config.TimeFormat), "[INFO]", msg)
}

func (config Config) Error(err error) {
	now := time.Now()
	fmt.Println(now.Format(config.TimeFormat), "[ERROR]", err.Error())
}

func (config Config) HttpRequest(req *http.Request) {
	now := time.Now()
	fmt.Println(now.Format(config.TimeFormat), "[REQUEST]", req.RemoteAddr, req.Method, req.URL.Path)
}

func (config Config) HttpError(req *http.Request, err error) {
	now := time.Now()
	fmt.Println(now.Format(config.TimeFormat), "[ERROR]", req.RemoteAddr, req.Method, req.URL.Path, "Error:", err.Error())
}
