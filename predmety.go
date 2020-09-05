package bakapi

import (
	"encoding/xml"
	"net/url"
)

type PredmetyResults struct {
	XMLName  xml.Name `xml:"results"`
	Predmety Predmety `xml:"predmety"`
	Result   int      `xml:"result"`
}

type Predmety struct {
	XMLName xml.Name  `xml:"predmety"`
	Predmet []Subject `xml:"predmet"`
}

type Subject struct {
	XMLName   xml.Name `xml:"predmet"`
	Nazev     string   `xml:"nazev"`
	Zkratka   string   `xml:"zkratka"`
	KodPred   string   `xml:"kod_pred"`
	Ucitel    string   `xml:"ucitel"`
	ZkratkaUc string   `xml:"zkratkauc"`
}

func GetPredmety(user User) PredmetyResults {
	params := url.Values{
		"hx": {user.Token},
		"pm": {"predmety"},
	}

	res := getRequest(user.Address, params)

	var predmety PredmetyResults
	xml.Unmarshal(res, &predmety)
	return predmety
}
