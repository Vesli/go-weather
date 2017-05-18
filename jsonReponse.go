package main

import (
	"encoding/json"
	"net/http"
)

func writeJSON(w http.ResponseWriter, obj interface{}) {
	jData, err := json.Marshal(obj)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jData)
}
