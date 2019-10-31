package function

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"net/url"

	handler "github.com/openfaas-incubator/go-function-sdk"
)

// Handle a function invocation
func Handle(req handler.Request) (handler.Response, error) {
	var err error

	if len(req.QueryString) == 0 {
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
		<link rel="stylesheet" href="https://stackpath.bootstrapcdn.com/bootstrap/4.3.1/css/bootstrap.min.css" integrity="sha384-ggOyR0iXCbMQv3Xipma34MD+dH/1fQ784/j6cY/iJTQUOhcWr7x9JvoRxT2MZw1T" crossorigin="anonymous">
		<style>
		` + getStyle() + ` 
		</style>
	  </head>
	  <body>
		<div class="site-wrapper">
	
		  <div class="site-wrapper-inner">
	
			<div class="cover-container">
	
			  <div class="masthead clearfix">
				<div class="inner">
				  <h3 class="masthead-brand">Gezellig</h3>
				  <nav>
					<ul class="nav masthead-nav">
					  <li class="active"><label for="site"><input id="site"></li>
					  <li><button id="crawl">crawl</button></li>
					  <li><button id="search">search</button></li>
					</ul>
				  </nav>
				</div>
			  </div>
	
			  <div id="content" class="inner cover">
				<h1 class="cover-heading">An experiment: Crawl and Analyze images.</h1>
				<p class="lead">Gezellig is a pure FaaS solution that crawls a site for images and indexes information about them for searching.</p>
				<p class="lead">
				  <a target="_blank" href="https://github.com/servernull/gezellig" class="btn btn-lg btn-default">Learn more</a>
				</p>
			  </div>
	
			  <div class="mastfoot">
				<div class="inner">
				  <p>A <a target="_blank" href="https://github.com/servernull">Server Null</a> project.</p>
				</div>
			  </div>
			</div>
		  </div>
		</div>
	
	
		<script src="https://ajax.googleapis.com/ajax/libs/jquery/1.12.4/jquery.min.js"></script>
		<script src="https://maxcdn.bootstrapcdn.com/bootstrap/3.3.7/js/bootstrap.min.js" integrity="sha384-Tc5IQib027qvyjSMfHjOMaLkfuWVxZxUPnCJA7l2mCWNIpG9mGCD8wGNIcPD7Txa" crossorigin="anonymous"></script>
	
		<script>
		` + getScript() + `
		</script>
	  </body>
	</html>`
}

func getScript() string {
	b, _ := ioutil.ReadFile("script.js")
	return string(b)
}

func getStyle() string {
	b, _ := ioutil.ReadFile("style.css")
	return string(b)
}
