package handler

import (
	//"encoding/json"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
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

	parts := strings.Split(r.URL.Path, "/")
	// error handling
	if len(parts) != 5 || parts[3] != COUNTRIES {
		http.Error(w, "Malformed URL", http.StatusBadRequest)
		log.Println("didn't work it seems")
		log.Println("LEN PARTS: ", len(parts))
		log.Println(parts[3] + " " + COUNTRIES)
		return
	}

	//decoder := json.NewDecoder(r.Body)

	resp, err := http.Get("http://universities.hipolabs.com/search?country=" + parts[4])
	if err != nil {
		log.Println("no response")
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, "No neighbouring countries", http.StatusBadRequest) //possibly print all?
		return
	}

	var country Country
	if err := json.Unmarshal(body, &country); err != nil {
		fmt.Println("Can not unmarshal JSON")
	}

	response := country.CountryName + "\n" + country.IsoCode

	w.Write([]byte(response))

	fmt.Println(string(body))
	w.Write(body)

}
