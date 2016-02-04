package main

import (
	"fmt"
	"github.com/rakyll/statik/fs"
	"log"
	"net/http"
	"os/exec"
	"runtime"
	// initialize binary data in statik dir
	_ "github.com/eyetoe/webstatic/statik"
)

// Before buildling, run `statik -src=./data`
// to generate the statik package.
// Note! when updating contents of ./data delete the ./statik dir so it can be recreated
func main() {
	statikFS, err := fs.New()
	if err != nil {
		log.Fatalf(err.Error())
	}

	// Serve the data from the root
	http.Handle("/", http.FileServer(statikFS))
	// Serve the data from the ./assets/ here as an example
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(statikFS)))

	openURL("http://localhost:8080/")
	http.ListenAndServe(":8080", nil)
}

func openURL(url string) error {
	var err error
	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", url).Start()
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", "http://localhost:4001/").Start()
	case "darwin":
		err = exec.Command("open", url).Start()
	default:
		err = fmt.Errorf("Cannot open URL %s on this platform", url)
	}
	return err
}
