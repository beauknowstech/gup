package main

import (
	"fmt"
	//"log"
	"net/http"
	"os"
	"strconv"
)

var (
	port    = 80
	ip      = "0.0.0.0"
	webroot = "."
)

func main() {

	// create file server handler
	fs := http.FileServer(http.Dir(webroot))
	path, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}
	strport := ":" + (strconv.Itoa(port))

	// start HTTP server with `fs` as the default handler
	fmt.Printf("Serving %s on HTTP port:%d\n", path, port)
	fmt.Println(http.ListenAndServe(strport, fs))

}
