package api

import (
	"encoding/json"
	"log/slog"
	"net/http"
)

func SendJSON(w http.ResponseWriter, response Response, statusCode int) {
	w.Header().Set("Content-Type", "application/json")

	data, err := json.Marshal(response)
	if err != nil {
		slog.Error("Failed to marshal json", "error", err)
		SendJSON(w, Response{Error: "Something went wrong..."}, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(statusCode)

	if _, err := w.Write(data); err != nil {
		slog.Error("Failed to write response to client", "error", err)
		return
	}
}
