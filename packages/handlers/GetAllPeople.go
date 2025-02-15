package handlers

import (
	"encoding/json"
	"fmt"
	"jalar/notionAutomator/packages/log_wrapper"
	"jalar/notionAutomator/packages/notion"
	"net/http"
)

func getAllPeople(w http.ResponseWriter, r *http.Request, notionsService notion.INotion, logService log_wrapper.ILogging) {

	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Headers", "*")
	w.Header().Add("Allow-Control-Request-Method", "POST, GET, OPTIONS")

	peoples, err := notionsService.GetAllPeople()
	if err != nil {
		logService.Error(fmt.Sprintf("Got error at GetAllpeople %s", err.Error()))
		handleError(w, fmt.Errorf("getting all people from notion : %w", err))
		return
	}

	data, err := json.Marshal(peoples)
	if err != nil {
		logService.Error(fmt.Sprintf("Got error at GetAllpeople %s", err.Error()))
		handleError(w, fmt.Errorf("getting all people from notion : %w", err))
		return
	}

	w.Write(data)
}
