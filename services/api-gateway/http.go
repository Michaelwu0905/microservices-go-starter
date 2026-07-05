package main

import (
	"encoding/json"
	"net/http"
	"ride-sharing/shared/contracts"
)

// 放置各种http handler

func handleTripPreview(w http.ResponseWriter, r *http.Request) {

	var reqBody previewTripRequest
	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil { // 解析body到reqBody
		http.Error(w, "failed to parse JSON data", http.StatusBadRequest)
		return
	}

	defer r.Body.Close()

	// Validation 简单校验
	if reqBody.UserID == "" {
		http.Error(w, "user ID is required", http.StatusBadRequest)
		return
	}

	// call trip service
	response := contracts.APIResponse{Data: "ok"}
	writeJSON(w, http.StatusCreated, response)

}
