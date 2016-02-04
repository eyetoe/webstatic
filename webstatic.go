package main

import (
	"log"
	"net/http"
	// relative path to the statik bin data
	_ "./statik"
	//_ "github.com/eyetoe/webstatic/statik"
	"github.com/rakyll/statik/fs"
)

// Before buildling, run `statik -src=./data`
// to generate the statik package.
// Then, run the main program and visit http://localhost:8080/public/hello.txt
func main() {
	statikFS, err := fs.New()
	if err != nil {
		log.Fatalf(err.Error())
	}

	// Serve the data from the root
	http.Handle("/", http.FileServer(statikFS))
	// Serve the data from the ./assets/ also for example
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(statikFS)))

	http.ListenAndServe(":8080", nil)
}
