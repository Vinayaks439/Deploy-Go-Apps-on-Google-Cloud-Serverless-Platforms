package main

import (
	"log"
	"net/http"
	"os"

	"github.com/GoogleCloudPlatform/DIY-Tools/gcp-data-drive/gcpdatadrive"
)

func main() {
	// Register the initial HTTP handler.
	http.HandleFunc("/", gcpdatadrive.GetJSONData)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}

	log.Printf("Listening on port %s", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
}
