package bakapi

import (
	"crypto/sha512"
	"encoding/base64"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
)

func hashString(str string) string {
	hasher := sha512.New()
	hasher.Write([]byte(str))
	return base64.URLEncoding.EncodeToString(hasher.Sum(nil))
}

func getRequest(address string, params url.Values) []byte {
	req, err := http.NewRequest("GET", address, nil)
	if err != nil {
		log.Print(err)
		os.Exit(1)
	}

	req.URL.RawQuery = params.Encode()

	resp, err := http.Get(req.URL.String())
	if err != nil {
		log.Print(err)
		os.Exit(1)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Print(err)
		os.Exit(1)
	}

	return body
}
