package logger

import (
	"fmt"
	"net/http"

	"github.com/fatih/color"
)

// https://go.dev/play/p/6E6grsIts9
// https://stackoverflow.com/questions/34017342/log-404-on-http-fileserver

type StatusRespWr struct {
	http.ResponseWriter // We embed http.ResponseWriter
	status              int
}

func (w *StatusRespWr) WriteHeader(status int) {
	w.status = status // Store the status for our own use
	w.ResponseWriter.WriteHeader(status)
}

func LoggingHandler(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		srw := &StatusRespWr{ResponseWriter: w}
		h.ServeHTTP(srw, r)
		if (srw.status >= 400) && (srw.status <= 599) {
			fmt.Print("\n"+r.Method+" ", r.URL.Path+" ")
			rederror := color.New(color.Bold, color.FgRed)
			rederror.Print(srw.status)
		} else {
			fmt.Print("\n"+r.Method+" ", r.URL.Path+" ")
			info := color.New(color.Bold, color.FgGreen)
			info.Print(srw.status)
		}
	})
}
