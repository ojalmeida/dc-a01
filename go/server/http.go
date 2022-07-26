package server

import (
	"fmt"
	"go-api/config"
	"net/http"
	"time"
)

var (
	m        *manager
	s        *http.Server
	infoLog  chan<- string
	errorLog chan<- string
)

func Start(i, e chan<- string) {

	s = &http.Server{

		Addr:              fmt.Sprintf("%s:%d", config.Config.Server.Addr, config.Config.Server.Port),
		ErrorLog:          nil,
		ReadHeaderTimeout: time.Duration(config.Config.Timeout.Read) * time.Second,
		WriteTimeout:      time.Duration(config.Config.Timeout.Write) * time.Second,
		IdleTimeout:       time.Duration(config.Config.Timeout.Read) * time.Second,
	}

	m = &manager{}

	infoLog = i
	errorLog = e

	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {

		done := m.pull(request, writer)

		<-done

	})

	go m.start()

	if err := s.ListenAndServe(); err != http.ErrServerClosed {

		errorLog <- err.Error()

		panic(err.Error())

	}

}
