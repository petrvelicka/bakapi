package bakapi

import (
	"encoding/xml"
	"net/url"
)

type Interfaces struct {
	XMLName   xml.Name    `xml:"interfaces"`
	Interface []Interface `xml:"interface"`
	Result    int         `xml:"result"`
}

type Interface struct {
	XMLName      xml.Name `xml:"interface"`
	Version      int      `xml:"version"`
	Deprecated   bool     `xml:"deprecated"`
	Experimental bool     `xml:"experimental"`
}

func GetInterfaces(user User) Interfaces {
	params := url.Values{
		"hx": {user.Token},
		"pm": {"interfaces"},
	}

	res := getRequest(user.Address, params)

	var interfaces Interfaces
	xml.Unmarshal(res, &interfaces)
	return interfaces
}
