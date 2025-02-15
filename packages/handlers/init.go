package handlers

import (
	"jalar/notionAutomator/packages/log_wrapper"
	"jalar/notionAutomator/packages/notion"
	"net/http"
)

func RegisterHandlers() {

	http.HandleFunc("/GetPeople", func(w http.ResponseWriter, r *http.Request) {
		getAllPeople(w, r, notion.NotionService, log_wrapper.LoggingService)
	})

	http.HandleFunc("/SpritualProject", func(w http.ResponseWriter, r *http.Request) {
		spritualProject(w, r, notion.NotionService)
	})

	http.HandleFunc("/ManhwaCopyProject", func(w http.ResponseWriter, r *http.Request) {
		manhwaCopyProject(w, r, notion.NotionService)
	})

	http.HandleFunc("/ManhwaOrgProject", func(w http.ResponseWriter, r *http.Request) {
		manhwaOrgProject(w, r, notion.NotionService)
	})

}
