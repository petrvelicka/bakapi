package bakapi

import (
	"encoding/xml"
	"net/url"
)

type PredvidacResults struct {
	XMLName xml.Name `xml:"results"`
	Typypru Typypru  `xml:"typypru"`
	Result  int      `xml:"result"`
}

type Typypru struct {
	XMLName xml.Name `xml:"typypru"`
	Typ     []Typ    `xml:"typ"`
}

type Typ struct {
	XMLName xml.Name `xml:"typ"`
	Zkratka string   `xml:"zkratka"`
	Nazev   string   `xml:"nazev"`
	Vaha    int      `xml:"vaha"`
}

func GetPredvidac(user User) PredvidacResults {
	params := url.Values{
		"hx": {user.Token},
		"pm": {"predvidac"},
	}

	res := getRequest(user.Address, params)

	var predvidac PredvidacResults
	xml.Unmarshal(res, &predvidac)
	return predvidac
}
