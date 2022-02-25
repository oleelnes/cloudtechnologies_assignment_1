package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

var StartupTime = time.Now()

func DiagHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		getDiag(w, r)
	default:
		http.Error(w, "Method not supported -- only GET is supported", http.StatusNotImplemented)
	}

}

//This method encodes the diag response
func getDiag(w http.ResponseWriter, r *http.Request) {
	var diag = Diagnostics{
		UniversityStatus: getExternalAPIStatus(UNI_API),
		CountryStatus:    getExternalAPIStatus(COUNTRIES_API),
		UpTime:           int(time.Since(StartupTime).Seconds()),
		Version:          VERSION,
	}

	//setting the format for the GET response
	json.Marshal(diag)
	w.Header().Add("content-type", "application/json")
	if err := json.NewEncoder(w).Encode(&diag); err != nil {
		log.Println("ERROR encoding JSON", err)
	}
}

//This method checks the status of an external API's site
//Returns an int which represents the status
func getExternalAPIStatus(APIpath string) int {
	resp, err := http.Get(APIpath)
	if err != nil || resp.StatusCode != http.StatusOK {
		return http.StatusBadGateway
	}
	return http.StatusOK
}
