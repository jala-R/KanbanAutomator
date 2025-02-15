package notion

import (
	httpWrapper "jalar/notionAutomator/packages/http_wrapper"
	"jalar/notionAutomator/packages/log_wrapper"
	"jalar/notionAutomator/packages/os_wrapper"
)

type INotion interface {
	GetAllPeople() ([]NotionPeople, error)
	AddItemToBoard(dbId, title, startDate, endDate string, people NotionPeople) error
}

type notionServiceImpl struct{}

func (notionServiceImpl) GetAllPeople() ([]NotionPeople, error) {
	return getAllPeople(httpWrapper.HttpService, os_wrapper.OsService)
}

func (notionServiceImpl) AddItemToBoard(dbId, title, startDate, endDate string, people NotionPeople) error {
	return addItemToBoard(dbId, title, startDate, endDate, people, httpWrapper.HttpService, os_wrapper.OsService, log_wrapper.LoggingService)
}

var NotionService = notionServiceImpl{}
