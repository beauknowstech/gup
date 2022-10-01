package filelist

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

// https://gosamples.dev/list-files/
func ListFilesrecursive(dirname string) {
	err := filepath.Walk(dirname,
		func(path string, _ os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			fmt.Print(path + " ")
			return nil
		})
	if err != nil {
		log.Fatal(err)
	}
}

// https://golang.cafe/blog/how-to-list-files-in-a-directory-in-go.html
func ListFiles(dirname string) {
	files, err := ioutil.ReadDir(dirname)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		fmt.Print(file.Name() + " ")
	}
}
