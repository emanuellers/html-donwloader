package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"regexp"
	"strings"
)

const (
	Hiperlinks   = `"(http://|https://).*?"` //
	InsideQuotes = `"(.*?)"`
)

type Html struct{}

func (r *Html) GetBufferResponse(url string) ([]byte, error) {
	request, err := http.Get(url)

	if err != nil {
		panic(err)
	}

	defer request.Body.Close()

	buffer, err := ioutil.ReadAll(request.Body)
	if err != nil {
		panic(err)
	}

	return buffer, nil
}
func (r Html) GetPageHtmlToFile(url string, filename string) error {
	buffer, err := r.GetBufferResponse(url)
	if err != nil {
		panic(err)
	}
	file, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	_, err = file.Write(buffer)
	defer file.Close()

	return err
}

func (r *Html) GetHiperlinks(buffer []byte, max int) []string {
	expression := regexp.MustCompile(Hiperlinks)
	matches := expression.FindAllString(string(buffer), len(buffer))
	return matches
}

func (r *Html) CheckFuntionalUrls(url string) int {
	request, err := http.Get(strings.ReplaceAll(url, "\"", ""))
	if err != nil {
		log.Panic(err)
	}

	return request.StatusCode

}
