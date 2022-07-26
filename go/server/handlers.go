package server

import (
	"encoding/json"
	"github.com/google/uuid"
	"go-api/db"
	"go-api/entities"
	"io/ioutil"
	"net/http"
	"time"
)

func insertItemHandler(w http.ResponseWriter, r *http.Request) (reqInfo requestInfo) {

	var err error
	var errs []error
	var clientResponse response
	var id uuid.UUID
	var out []byte
	var processingTime time.Duration
	var init time.Time

	defer func() {

		processingTime = time.Now().Sub(init)

		reqInfo = requestInfo{
			R:              r,
			ID:             id,
			ResponseStatus: clientResponse.Status,
			SentData:       out,
			Errors:         errs,
			ProcessingTime: processingTime,
			Handler:        "insertItemHandler",
		}

	}()

	type requestStruct struct {
		Text string `json:"text,omitempty"`
	}

	rs := requestStruct{}

	init = time.Now()

	id = uuid.New()

	defer r.Body.Close()

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "*")

	// CORS
	if r.Method == http.MethodOptions {

		return

	}

	body, err := ioutil.ReadAll(r.Body)

	err = json.Unmarshal(body, &rs)

	if err != nil {
		errs = append(errs, err)
		clientResponse.Status = http.StatusBadRequest
		clientResponse.Msg = "wrong json format"
		return
	}

	item := entities.Item{
		Text: rs.Text,
		Date: time.Now(),
	}

	insertedItem, err := db.CreateItem(r.Context(), item)

	if err != nil {
		errs = append(errs, err)
		clientResponse.Status = http.StatusInternalServerError
		clientResponse.Msg = "an error occurred"
		return
	}

	clientResponse.Status = http.StatusCreated
	clientResponse.Data, err = json.Marshal(insertedItem)

	if err != nil {
		errs = append(errs, err)
	}

	return

}

// notFoundHandler is called when a HTTP request does not match any defined routes.
func notFoundHandler(w http.ResponseWriter, r *http.Request) (reqInfo requestInfo) {

	var err error
	var errs []error
	var res response
	var id uuid.UUID
	var out []byte
	var processingTime time.Duration

	init := time.Now()

	defer func() {

		reqInfo = requestInfo{
			R:              r,
			ID:             id,
			ResponseStatus: res.Status,
			SentData:       out,
			Errors:         errs,
			ProcessingTime: processingTime,
			Handler:        "notFoundHandler",
		}

	}()

	defer func() {

		processingTime = time.Now().Sub(init)

		out, err = json.Marshal(res)

		if err != nil {
			w.WriteHeader(http.StatusBadGateway)
			return
		}

		if res.Status != 200 {

			w.WriteHeader(res.Status)

		}

		_, err = w.Write(out)

		if err != nil {
			w.WriteHeader(http.StatusBadGateway)
			return
		}

	}()

	res.Status = 404
	res.Msg = "not found"

	return

}
