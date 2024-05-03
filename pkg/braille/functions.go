package braille

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func GetUrlIntParam(r *http.Request, param string, defaultValue int) int {
	if value, err := strconv.Atoi(GetUrlStringParam(r, param, "")); err == nil {
		return value
	}

	return defaultValue
}

func GetQueryString(r *http.Request, key, defaultValue string) string {
	query := r.URL.Query()

	if value, ok := query[key]; ok {
		return value[0]
	}

	return defaultValue
}

func GetUrlStringParam(r *http.Request, param, defaultValue string) string {
	vars := mux.Vars(r)

	if v, ok := vars[param]; ok {
		return v
	}

	return defaultValue
}

func Repply(w http.ResponseWriter, data any) {
	response, err := json.Marshal(data)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}
