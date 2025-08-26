package main

import (
	"flag"
	"fmt"
	"html/template"
	"io"
	"log"
	"net"
	"net/http"
	"os"
	"path/filepath"
	"strconv"

	"github.com/n000b3r/gup/internal/filelist"
	"github.com/n000b3r/gup/internal/iplist"
	"github.com/n000b3r/gup/internal/logger"
	"github.com/fatih/color"
)

func main() {
	log.SetFlags(0)

	pwd, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}

	webroot := flag.String("d", pwd, "Specify which directory")
	port := flag.Int("p", 80, "Specify which port to listen on")
	recursive := flag.Bool("r", false, "Recursive or nah?")
	flag.Parse()

	strport := ":" + strconv.Itoa(*port)
	fs := logger.LoggingHandler(http.FileServer(http.Dir(*webroot)))

	bold := color.New(color.Bold, color.Underline)
	bold.Print("Local Directory:")
	fmt.Print(" " + *webroot)
	bold.Print("\nListening on:\n")
	iplist.LocalIP()
	bold.Print("Port:")
	fmt.Print(" " + strconv.Itoa(*port))
	bold.Print("\nFiles:\n")
	if *recursive {
		filelist.ListFilesrecursive(*webroot)
	} else {
		filelist.ListFiles(*webroot)
	}
	fmt.Println()

	mux := http.NewServeMux()
	mux.Handle("/", rootHandler(*webroot, fs))

	log.Fatal(http.ListenAndServe(strport, mux))
}

func rootHandler(uploadDir string, fs http.Handler) http.Handler {
	green := color.New(color.FgGreen).SprintFunc()
	tmpl := template.Must(template.New("upload").Parse(`
<!DOCTYPE html>
<html>
<head>
    <meta charset="utf-8">
    <title>Upload & Files</title>
</head>
<body>
    <h1>Upload a file</h1>
    <form method="POST" action="/" enctype="multipart/form-data">
        <input type="file" name="file" />
        <input type="submit" value="Upload" />
    </form>
</body>
</html>
`))

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/" {
			switch r.Method {
			case http.MethodGet:
				tmpl.Execute(w, nil)
				return

			case http.MethodPost:
				r.ParseMultipartForm(10 << 20)
				f, hdr, err := r.FormFile("file")
				if err != nil {
					http.Error(w, "Failed to read upload: "+err.Error(), http.StatusBadRequest)
					return
				}
				defer f.Close()

				dstPath := filepath.Join(uploadDir, hdr.Filename)
				dst, err := os.Create(dstPath)
				if err != nil {
					http.Error(w, "Failed to save file: "+err.Error(), http.StatusInternalServerError)
					return
				}
				defer dst.Close()

				if _, err := io.Copy(dst, f); err != nil {
					http.Error(w, "Failed to save file: "+err.Error(), http.StatusInternalServerError)
					return
				}

				host, _, err := net.SplitHostPort(r.RemoteAddr)
				if err != nil {
					host = r.RemoteAddr
				}
				fmt.Println("\n" + host + " Received: " + green(hdr.Filename))
				fmt.Fprintf(w, "\n%s Received: %s", host, green(hdr.Filename))
				return

			default:
				http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
				return
			}
		}
		fs.ServeHTTP(w, r)
	})
}

