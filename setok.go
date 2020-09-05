package bakapi

import (
	"encoding/xml"
	"net/url"
)

type SetOkResults struct {
	XMLName xml.Name `xml:"results"`
	Result  int      `xml:"result"`
}

func SetOk(user User, id int) SetOkResults {
	params := url.Values{
		"hx":  {user.Token},
		"pm":  {"setok"},
		"pmd": {string(id)},
	}

	res := getRequest(user.Address, params)

	var setok SetOkResults
	xml.Unmarshal(res, &setok)
	return setok
}
