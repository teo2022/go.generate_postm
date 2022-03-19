package controller

import (
	"github.com/teo2022/go.generate_postm/models"
	"github.com/teo2022/go.generate_postm/utils"
	"io/ioutil"
	"log"
	"strings"
)

//Находим роуты который принимают боду
func GetRawBody(group []models.GroupRoute, model []models.ListModel, AllFolder []models.ListCatalog) []models.GroupRoute {
	for i, v := range group {
		for i2, m := range v.Group {
			if m.MetodFunc == "POST" || m.MetodFunc == "PUT" || m.MetodFunc == "DELETE" {
				group[i].Group[i2].RawBody = GetModelFunc(m, AllFolder, model)
			}
		}
	}
	return group
}

func GetModelFunc(route models.ListRoute, AllFolder []models.ListCatalog, model []models.ListModel) string {
	for _, v := range AllFolder {
		if v.Catalog == route.Folder {
			funcMetod := getFuncFile(route.Function, v)
			modelSearch := SearchModelCodeIn(funcMetod)
			arModels := utils.ClearArrString(strings.Split(modelSearch, "."))
			if len(arModels) > 0 {
				for _, v2 := range model {
					if v2.Name == arModels[1] {
						return v2.Jsonschema
					}
				}
			}
			break
		}
	}
	return ""
}

func SearchModelCodeIn(code string) string {
	lines := strings.Split(code, "\n")
	isStrucrute := false
	var name string
	for _, line := range lines {
		if strings.Contains(line, "err := json.NewDecoder(r.Body).Decode(&") {
			isStrucrute = true
		}
		if strings.Contains(line, "err := json.NewDecoder(r.Body).Decode(") {
			isStrucrute = true
		}
		if isStrucrute == true {
			name = strings.Replace(line, "err := json.NewDecoder(r.Body).Decode(&", "", -1)
			name = strings.Replace(name, "err := json.NewDecoder(r.Body).Decode(", "", -1)
			name = strings.Replace(name, ")", "", -1)
			name = strings.Replace(name, "\t", "", -1)
			model := getModelLine(name, code)
			return model
		}
	}
	return ""
}

func getModelLine(model string, code string) string {
	lines := strings.Split(code, "\n")
	isStrucrute := false
	var name string
	for _, line := range lines {
		if strings.Contains(line, model) && strings.Contains(line, ":="){
			isStrucrute = true
		}
		if isStrucrute == true {
			name = strings.Replace(line, model,"", -1)
			name = strings.Replace(name, ":=", "", -1)
			name = strings.Replace(name, "\t", "", -1)
			name = strings.Replace(name, " ", "", -1)
			name = strings.Replace(name, "&", "", -1)
			name = strings.Replace(name, "{", "", -1)
			name = strings.Replace(name, "}", "", -1)
			return name
		}
	}
	return ""
}


func getFuncFile(search string, catalog models.ListCatalog) string {
	var finSting string
	for _, v := range catalog.Files {
		input, err := ioutil.ReadFile(catalog.Patch + "/" + v)
		if err != nil {
			log.Fatalln(err)
		}

		lines := strings.Split(string(input), "\n")

		isStrucrute := false
		var str []string
		var name  string
		var countCreate int
		for _, line := range lines {
			if strings.Contains(line, search) &&  strings.Contains(line, "w http.ResponseWriter, r *http.Request"){
				isStrucrute = true
				name = line
				name = strings.Replace(name, "type", "", -1)
				name = strings.Replace(name, "struct", "", -1)
				name = strings.Replace(name, "{", "", -1)
				name = strings.Replace(name, " ", "", -1)
			}
			if isStrucrute == true {
				str = append(str, line)
				if strings.Contains(line, "{") {
					countCreate = countCreate + 1
				}
				if strings.Contains(line, "}") {
					countCreate = countCreate - 1
				}
				if countCreate == 0 {
					isStrucrute = false
					finSting = strings.Join(str, "\n")
					break
				}
			}
		}
	}
	return finSting
}
