package function

import (
	"net/http"
	"net/url"

	handler "github.com/openfaas-incubator/go-function-sdk"
)

const index = `
<!DOCTYPE html>
<html lang="en">
  <head>
    <meta charset="utf-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width, initial-scale=1">
    <!-- The above 3 meta tags *must* come first in the head; any other head content must come *after* these tags -->
    <title>Bootstrap 101 Template</title>

    <!-- Bootstrap -->
	<link rel="stylesheet" href="https://stackpath.bootstrapcdn.com/bootstrap/4.3.1/css/bootstrap.min.css" integrity="sha384-ggOyR0iXCbMQv3Xipma34MD+dH/1fQ784/j6cY/iJTQUOhcWr7x9JvoRxT2MZw1T" crossorigin="anonymous">

    <!-- HTML5 shim and Respond.js for IE8 support of HTML5 elements and media queries -->
    <!-- WARNING: Respond.js doesn't work if you view the page via file:// -->
    <!--[if lt IE 9]>
      <script src="https://oss.maxcdn.com/html5shiv/3.7.3/html5shiv.min.js"></script>
      <script src="https://oss.maxcdn.com/respond/1.4.2/respond.min.js"></script>
	<![endif]-->
	<style>
	/*
	* Globals
	*/
   
   /* Links */
   a,
   a:focus,
   a:hover {
	 color: #fff;
   }
   
   /* Custom default button */
   .btn-default,
   .btn-default:hover,
   .btn-default:focus {
	 color: #333;
	 text-shadow: none; 
	 background-color: #fff;
	 border: 1px solid #fff;
   }
   
   
   /*
	* Base structure
	*/
   
   html,
   body {
	 height: 100%;
	 background-color: #333;
   }
   body {
	 color: #fff;
	 text-align: center;
	 text-shadow: 0 1px 3px rgba(0,0,0,.5);
   }
   
   /* Extra markup and styles for table-esque vertical and horizontal centering */
   .site-wrapper {
	 display: table;
	 width: 100%;
	 height: 100%; /* For at least Firefox */
	 min-height: 100%;
	 -webkit-box-shadow: inset 0 0 100px rgba(0,0,0,.5);
			 box-shadow: inset 0 0 100px rgba(0,0,0,.5);
   }
   .site-wrapper-inner {
	 display: table-cell;
	 vertical-align: top;
   }
   .cover-container {
	 margin-right: auto;
	 margin-left: auto;
   }
   
   /* Padding for spacing */
   .inner {
	 padding: 30px;
   }
   
   
   /*
	* Header
	*/
   .masthead-brand {
	 margin-top: 10px;
	 margin-bottom: 10px;
   }
   
   .masthead-nav > li {
	 display: inline-block;
   }
   .masthead-nav > li + li {
	 margin-left: 20px;
   }
   .masthead-nav > li > a {
	 padding-right: 0;
	 padding-left: 0;
	 font-size: 16px;
	 font-weight: bold;
	 color: #fff; /* IE8 proofing */
	 color: rgba(255,255,255,.75);
	 border-bottom: 2px solid transparent;
   }
   .masthead-nav > li > a:hover,
   .masthead-nav > li > a:focus {
	 background-color: transparent;
	 border-bottom-color: #a9a9a9;
	 border-bottom-color: rgba(255,255,255,.25);
   }
   .masthead-nav > .active > a,
   .masthead-nav > .active > a:hover,
   .masthead-nav > .active > a:focus {
	 color: #fff;
	 border-bottom-color: #fff;
   }
   
   @media (min-width: 768px) {
	 .masthead-brand {
	   float: left;
	 }
	 .masthead-nav {
	   float: right;
	 }
   }
   
   
   /*
	* Cover
	*/
   
   .cover {
	 padding: 0 20px;
   }
   .cover .btn-lg {
	 padding: 10px 20px;
	 font-weight: bold;
   }
   
   
   /*
	* Footer
	*/
   
   .mastfoot {
	 color: #999; /* IE8 proofing */
	 color: rgba(255,255,255,.5);
   }
   
   
   /*
	* Affix and center
	*/
   
   @media (min-width: 768px) {
	 /* Pull out the header and footer */
	 .masthead {
	   position: fixed;
	   top: 0;
	 }
	 .mastfoot {
	   position: fixed;
	   bottom: 0;
	 }
	 /* Start the vertical centering */
	 .site-wrapper-inner {
	   vertical-align: middle;
	 }
	 /* Handle the widths */
	 .masthead,
	 .mastfoot,
	 .cover-container {
	   width: 100%; /* Must be percentage or pixels for horizontal alignment */
	 }
   }
   
   @media (min-width: 992px) {
	 .masthead,
	 .mastfoot,
	 .cover-container {
	   width: 700px;
	 }
   }
   
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

          <div class="inner cover">
            <h1 class="cover-heading">Cover your page.</h1>
            <p class="lead">Cover is a one-page template for building simple and beautiful home pages. Download, edit the text, and add your own fullscreen background photo to make it your own.</p>
            <p class="lead">
              <a href="#" class="btn btn-lg btn-default">Learn more</a>
            </p>
          </div>

          <div class="mastfoot">
            <div class="inner">
              <p>Cover template for <a href="http://getbootstrap.com">Bootstrap</a>, by <a href="https://twitter.com/mdo">@mdo</a>.</p>
            </div>
          </div>

        </div>

      </div>

    </div>


    <!-- jQuery (necessary for Bootstrap's JavaScript plugins) -->
    <script src="https://ajax.googleapis.com/ajax/libs/jquery/1.12.4/jquery.min.js"></script>
    <!-- Include all compiled plugins (below), or include individual files as needed -->
	<script src="https://maxcdn.bootstrapcdn.com/bootstrap/3.3.7/js/bootstrap.min.js" integrity="sha384-Tc5IQib027qvyjSMfHjOMaLkfuWVxZxUPnCJA7l2mCWNIpG9mGCD8wGNIcPD7Txa" crossorigin="anonymous"></script>

	<script>
		document.getElementById("crawl").addEventListener("click", function(){
			fetch('?action=crawl', {
				method: 'post',
				body: document.getElementById("site").value
			  }).then(function(response) {
				return response.json();
		  	  }).then(function(data) {
				//todo handle json
			});
		});
		document.getElementById("search").addEventListener("click", function(){
			fetch('?action=search', {
				method: 'post',
				body: document.getElementById("site").value
			  }).then(function(response) {
				return response.json();
		  	  }).then(function(data) {
				  //todo handle json
			});
	  	});
	</script>
  </body>
</html>`

// Handle a function invocation
func Handle(req handler.Request) (handler.Response, error) {
	var err error

	if len(req.QueryString) == 0 {
		return handler.Response{
			Body:       []byte(index),
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

	action := values.Get("action")
	switch action {
	case "crawl":
		headers := http.Header{}
		headers.Add("Content-Type", "application/json")
		return handler.Response{
			Body:       []byte("{\"status\": \"crawler\"}"),
			StatusCode: http.StatusOK,
			Header:     headers,
		}, err
	case "search":
		headers := http.Header{}
		headers.Add("Content-Type", "application/json")
		return handler.Response{
			Body:       []byte("{\"status\": \"search\"}"),
			StatusCode: http.StatusOK,
			Header:     headers,
		}, err
	}

	return handler.Response{
		Body:       []byte(""),
		StatusCode: http.StatusBadRequest,
	}, err
}
