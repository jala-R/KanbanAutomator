package notion

import (
	"encoding/json"
	"fmt"
	"io"
	httpWrapper "jalar/notionAutomator/packages/http_wrapper"
	"jalar/notionAutomator/packages/os_wrapper"
	"net/http"
)

func getAllPeople(httpService httpWrapper.IHttp, osService os_wrapper.IOs) ([]NotionPeople, error) {

	//get url
	url := FullUrl + UsersRoute

	//get token
	token := osService.Getenv(TokenKey)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		err = fmt.Errorf("building request: %w", err)
		return nil, err
	}

	addHeaders(req, token)

	res, err := httpService.Do(req)
	if err != nil {
		err = fmt.Errorf("got http error: %w", err)
		return nil, err
	}
	if res.StatusCode < 200 || res.StatusCode >= 300 {
		body, _ := io.ReadAll(res.Body)
		return nil, fmt.Errorf("got Error status code %d with message %s", res.StatusCode, body)
	}
	defer res.Body.Close()

	body, _ := io.ReadAll(res.Body)

	var unmarshalledResp GetPeopleResponse
	json.Unmarshal(body, &unmarshalledResp)

	return unmarshalledResp.Results, nil

}
