package bakapi

import (
	"encoding/xml"
	"net/url"
)

type Login struct {
	XMLName xml.Name `xml:"results"`
	Verze   string   `xml:"verze"`
	Jmeno   string   `xml:"jmeno"`
	Typ     string   `xml:"typ"`
	Strtyp  string   `xml:"strtyp"`
	Skola   string   `xml:"skola"`
	Trida   string   `xml:"trida"`
	Rocnik  string   `xml:"rocnik"`
	Moduly  string   `xml:"moduly"`
	Params  Params   `xml:"params"`
	Result  int      `xml:"result"`
}

type Params struct {
	XMLName     xml.Name `xml:"params"`
	NewMarkDays int      `xml:"newmarkdays"`
}

func GetLogin(user User) Login {
	params := url.Values{
		"hx": {user.Token},
		"pm": {"login"},
	}

	res := getRequest(user.Address, params)

	var login Login
	xml.Unmarshal(res, &login)
	return login
}
