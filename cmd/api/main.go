package main

import (
	"fmt"
	"jalar/notionAutomator/packages/handlers"
	"log"
	"net/http"
)

func main() {
	handlers.RegisterHandlers()

	fmt.Println("Server starting in 80")
	log.Fatal(http.ListenAndServe(":80", nil))

}
