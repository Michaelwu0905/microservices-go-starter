package main

import (
	"encoding/json"
	"net/http"
)

// 封装json响应写入逻辑
func writeJSON(w http.ResponseWriter, status int, data any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(data)
}
