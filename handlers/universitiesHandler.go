package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"
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
	parts := strings.Split(r.URL.Path, "/")
	// error handling
	if len(parts) != 5 || parts[3] != UNI_INFO {
		http.Error(w, "Malformed URL", http.StatusBadRequest)
		return
	}

	var universities []University

	universities = getUniversityInformation(parts[4], w)

	for i := range universities {
		var countryName = universities[i].CountryName
		log.Println(countryName)
		var additionalCountryInformation []AdditionCountryInformation
		additionalCountryInformation = getCountryInformation(countryName, w)
		universities[i].AdditionCountryInformation.Languages = additionalCountryInformation[0].Languages
		universities[i].AdditionCountryInformation.LocationMap.OpenStreetMaps = additionalCountryInformation[0].LocationMap.OpenStreetMaps

	}
	json.Marshal(universities) // are this one and ..

	w.Header().Add("content-type", "application/json") // this one needed?

	if err := json.NewEncoder(w).Encode(&universities); err != nil {
		log.Println("ERROR encoding JSON", err)
	}

}

func getUniversityInformation(universityName string, w http.ResponseWriter) []University {
	resp, err := http.Get("http://universities.hipolabs.com/search?name=" + universityName)
	if err != nil {
		log.Println("no response")
	}

	defer resp.Body.Close()

	var universities []University

	//decoding
	if err := json.NewDecoder(resp.Body).Decode(&universities); err != nil {
		log.Println("Can not decode JSON", err)
		http.Error(w, "Unable decode", http.StatusBadRequest) //todo: set correct error status
	}

	return universities
}

func getCountryInformation(countryName string, w http.ResponseWriter) []AdditionCountryInformation {
	resp, err := http.Get("https://restcountries.com/v3.1/name/%7B" + countryName + "%7D")
	if err != nil {
		log.Println("no response")
	}

	defer resp.Body.Close()

	var additionCountryInfo []AdditionCountryInformation

	if err := json.NewDecoder(resp.Body).Decode(&additionCountryInfo); err != nil {
		log.Println("Can not decode JSON", err)
		//http.Error(w, "Unable decode", http.StatusBadRequest) //todo: set correct error status
	}

	return additionCountryInfo
}
