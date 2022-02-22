package handler

import (
	//"encoding/json"
	"net/http"
)

func UniversitiesHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		getUniversity(w, r)
	default:
		http.Error(w, "Method not supported -- only GET is supported", http.StatusNotImplemented)
	}

}

func getUniversity(w http.ResponseWriter, r *http.Request) {

}
