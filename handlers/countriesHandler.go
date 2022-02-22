package handler

import (
	//"encoding/json"
	"net/http"
)

func CountriesHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		getCountry(w, r)
	default:
		http.Error(w, "Method not supported -- only GET is supported", http.StatusNotImplemented)
	}

}

func getCountry(w http.ResponseWriter, r *http.Request) {

	//decoder := json.NewDecoder(r.Body)

}
