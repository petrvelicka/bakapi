package bakapi

import (
	"encoding/xml"
	"net/url"
)

type AbsenceResult struct {
	XMLName xml.Name `xml:"results"`
	Absence Absence  `xml:"absence"`
	Result  int      `xml:"result"`
}

type Absence struct {
	XMLName     xml.Name    `xml:"absence"`
	Hranice     string      `xml:"hranice"`
	Seznam      Seznam      `xml:"seznam"`
	Zameskanost Zameskanost `xml:"zameskanost"`
}

type Seznam struct {
	XMLName xml.Name `xml:"seznam"`
	Abs     []Abs    `xml:"abs"`
}

type Abs struct {
	XMLName xml.Name `xml:"abs"`
	Datum   string   `xml:"datum"`
	Den     string   `xml:"den"`
	A       int      `xml:"A"`
	AOk     int      `xml:"AOk"`
	AMiss   int      `xml:"AMiss"`
	ALate   int      `xml:"ALate"`
	ASoon   int      `xml:"ASoon"`
	ASchool int      `xml:"ASchool"`
}

type Zameskanost struct {
	XMLName xml.Name  `xml:"zameskanost"`
	Nadpis  string    `xml:"nadpis"`
	Predmet []Predmet `xml:"predmet"`
}

type Predmet struct {
	XMLName   xml.Name `xml:"predmet"`
	Nazev     string   `xml:"nazev"`
	Oduceno   int      `xml:"oduceno"`
	Absbase   int      `xml:"absbase"`
	Absschool int      `xml:"absschool"`
	Abssoon   int      `xml:"abssoon"`
	Abslate   int      `xml:"abslate"`
}

func GetAbsence(user User) AbsenceResult {
	params := url.Values{
		"hx": {user.Token},
		"pm": {"absence"},
	}

	res := getRequest(user.Address, params)

	var absenceResult AbsenceResult
	xml.Unmarshal(res, &absenceResult)
	return absenceResult
}
