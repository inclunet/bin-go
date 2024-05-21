package server

import (
	"encoding/json"
	"errors"
	"net/http"
)

type Response struct {
	StatusCode int
	Body       any `json:"body,omitempty"`
}

func (r *Response) SendHasFile(w http.ResponseWriter) error {
	w.WriteHeader(r.StatusCode)
	w.Header().Set("Content-Type", "application/json")

	if r.Body == nil {
		return errors.New("response body is empty")
	}

	if r.StatusCode >= 400 {
		body, ok := r.Body.(*ErrorBody)

		if !ok {
			return errors.New("response body is not an error")
		}

		_, err := w.Write([]byte(body.Message))

		return err
	}

	_, err := w.Write(r.Body.([]byte))

	return err
}

func (r *Response) SendHasJson(w http.ResponseWriter) error {
	w.WriteHeader(r.StatusCode)
	w.Header().Set("Content-Type", "application/json")

	if r.Body == nil {
		return errors.New("response body is empty")
	}

	return json.NewEncoder(w).Encode(r.Body)
}
