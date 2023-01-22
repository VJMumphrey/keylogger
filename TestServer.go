package main

import "net/http"

func TestServer () {
	
	// used to run a server that will store the log
	http.ListenAndServe(":8090", nil)
}