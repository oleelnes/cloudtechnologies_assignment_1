package handler

type University struct {
	Name        string   `json:"name,omitempty"`
	CountryName string   `json:"country,omitempty"`
	IsoCode     string   `json:"alpha_two_code"`
	WebPages    []string `json:"web_pages,omitempty"`
	AdditionCountryInformation
}

type Country struct {
	CountryName string            `json:"name,omitempty"`
	Languages   map[string]string `json:"languages,omitempty"`
	IsoCode     string            `json:"isocode,omitempty"`
	LocationMap string            `json:"maps,omitempty"`
	Neighbours  []string          `json:"neighbors,omitempty"`
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
