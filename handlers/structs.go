package handler

type University struct {
	Name        string   `json:"name,omitempty"`
	CountryName string   `json:"country,omitempty"`
	IsoCode     string   `json:"alpha_two_code"`
	WebPages    []string `json:"web_pages,omitempty"`
	AdditionCountryInformation
}

type Country struct {
	CountryName struct {
		Name string `json:"common,omitempty"`
	} `json:"name"`
	IsoCode    string   `json:"cca2,omitempty"`
	Neighbours []string `json:"borders,omitempty"`
	AdditionCountryInformation
}

type Diagnostics struct {
	Version string `json:"version,omitempty"`
}

type AdditionCountryInformation struct {
	LocationMap struct {
		OpenStreetMaps string `json:"openstreetmaps"`
	} `json:"maps"`
	Languages map[string]string `json:"languages"`
}
