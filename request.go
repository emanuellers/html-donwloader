package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
)

const (
	HrefAttribute = `href="(.*?)".*?`
)

type Request struct{}

func (r *Request) GetBufferResponse(url string) ([]byte, error) {
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
func (r Request) GetPageHtmlToFile(url string, filename string) error {
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

func (r *Request) GetHiperlinks(buffer []byte, max int) []string {
	expression := regexp.MustCompile(HrefAttribute)
	matches := expression.FindAllString(string(buffer), len(buffer))

	//for _, match := range matches {
	//	match = match[4:]
	//	fmt.Println(match[4:])
	//}

	fmt.Println(matches)
	return nil
}
