package main

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"ride-sharing/shared/contracts"
)

func handleTripPreview(w http.ResponseWriter, r *http.Request) {
	var reqBody previewTripRequest
	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		http.Error(w, "failed to parse JSON data", http.StatusBadRequest)
		return
	}

	defer r.Body.Close()

	if reqBody.UserID == "" {
		http.Error(w, "user ID is required", http.StatusBadRequest)
		return
	}

	jsonBody, _ := json.Marshal(reqBody)
	reader := bytes.NewBuffer(jsonBody)

	res, err := http.Post("http://trip-service:8083/preview", "application/json", reader)
	if err != nil {
		log.Print(err)
		return
	}

	defer res.Body.Close()

	var resBody any
	if err := json.NewDecoder(res.Body).Decode(&resBody); err != nil {
		http.Error(w, "failed to parse JSON data form trip service", http.StatusBadRequest)
		return
	}

	response := contracts.APIResponse{Data: resBody}

	writeJSON(w, http.StatusCreated, response)
}
