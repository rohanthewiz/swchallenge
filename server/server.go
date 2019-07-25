package server

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"swchallenge/logger"
)

type requestParams struct {
	Username string `json:"username"`
	Timestamp int64 `json:"unix_timestamp"`
	EventUUID string `json:"event_uuid"`
	IPAddress string `json:"ip_address"`
}

func HandleV1(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/v1" && r.URL.Path != "/v1/" {
		http.Error(w, "Sorry, we could not find that page", http.StatusNotFound)
		return
	}

	if r.Method == "POST" {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			logger.LogErr(err)
			http.Error(w, "Error when reading the request", 500)
			return
		}
		defer r.Body.Close()

		var reqParams requestParams
		err = json.Unmarshal(body, &reqParams)
		if err != nil {
			logger.LogErr(err)
			http.Error(w, "Error when reading the request", 500)
			return
		}
		fmt.Printf("%#v\n", reqParams)

		// Response
		resp, err := json.Marshal(reqParams)
		if err != nil {
			logger.LogErr(err)
			http.Error(w, "Error when processing the response", 500)
			return
		}
		w.Header().Set("content-type", "application/json")
		_, _ = w.Write(resp)

	} else {
		fmt.Fprint(w, "Sorry, only the POST method is supported")
	}
}