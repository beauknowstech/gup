package main

import (
	"flag"
	"fmt"
	"gup/internal/filelist"
	"net/http"
	"strconv"
)

func main() {

	recursive := flag.Bool("r", false, "Recursive or nah?")
	webroot := flag.String("d", ".", "Specify which directory")
	port := flag.Int("p", 80, "Specify which port to listen on")
	flag.Parse()

	// create file server handler
	fs := http.FileServer(http.Dir(*webroot))
	//path, err := os.Getwd()
	//if err != nil {
	//	fmt.Println(err)
	//}
	strport := ":" + (strconv.Itoa(*port))

	// start HTTP server with `fs` as the default handler
	fmt.Printf("Directory: %s\nPort: %d\n", *webroot, *port)
	// print list of files

	if *recursive {
		fmt.Print("Files:\n")
		filelist.ListFilesrecursive(*webroot)
	} else {
		fmt.Print("Files:\n")
		filelist.ListFiles(*webroot)
	}

	fmt.Println(http.ListenAndServe(strport, fs))

}
