package main

import (
	"fmt"
	"log"
	"net/http"
	"rest-api/controller"
)

func main() {
	http.HandleFunc("/ConnectionSocialNetworks/facebook/login", controller.Login)
	http.HandleFunc("/ConnectionSocialNetworks/facebook/redirect", controller.Redirect)
	http.HandleFunc("/ConnectionSocialNetworks/facebook/error", controller.Error)
	http.HandleFunc("/ConnectionSocialNetworks/facebook/media", controller.Media)
	fmt.Print("Started running on http://localhost:7102\n")
	log.Fatal(http.ListenAndServeTLS(":7102", "certs/cert.pem", "certs/key.pem", nil))
}
