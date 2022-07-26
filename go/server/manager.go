package server

import (
	"go-api/config"
	"net/http"
)

// pull inserts a request and his writer to the request queue to be processed
func (m *manager) pull(req *http.Request, writer http.ResponseWriter) chan bool {

	done := make(chan bool)

	data := item{

		req,
		writer,
		done,
	}

	m.queue <- data

	return done

}

// start inits queue and waits for new incoming items to be routed by routeRequest
func (m *manager) start() {

	m.queue = make(chan item, config.Config.Performance.MaxNumberOfWorkers)

	for {

		if m.activeWorkers < config.Config.Performance.MaxNumberOfWorkers {

			data := <-m.queue

			req := data.req
			writer := data.writer
			done := data.done

			m.activeWorkers++

			go routeRequest(req, writer, done, infoLog, errorLog)

		}

	}

}
