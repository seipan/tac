package sdk

import (
	"log"
	"net/http"
	"net/url"
	"strings"
)

const (
	host = "http://localhost:8091/add"
)

func Post(content string) {
	form := url.Values{}
	form.Add("title", content)
	req, err := http.NewRequest("POST", host, strings.NewReader(form.Encode()))
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	_, err = http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
}
