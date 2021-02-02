package main

import (
	"net/http"
	
	"github.com/shortdaddy0711/goWeb/myapp"
)

func main() {

	http.ListenAndServe(":3000", myapp.NewHTTPHandler())
}
