package handler

import (
	"net/http"
	"encoding/json"
)

// respondJSON制造JSON格式的工资应答。
func respondJSON(writer http.ResponseWriter, status int, payload interface{}) {
	response, err := json.Marshal(payload)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		writer.Write([]byte(err.Error()))
		return
	}
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(status)
	writer.Write([]byte(response))
	return
}

// respondError制造JSON格式的错误应答。
func respondError(writer http.ResponseWriter, code int, message string) {
	respondJSON(writer, code, map[string]string{"error": message})
	return
}
