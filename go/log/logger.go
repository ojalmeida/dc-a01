package log

import (
	"go-api/config"
	"log"
	"os"
)

var (
	errorLogger  *log.Logger
	infoLogger   *log.Logger
	errorChannel chan string
	infoChannel  chan string

	channelBuffer = config.Config.Performance.MaxNumberOfWorkers
)

// Start returns two channels where info and error logs must be sent to been stored
func Start() (chan<- string, chan<- string) {

	infoLogger = log.New(os.Stdout, "INFO (go-api): ", log.Ldate|log.Ltime|log.LUTC)

	errorLogger = log.New(os.Stderr, "ERROR (go-api): ", log.Ldate|log.Ltime|log.LUTC)

	errorChannel = make(chan string, channelBuffer)
	infoChannel = make(chan string, channelBuffer)

	go storeLogs()

	return infoChannel, errorChannel

}

// storeLogs listens to info and errors channels, logging information to stdout and stderr, respectively.
func storeLogs() {

	for {

		select {

		case info := <-infoChannel:

			infoLogger.Println(info)

		case err := <-errorChannel:

			errorLogger.Println(err)

		}

	}

}
