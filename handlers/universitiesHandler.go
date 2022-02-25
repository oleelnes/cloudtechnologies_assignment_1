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

//Method that is responsible for encoding University json entries through
//fetching information from external APIs
func getUniversity(w http.ResponseWriter, r *http.Request) {
	//splitting the url into different parts
	parts := strings.Split(r.URL.Path, "/")

	//checks if the url is valid
	if len(parts) != 5 || parts[3] != UNI_INFO {
		http.Error(w, "Malformed URL", http.StatusBadRequest)
		return
	}

	//replacing empty space with %20; replacing {} with empty space
	//todo: better way to solve this?
	var universityName = strings.ReplaceAll(parts[4], " ", "%20")

	//Information about universities retrieved from the below method
	var universities = getUniversityInformation(universityName, w)
	if universities == nil {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	//fetching and implementing information about country that was not retrieved from hipolabs
	for i := range universities {
		var additionalCountryInformation = getAdditionalCountryInformation(universities[i].CountryName, w)
		universities[i].AdditionCountryInformation.Languages = additionalCountryInformation[0].Languages
		universities[i].AdditionCountryInformation.LocationMap.OpenStreetMaps = additionalCountryInformation[0].LocationMap.OpenStreetMaps
	}

	json.Marshal(universities)
	w.Header().Add("content-type", "application/json")

	var universitiesReply []UniversityReply
	for i := range universities {
		var universityReply = UniversityReply{
			Name:        universities[i].Name,
			CountryName: universities[i].CountryName,
			IsoCode:     universities[i].IsoCode,
			WebPages:    universities[i].WebPages,
			Languages:   universities[i].Languages,
			LocationMap: universities[i].LocationMap.OpenStreetMaps,
		}
		universitiesReply = append(universitiesReply, universityReply)
	}

	//setting the format for the GET response
	json.Marshal(universitiesReply)
	w.Header().Add("content-type", "application/json")
	if err := json.NewEncoder(w).Encode(&universitiesReply); err != nil {
		log.Println("ERROR encoding JSON", err)
	}

}

//Method that performs a GET action to find information about a university given
//a string universityName and writes with http.Responsewriter w
//Returns an array of the struct University
func getUniversityInformation(universityName string, w http.ResponseWriter) []University {
	resp, err := http.Get("http://universities.hipolabs.com/search?name=" + universityName)
	if err != nil {
		log.Println("no response")
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
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

//Method that finds universities through a name and country search
//returns a slice of struct University
func getUniversityByNameAndCountry(universityName string, countryName string, w http.ResponseWriter) []University {
	resp, err := http.Get("http://universities.hipolabs.com/search?name=" + universityName + "&country=" + countryName)
	if err != nil {
		log.Println("no response")
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
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
