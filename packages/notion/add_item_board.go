package notion

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"jalar/notionAutomator/packages/http_wrapper"
	"jalar/notionAutomator/packages/log_wrapper"
	"jalar/notionAutomator/packages/os_wrapper"
)

func addItemToBoard(dbId, title, startDate, endDate string, people NotionPeople, httpService http_wrapper.IHttp, osService os_wrapper.IOs, logService log_wrapper.ILogging) error {

	fmt.Println(dbId, title, startDate, endDate, people)
	url := FullUrl + Pages

	//getToken
	token := osService.Getenv(TokenKey)

	//pepare request body
	bodyObj := PagesRequestBody{}
	bodyObj.Populate(dbId, title, startDate, endDate, people)

	body, err := json.Marshal(bodyObj)
	if err != nil {
		return fmt.Errorf("got error in marshalling body %w", err)
	}

	req, err := httpService.NewRequest("POST", url, bytes.NewReader(body))
	if err != nil {
		return fmt.Errorf("got error while building Request: %w", err)
	}

	addHeaders(req, token)

	//add content type
	req.Header.Add("Content-Type", "application/json")

	resp, err := httpService.Do(req)
	if err != nil {
		return fmt.Errorf("got error while sending request: %w", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("got error status code %d, with message %s", resp.StatusCode, string(body))
	}

	logService.Info(fmt.Sprintf("Added item to board title: %s", title))

	return nil
}
