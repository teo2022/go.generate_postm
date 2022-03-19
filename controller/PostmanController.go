package controller

import (
	"encoding/json"
	"fmt"
	"github.com/teo2022/go.generate_postm/models"
	"github.com/teo2022/go.generate_postm/utils"
	"io/ioutil"
	"strings"
)

func GeneratePostman(patch string, list []models.GroupRoute, host string, port string, name string) []models.PostmanFolders {
	var postman models.Postman
	var varibles []models.PostmanVariables
	varibles = append(varibles, models.PostmanVariables{Key: "base", Value: fmt.Sprintf("%v:%v", host, port)})

	postman.Info.PostmanId = "06a8b372-5aec-4fe3-9d98-5687f5576b51"
	postman.Info.Name = name
	postman.Info.Schema = "https://schema.getpostman.com/json/collection/v2.1.0/collection.json"
	postman.Variable = varibles
	var allPostmanFolder []models.PostmanFolders
	for _, v := range list {
		var postmanFolder models.PostmanFolders
		var allRoute []models.PostmanRoute
		postmanFolder.Name = v.Name
		for _, v2 := range v.Group {
			var postmanRoute models.PostmanRoute
			postmanRoute.Name = v2.Name
			postmanRoute.Request.Method = v2.MetodFunc
			postmanRoute.Request.Url.Raw = v2.Link
			postmanRoute.Request.Url.Host = []string{"{{base}}"}
			postmanRoute.Request.Url.Path = utils.ClearArrString(strings.Split(v2.Link, "/"))
			postmanRoute.Request.Body.Raw = v2.RawBody
			postmanRoute.Request.Body.Mode = "raw"
			postmanRoute.Request.Body.Options.Raw.Language = "json"
			postmanRoute.Response = GetResponce(v2, host, port)
			allRoute = append(allRoute, postmanRoute)
		}
		postmanFolder.Item = allRoute
		allPostmanFolder = append(allPostmanFolder, postmanFolder)
	}

	postman.Item = allPostmanFolder

	file, _ := json.MarshalIndent(postman, " ", " ")
	err2 := ioutil.WriteFile(fmt.Sprintf("%v/%v.json", patch, name), file, 0644)
	if err2 != nil {
		fmt.Println(err2)
	}

	return allPostmanFolder
}

func GetResponce(v2 models.ListRoute, host string, port string) []models.PostmanResponse {
	var finResponce []models.PostmanResponse
	var postmanResponse models.PostmanResponse
	postmanResponse.Name = v2.Name
	postmanResponse.OriginalRequest.Method = v2.MetodFunc
	postmanResponse.OriginalRequest.Url.Raw = v2.Link
	postmanResponse.OriginalRequest.Url.Host = []string{host}
	postmanResponse.OriginalRequest.Url.Path = utils.ClearArrString(strings.Split(v2.Link, "/"))
	postmanResponse.OriginalRequest.Body.Raw = v2.RawBody
	postmanResponse.OriginalRequest.Body.Options.Raw.Language = "json"
	postmanResponse.OriginalRequest.Body.Mode = "raw"
	postmanResponse.PostmanPreviewlanguage = "json"
	postmanResponse.Status = "OK"
	postmanResponse.Code = 200
	finResponce = append(finResponce, postmanResponse)
	return finResponce
}
