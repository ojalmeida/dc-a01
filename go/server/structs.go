package server

import (
	"encoding/json"
	"github.com/google/uuid"
	"net/http"
	"regexp"
	"sync"
	"time"
)

type item struct {
	req    *http.Request
	writer http.ResponseWriter
	done   chan bool
}

type manager struct {

	/* queue is a slice of a http.Request that contains request data, a http.ResponseWriter that provides communication
	with the client and a boolean channel that carries the conclusion of request processing */
	queue         chan item
	queueMutex    sync.Mutex
	activeWorkers int
}

// requestInfo carries data of requests, such processing time and written body, to upper layers
type requestInfo struct {
	R              *http.Request
	ID             uuid.UUID
	ResponseStatus int
	SentData       []byte
	Errors         []error
	ProcessingTime time.Duration
	Handler        string
}

type response struct {
	Status int             `json:"status,omitempty"`
	Data   json.RawMessage `json:"data,omitempty"`
	Msg    string          `json:"msg,omitempty"`
}

type route struct {
	pattern *regexp.Regexp
	method  *regexp.Regexp
	fn      func(w http.ResponseWriter, r *http.Request) requestInfo
}

func (r route) handle(writer http.ResponseWriter, req *http.Request) (reqInfo requestInfo) {

	reqInfo = r.fn(writer, req)

	return

}

func (r route) match(req *http.Request) bool {

	return r.pattern.MatchString(req.URL.RequestURI()) && r.method.MatchString(req.Method)

}
