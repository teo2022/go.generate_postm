package controller

import (
	"fmt"
	"github.com/teo2022/go.generate_postm/models"
	"github.com/teo2022/go.generate_postm/utils"
	"io/ioutil"
	"log"
	"regexp"
	"strings"
)

func GetStruct(list []models.ListCatalog) []models.ListModel {
	var finSting []models.ListModel
	for _, v := range list {
		for _, r := range v.Files {
			if !strings.Contains(r, ".go") {
				continue
			}
			input, err := ioutil.ReadFile(v.Patch + "/" + r)
			if err != nil {
				log.Fatalln(err)
			}
			lines := strings.Split(string(input), "\n")
			isStrucrute := false
			isManyLine := false
			stMany := ""
			var str []string
			var name string
			var countCreate int
			for i, line := range lines {
				if strings.Contains(line, "type") {
					if strings.Contains(lines[i+1], "struct") || strings.Contains(lines[i+2], "struct") {
						isStrucrute = true
						isManyLine = true
						line = strings.Replace(line, "(", "", -1)
					}
				}

				if isManyLine == true {
					if strings.Contains(line, "//") {
						continue
					}
					line = strings.Replace(line, "\t", "", -1)
					stMany = stMany + line
				}
				if isManyLine == true && strings.Contains(line, "{") {
					isManyLine = false
					line = stMany
					stMany = ""
				}
				if isManyLine == true {
					continue
				}

				if strings.Contains(line, "type") && strings.Contains(line, "struct") {
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
						var output models.ListModel
						output.Name = name
						output.Structure = strings.Join(str, "\n")
						finSting = append(finSting, output)
						str = []string{}
					}
				}
			}
		}
	}
	finSting = GetJsonStructure(finSting)
	return finSting
}

func GetJsonStructure(list []models.ListModel) []models.ListModel {

	for i, v := range list {
		var raw string
		raw = v.Structure
		raw = strings.Replace(raw, "type", "", -1)
		raw = strings.Replace(raw, v.Name, "", -1)
		raw = strings.Replace(raw, ",omitempty", "", -1)
		lines := strings.Split(raw, "\n")
		var finJson []string
		for i, line := range lines {
			if i == 0 {
				line = strings.Replace(line, "struct", "", -1)
			}
			line = strings.ToLower(line)
			re := regexp.MustCompile("json:\"(\\w+)\"")
			match := re.FindAllStringSubmatch(line, -1)
			rawLine := ""
			isIgnore := false
			if len(match) < 1 {
				if !strings.Contains(line, "-") {
					unknow := utils.ClearArrString(strings.Split(strings.Trim(line, " "), " "))
					if len(unknow) > 1 {
						rawLine = fmt.Sprintf("    \"%v\":%v", unknow[0], GetTypeString(unknow[1]))
					} else {
						rawLine = strings.Trim(line, " ")
					}
				} else {
					isIgnore = true
				}
			}
			if len(match) >= 1 {
				gType := GetTypeString(line)
				if gType != "???" {
					rawLine = fmt.Sprintf("    \"%v\":%v", match[0][1], gType)
				} else {
					isIgnore = true
				}
			}
			if !isIgnore {
				if match != nil && match[0][1] == "id" {
					continue
				}
				finJson = append(finJson, rawLine)
			}
		}
		finJson = setComma(finJson)
		list[i].Jsonschema = strings.Join(finJson, "\n")
	}
	return list
}

func GetTypeString(line string) interface{} {
	if strings.Contains(line, "uint") {
		return 0
	}
	if strings.Contains(line, "int") {
		return 0
	}
	if strings.Contains(line, "int64") {
		return 0
	}
	if strings.Contains(line, "float64") {
		return 0.1
	}
	if strings.Contains(line, "string") {
		return "\"text\""
	}
	if strings.Contains(line, "bool") {
		return "false"
	}
	if strings.Contains(line, "time.Time") || strings.Contains(line, "time.time") {
		return "\"2022-03-14T05:40:00.000Z\""
	}
	if strings.Contains(line, "struct") {
		return "\"text\""
	}

	return "???"
}

func setComma(list []string) []string {
	for i, _ := range list {
		if i < len(list)-2 && i > 0 {
			list[i] = list[i] + ","
		}
		if strings.Contains(list[i], "\t") {
			list[i] = strings.Replace(list[i], "\t", "", -1)
		}
	}
	return list
}
