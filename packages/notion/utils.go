package notion

import (
	"net/http"
)

const (
	FullUrl    = "https://api.notion.com/v1"
	UsersRoute = "/users"
	Pages      = "/pages"
	TokenKey   = "NOTIONTOKEN"
)

func addHeaders(req *http.Request, token string) {
	req.Header.Add("Notion-Version", "2022-06-28")

	//add bearer token
	token = "Bearer " + token
	req.Header.Add("Authorization", token)
}
