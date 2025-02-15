package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"jalar/notionAutomator/packages/notion"
	"net/http"
)

type SpritualProjectRequest struct {
	Projectname    string   `json:"project_name,omitempty"`
	Script         WorkItem `json:"script,omitempty"`
	TitleThumbnail WorkItem `json:"title_thumbnail,omitempty"`
	Edit           WorkItem `json:"edit,omitempty"`
	Schedule       WorkItem `json:"schedule,omitempty"`
}

type WorkItem struct {
	Assign    notion.NotionPeople `json:"assign"`
	StartDate string              `json:"start_date"`
	EndDate   string              `json:"end_date"`
}

const DBId = "17ca844dadf480c3ad41defece223838"

func spritualProject(w http.ResponseWriter, r *http.Request, notionService notion.INotion) {
	defer r.Body.Close()

	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Headers", "*")
	w.Header().Add("Allow-Control-Request-Method", "POST, GET, OPTIONS")

	data, err := io.ReadAll(r.Body)
	var reqBody SpritualProjectRequest
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(data))

	err = json.Unmarshal(data, &reqBody)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(reqBody)
	//project name
	//script
	err = notionService.AddItemToBoard(DBId, reqBody.Projectname+"-Script", reqBody.Script.StartDate, reqBody.Script.EndDate, reqBody.Script.Assign)
	if err != nil {
		fmt.Println(err)
		return
	}
	//title thumbnail
	err = notionService.AddItemToBoard(DBId, reqBody.Projectname+"-Title & Thumbnail", reqBody.TitleThumbnail.StartDate, reqBody.TitleThumbnail.EndDate, reqBody.TitleThumbnail.Assign)
	if err != nil {
		fmt.Println(err)
		return
	}
	//edit
	err = notionService.AddItemToBoard(DBId, reqBody.Projectname+"-Edit", reqBody.Edit.StartDate, reqBody.Edit.EndDate, reqBody.Edit.Assign)
	if err != nil {
		fmt.Println(err)
		return
	}

	//schedule
	err = notionService.AddItemToBoard(DBId, reqBody.Projectname+"-Schedule", reqBody.Schedule.StartDate, reqBody.Schedule.EndDate, reqBody.Schedule.Assign)
	if err != nil {
		fmt.Println(err)
		return
	}
}

type ManhwaCopyProjectRequest struct {
	Projectname     string   `json:"project_name"`
	ImagesFromVideo WorkItem `json:"images_from_video"`
	BucketImages    WorkItem `json:"bucket_images"`
	OrgScript       WorkItem `json:"org_script"`
	RewriteScript   WorkItem `json:"rewrite_script"`
	ImageScreenshot WorkItem `json:"image_screenshot"`
	Script          WorkItem `json:"script"`
	TitleThumbnail  WorkItem `json:"title_thumbnail"`
	Edit            WorkItem `json:"edit"`
	Schedule        WorkItem `json:"schedule"`
}

func manhwaCopyProject(w http.ResponseWriter, r *http.Request, notionService notion.INotion) {
	defer r.Body.Close()

	data, _ := io.ReadAll(r.Body)
	var reqBody ManhwaCopyProjectRequest

	json.Unmarshal(data, &reqBody)

	//get images from video
	go notionService.AddItemToBoard(DBId, reqBody.Projectname+"- Images From Video", reqBody.ImagesFromVideo.StartDate, reqBody.ImagesFromVideo.EndDate, reqBody.ImagesFromVideo.Assign)

	//bucket images
	go notionService.AddItemToBoard(DBId, reqBody.Projectname+"- Bucket Images", reqBody.BucketImages.StartDate, reqBody.BucketImages.EndDate, reqBody.BucketImages.Assign)

	//get org script
	go notionService.AddItemToBoard(DBId, reqBody.Projectname+"- Org Script", reqBody.OrgScript.StartDate, reqBody.OrgScript.EndDate, reqBody.OrgScript.Assign)

	//rewrite script
	go notionService.AddItemToBoard(DBId, reqBody.Projectname+"-Rewrite Script", reqBody.RewriteScript.StartDate, reqBody.RewriteScript.EndDate, reqBody.RewriteScript.Assign)

	//image screenshots
	go notionService.AddItemToBoard(DBId, reqBody.Projectname+"-Image ScreenShots", reqBody.ImageScreenshot.StartDate, reqBody.ImageScreenshot.EndDate, reqBody.ImageScreenshot.Assign)

	//edit
	go notionService.AddItemToBoard(DBId, reqBody.Projectname+"-Edit", reqBody.Edit.StartDate, reqBody.Edit.EndDate, reqBody.Edit.Assign)

	//thumbnail & title
	go notionService.AddItemToBoard(DBId, reqBody.Projectname+"-Thumbnail & title", reqBody.TitleThumbnail.StartDate, reqBody.TitleThumbnail.EndDate, reqBody.TitleThumbnail.Assign)

	//schedule
	go notionService.AddItemToBoard(DBId, reqBody.Projectname+"-Schedule", reqBody.Schedule.StartDate, reqBody.Schedule.EndDate, reqBody.Schedule.Assign)

}

type ManhwaOrgProjectRequest struct {
	Projectname       string   `json:"project_name"`
	ImageWithWords    WorkItem `json:"image_with_words"`
	ImageWithoutWords WorkItem `json:"image_without_words"`
	Script            WorkItem `json:"script"`
	TitleThumbnail    WorkItem `json:"title_thumbnail"`
	Edit              WorkItem `json:"edit"`
	Schedule          WorkItem `json:"schedule"`
}

func manhwaOrgProject(w http.ResponseWriter, r *http.Request, notionService notion.INotion) {
	defer r.Body.Close()

	data, _ := io.ReadAll(r.Body)
	var reqBody ManhwaOrgProjectRequest

	json.Unmarshal(data, &reqBody)

	//get images with words
	go notionService.AddItemToBoard(DBId, reqBody.Projectname+"-Images with words", reqBody.ImageWithWords.StartDate, reqBody.ImageWithWords.EndDate, reqBody.ImageWithWords.Assign)

	//images without words
	go notionService.AddItemToBoard(DBId, reqBody.Projectname+"-Images without words", reqBody.ImageWithoutWords.StartDate, reqBody.ImageWithoutWords.EndDate, reqBody.ImageWithoutWords.Assign)

	//rewrite script
	go notionService.AddItemToBoard(DBId, reqBody.Projectname+"-Script", reqBody.Script.StartDate, reqBody.Script.EndDate, reqBody.Script.Assign)

	//edit
	go notionService.AddItemToBoard(DBId, reqBody.Projectname+"-Edit", reqBody.Edit.StartDate, reqBody.Edit.EndDate, reqBody.Edit.Assign)

	//thumbnail & title
	go notionService.AddItemToBoard(DBId, reqBody.Projectname+"-Thumbnail & title", reqBody.TitleThumbnail.StartDate, reqBody.TitleThumbnail.EndDate, reqBody.TitleThumbnail.Assign)

	//schedule
	go notionService.AddItemToBoard(DBId, reqBody.Projectname+"-Schedule", reqBody.Schedule.StartDate, reqBody.Schedule.EndDate, reqBody.Schedule.Assign)

}
