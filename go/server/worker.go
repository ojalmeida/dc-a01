package server

import (
	"context"
	"fmt"
	"go-api/config"
	"net/http"
	"strings"
	"time"
)

// routeRequest iterate over routes and, if matched any, delegates handling
func routeRequest(req *http.Request, writer http.ResponseWriter, done chan<- bool, info, error chan<- string) {

	var reqInfo requestInfo

	ctx := req.Context()

	ctx, cancel := context.WithTimeout(req.Context(), time.Duration(config.Config.Timeout.Wait)*time.Second)
	req = req.WithContext(ctx)
	defer cancel()

	var handled bool
	for _, r := range routes {

		if r.match(req) {

			reqInfo = r.handle(writer, req)
			handled = true
			break

		}

	}

	if !handled {

		reqInfo = notFoundHandler(writer, req)

	}

	// {id: string} {handlerName: string} {ip: string} {method: string} {URL: string} {send: string} {referer: string} {userAgent: string} {status: int} {processingTime: float64}

	info <- fmt.Sprintf(`%s %s %s %s "%s" "%s", "%s" "%s" %d %.3f`,

		reqInfo.ID.String(),
		reqInfo.Handler,
		reqInfo.R.Header.Get("X-Forwarded-For"),
		reqInfo.R.Method,
		reqInfo.R.URL.String(),
		reqInfo.SentData,
		reqInfo.R.Header.Get("Referer"),
		reqInfo.R.Header.Get("User-Agent"),
		reqInfo.ResponseStatus,
		float64(reqInfo.ProcessingTime.Milliseconds())/1000,
	)

	if len(reqInfo.Errors) > 0 {

		var e []string

		for _, err := range reqInfo.Errors {

			e = append(e, fmt.Sprintf(`"%s"`, err.Error()))

		}

		error <- fmt.Sprintf(`%s | %s`, reqInfo.ID.String(), strings.Join(e, ","))

	}

	m.activeWorkers--
	done <- true

}
