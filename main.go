package main

import (
	"flag"
	"fmt"
	"gup/internal/filelist"
	"gup/internal/iplist"
	"net/http"
	"strconv"

	"github.com/fatih/color"
)

func main() {

	webroot := flag.String("d", ".", "Specify which directory")
	port := flag.Int("p", 80, "Specify which port to listen on")
	recursive := flag.Bool("r", false, "Recursive or nah?")
	flag.Parse()

	// create file server handler
	fs := http.FileServer(http.Dir(*webroot))
	//path, err := os.Getwd()
	//if err != nil {
	//	fmt.Println(err)
	//}
	strport := ":" + (strconv.Itoa(*port))

	bold := color.New(color.Bold, color.Underline)
	//print info
	//bold.Printf("Directory: %s\nPort: %d\n", *webroot, *port)
	bold.Print("Local Directory:")
	fmt.Print(" " + *webroot)
	// Print local IP's
	bold.Print("\nListening on:\n")
	iplist.LocalIP()
	bold.Print("\nPort:")
	fmt.Print(" " + (strconv.Itoa(*port)))

	// print list of files

	if *recursive {
		bold.Print("\nFiles:\n")
		filelist.ListFilesrecursive(*webroot)
	} else {
		bold.Print("\nFiles:\n")
		filelist.ListFiles(*webroot)
	}
	// start HTTP server with `fs` as the default handler
	fmt.Println(http.ListenAndServe(strport, fs))

}
