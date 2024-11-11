package resp

import (
	"encoding/json"
	"net/http"
)

func JSON(w http.ResponseWriter, status int, v interface{}) {
	w.Header().Set("Content-Type", "application/json")

	data, err := json.Marshal(v)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
	w.WriteHeader(status)

	_, err = w.Write(data)
}

func Error(w http.ResponseWriter, status int, err string) {
	JSON(w, status, envelope{
		"error": err,
	})
}

type envelope map[string]interface{}
