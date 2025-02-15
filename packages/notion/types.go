package notion

import "github.com/google/uuid"

type PagesRequestBody struct {
	Parent     Parent       `json:"parent"`
	Properties PageProperty `json:"properties"`
}

func (body *PagesRequestBody) Populate(dbId, title, start, end string, people NotionPeople) {
	body.Parent.Populate(dbId)
	body.Properties.Populate(start, end, title, people)
}

type Parent struct {
	DbID string `json:"database_id"`
}

func (parent *Parent) Populate(dbId string) {
	parent.DbID = dbId
}

type PageProperty struct {
	Name   Name   `json:"Name"`
	Assign Assign `json:"Assign"`
	Date   Date   `json:"Date"`
}

func (p *PageProperty) Populate(start, end, title string, people NotionPeople) {
	p.Name.Populate(title)
	p.Assign.Populate(people)
	p.Date.Populate(start, end)
}

type NotionObject struct {
	Id   string `json:"id"`
	Type string `json:"type"`
}

type Titleobj struct {
	Type string            `json:"type"`
	Text map[string]string `json:"text"`
}

type Name struct {
	NotionObject
	Title []Titleobj `json:"title"`
}

func (p *Name) Populate(title string) {
	p.Id = uuid.NewString()
	p.Type = "title"

	p.Title = append(p.Title, Titleobj{
		"text",
		map[string]string{
			"content": title,
		},
	})
}

type Assign struct {
	NotionObject
	People []NotionPeople `json:"people"`
}

func (p *Assign) Populate(people NotionPeople) {
	p.Id = uuid.NewString()
	p.Type = "people"
	p.People = append(p.People, people)
}

type Date struct {
	NotionObject
	DateDetails DateDetails `json:"date"`
}

func (p *Date) Populate(start, end string) {
	p.Id = uuid.NewString()
	p.Type = "date"

	p.DateDetails = DateDetails{
		Start: start,
		End:   end,
	}
}

type DateDetails struct {
	Start    string `json:"start"`
	End      string `json:"end"`
	TimeZone any    `json:"time_zone"`
}

type GetPeopleResponse struct {
	Results []NotionPeople `json:"results"`
}

type NotionPeople struct {
	Object    string       `json:"object"`
	Id        string       `json:"id"`
	Name      string       `json:"name"`
	AvatarUrl string       `json:"avatar_url"`
	Type      string       `json:"type"`
	Person    NotionPerson `json:"person"`
}

type NotionPerson struct {
	Email string `json:"email"`
}
