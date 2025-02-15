package main

import (
	"fmt"
	"jalar/notionAutomator/packages/handlers"
	"net/http"
)

func main() {
	handlers.RegisterHandlers()

	fmt.Println("Server starting in 80")
	http.ListenAndServe(":80", nil)

}
