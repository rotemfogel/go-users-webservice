package main

import (
	"log"
	"me.rotemfo/webservice/controllers"
	"net/http"
	"strconv"
)

func main() {
	port := 3000
	_, err := startWebServer(&port)
	if err != nil {
		log.Fatal(err)
	}
	log.Print("Web server started on port", port)
}

func startWebServer(port *int) (*int, error) {
	controllers.RegisterControllers()
	log.Println("Start web server on port", *port)
	sPort := ":" + strconv.Itoa(*port)
	return port, http.ListenAndServe(sPort, nil)
}
