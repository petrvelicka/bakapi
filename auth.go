package bakapi

import (
	"encoding/xml"
	"net/url"
	"strings"
	"time"
)

type HxResults struct {
	XMLName xml.Name `xml:"results"`
	Res     int      `xml:"res"`
	Typ     string   `xml:"typ"`
	Ikod    string   `xml:"ikod"`
	Salt    string   `xml:"salt"`
}

func GetHashedPwd(login, pwd, address string) string {
	getHx := url.Values{
		"gethx": {login},
	}
	res := getRequest(address, getHx)

	var hxResults HxResults
	xml.Unmarshal(res, &hxResults)

	pwdSalt := hxResults.Salt + hxResults.Ikod + hxResults.Typ + pwd
	pwdHashed := hashString(pwdSalt)
	pwdHashed = strings.Replace(pwdHashed, "_", "/", -1)
	return pwdHashed
}

func GetToken(login, hashedpwd string) string {
	tokenSalt := "*login*" + login + "*pwd*" + hashedpwd + "*sgn*ANDR" + time.Now().Format("20060102")
	token := hashString(tokenSalt)
	return token
}

func Authentificate(login, hashedpwd, address string) User {
	token := GetToken(login, hashedpwd)

	return User{token, address}
}
