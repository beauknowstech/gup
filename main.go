package main

import (
	"flag"
	"fmt"
	"gup/internal/filelist"
	"gup/internal/iplist"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/fatih/color"
)

func main() {
	pwd, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}

	webroot := flag.String("d", pwd, "Specify which directory")
	port := flag.Int("p", 80, "Specify which port to listen on")
	recursive := flag.Bool("r", false, "Recursive or nah?")
	flag.Parse()

	// create file server handler
	fs := http.FileServer(http.Dir(*webroot))
	//if err != nil {
	//	fmt.Println(err)
	//}
	strport := ":" + (strconv.Itoa(*port))

	bold := color.New(color.Bold, color.Underline)
	//print info
	bold.Print("Local Directory:")
	fmt.Print(" " + *webroot)
	// Print local IP's
	bold.Print("\nListening on:\n")
	iplist.LocalIP()
	bold.Print("Port:")
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
