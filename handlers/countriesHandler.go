package handler

import (
	//"encoding/json"
	"encoding/json"
	"log"
	"net/http"
	"regexp"

	"strconv"
	"strings"
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
	//splitting the url into different parts
	parts := strings.Split(r.URL.Path, "/")

	//checks if the url is valid
	if len(parts) != 6 || parts[3] != COUNTRIES {
		http.Error(w, "Malformed URL", http.StatusBadRequest)
		return
	}

	//find country first, then:
	//find all universities consisting of the given name, then
	//select only the universities that are in the correct countries and its neighbors

	//replacing empty space with %20

	var countryName = strings.ReplaceAll(parts[4], " ", "%20")
	var universityName = strings.ReplaceAll(parts[5], " ", "%20")

	var countryInfo = getCountryInfo(countryName, w, "name")
	var universities = getUniversityByNameAndCountry(universityName, countryInfo[0].CountryName.Name, w)

	for i := range countryInfo {
		for j := range countryInfo[i].Neighbours {
			var neighbour = getCountryInfo(countryInfo[i].Neighbours[j], w, "alpha")
			countryInfo = append(countryInfo, neighbour[0])
		}
	}

	//this if-statement checks whether or not a limit is provided,
	//if it is, it will add universities in neighbouring countries as long as the
	//number of universities don't exceed the limit,
	//and if not, it will add all the universities in the neighbouring countries (with the same name)
	if strings.Contains(r.URL.RawQuery, "limit") {
		limitString := r.URL.RawQuery
		re, err := regexp.Compile(`[^0-9]`) //todo: add error if
		if err != nil {
			log.Println("error compiling regex", err)
		}
		limitString = re.ReplaceAllString(limitString, "")
		log.Println("the limit is:" + limitString)
		limitInt, err := strconv.Atoi(limitString)
		if err != nil {
			log.Println("string to int conversion failed", err)
		}
		for i := range countryInfo {
			var universityNeighbour = getUniversityByNameAndCountry(universityName, countryInfo[i].CountryName.Name, w)
			if len(universityNeighbour) > 0 {
				for j := 0; j < limitInt; j++ {
					universities = append(universities, universityNeighbour[j])
					log.Println(universityNeighbour[j].Name)
				}
			}
		}

	} else {
		for i := range countryInfo {
			var universityNeighbour = getUniversityByNameAndCountry(universityName, countryInfo[i].CountryName.Name, w)
			if len(universityNeighbour) > 0 {
				for j := 0; j < len(universityNeighbour); j++ {
					universities = append(universities, universityNeighbour[j])
					log.Println(universityNeighbour[j].Name)
				}
			}
		}
	}

	//fetching and implementing information about country that was not retrieved from hipolabs
	for i := range universities {
		var additionalCountryInformation = getAdditionalCountryInformation(universities[i].CountryName, w)
		universities[i].AdditionCountryInformation.Languages = additionalCountryInformation[0].Languages
		universities[i].AdditionCountryInformation.LocationMap.OpenStreetMaps = additionalCountryInformation[0].LocationMap.OpenStreetMaps
	}

	//setting the format for the GET response
	json.Marshal(universities)
	w.Header().Add("content-type", "application/json")
	if err := json.NewEncoder(w).Encode(&universities); err != nil {
		log.Println("ERROR encoding JSON", err)
	}

}

//Method that performs a GET action to find information about a university given
//a string universityName and writes with http.Responsewriter w
//Returns an array of the struct University
func getCountryInfo(countryName string, w http.ResponseWriter, searchPath string) []Country {
	resp, err := http.Get("https://restcountries.com/v3.1/" + searchPath + "/" + countryName)
	if err != nil {
		log.Println("no response")
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
	}
	defer resp.Body.Close()

	var country []Country

	//decoding
	if err := json.NewDecoder(resp.Body).Decode(&country); err != nil {
		log.Println("Can not decode JSON", err)
		http.Error(w, "Unable decode", http.StatusBadRequest) //todo: set correct error status
	}

	return country
}

func getAdditionalCountryInformation(countryName string, w http.ResponseWriter) []AdditionCountryInformation {
	resp, err := http.Get("https://restcountries.com/v3.1/name/%7B" + countryName + "%7D")
	if err != nil {
		log.Println("no response")
		http.Error(w, "Requested resource was not found", http.StatusNotFound)
	}

	defer resp.Body.Close()

	var additionCountryInfo []AdditionCountryInformation

	if err := json.NewDecoder(resp.Body).Decode(&additionCountryInfo); err != nil {
		log.Println("Can not decode JSON", err)
		http.Error(w, "Unable decode", http.StatusBadRequest) //todo: set correct error status
	}

	return additionCountryInfo
}
