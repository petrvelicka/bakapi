package bakapi

import (
	"encoding/xml"
	"net/url"
)

type KomDelResults struct {
	XMLName xml.Name `xml:"results"`
	Code    int      `xml:"code"`
	Message string   `xml:"message"`
	Result  int      `xml:"result"`
}

func KomDel(user User, id int) KomDelResults {
	params := url.Values{
		"hx":  {user.Token},
		"pm":  {"komdel"},
		"pmd": {string(id)},
	}

	res := getRequest(user.Address, params)

	var komdel KomDelResults
	xml.Unmarshal(res, &komdel)
	return komdel
}
