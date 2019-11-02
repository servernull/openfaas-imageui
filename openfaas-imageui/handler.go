package function

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"net/url"

	handler "github.com/openfaas-incubator/go-function-sdk"
)

//Handle incoming requests
func Handle(req handler.Request) (handler.Response, error) {
	var err error

	if len(req.QueryString) == 0 && req.Method == "GET" {
		return handler.Response{
			Body:       []byte(getHtml()),
			StatusCode: http.StatusOK,
		}, err
	}

	values, err := url.ParseQuery(req.QueryString)
	if err != nil {
		return handler.Response{
			Body:       []byte(req.QueryString),
			StatusCode: http.StatusBadRequest,
		}, err
	}

	headers := http.Header{}
	headers.Add("Content-Type", "application/json")
	action := values.Get("action")
	switch action {
	case "crawl":
		url := req.Body
		resp, err := http.Post("http://gateway.openfaas:8080/function/openfaas-imagecrawler", "application/json", bytes.NewBuffer(url))
		if err != nil {
			return handler.Response{
				Body:       []byte("{\"status\": \"error\"}"),
				StatusCode: http.StatusBadRequest,
				Header:     headers,
			}, err
		}
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		return handler.Response{
			Body:       []byte(body),
			StatusCode: http.StatusOK,
			Header:     headers,
		}, err
	case "search":
		url := req.Body
		resp, err := http.Post("http://gateway.openfaas:8080/function/openfaas-imagesearch", "application/json", bytes.NewBuffer(url))
		if err != nil {
			return handler.Response{
				Body:       []byte("{\"status\": \"error\"}"),
				StatusCode: http.StatusBadRequest,
				Header:     headers,
			}, err
		}
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		return handler.Response{
			Body:       []byte(body),
			StatusCode: http.StatusOK,
			Header:     headers,
		}, err
	case "clear":
		resp, err := http.Post("http://gateway.openfaas:8080/function/openfaas-elastic", "application/json", bytes.NewBuffer([]byte("clear")))
		if err != nil {
			return handler.Response{
				Body:       []byte("{\"status\": \"error\"}"),
				StatusCode: http.StatusBadRequest,
				Header:     headers,
			}, err
		}
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		return handler.Response{
			Body:       []byte(body),
			StatusCode: http.StatusOK,
			Header:     headers,
		}, err
	}

	return handler.Response{
		Body:       []byte("{\"status\": \"error\"}"),
		StatusCode: http.StatusBadRequest,
		Header:     headers,
	}, err
}

func getHtml() string {
	return `
	<!DOCTYPE html>
	<html lang="en">
	  <head>
		<meta charset="utf-8">
		<meta http-equiv="X-UA-Compatible" content="IE=edge">
		<meta name="viewport" content="width=device-width, initial-scale=1">
		<title>Gezellig</title>
		<link rel="stylesheet" href="https://fonts.googleapis.com/icon?family=Material+Icons">
		<link rel="stylesheet" href="https://code.getmdl.io/1.3.0/material.indigo-pink.min.css">
		<script defer src="https://code.getmdl.io/1.3.0/material.min.js"></script>
		<style>
		` + getStyle() + ` 
		</style>
	  </head>
	  <body>
		<noscript><strong>We're sorry but ui doesn't work properly without JavaScript enabled. Please enable it to continue.</strong></noscript>
		<div id=app></div>
		<script>
		` + getVendorScript() + `
		</script>	
		<script>
		` + getAppScript() + `
		</script>
	  </body>
	</html>`
}

func getVendorScript() string {
	b, _ := ioutil.ReadFile("chunk-vendors.js")
	return string(b)
}

func getAppScript() string {
	b, _ := ioutil.ReadFile("app.js")
	return string(b)
}

func getStyle() string {
	b, _ := ioutil.ReadFile("app.css")
	return string(b)
}
