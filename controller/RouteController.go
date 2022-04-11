package controller

import (
	"github.com/teo2022/go.generate_postm/models"
	"github.com/teo2022/go.generate_postm/utils"
	"io/ioutil"
	"log"
	"regexp"
	"strings"
)

func GroupRoute(list []models.ListRoute) []models.GroupRoute {
	var groupRoute []models.GroupRoute
	for _, route := range list {
		arName := utils.ClearArrString(strings.Split(route.Link, "/"))
		if len(arName) > 1 {

			oldId := GetIdGroupRout(groupRoute, arName[0])
			if oldId == -1 {
				var newR []models.ListRoute
				newR = append(newR, route)
				groupRoute = append(groupRoute, models.GroupRoute{Name: arName[0], Group: newR})
			} else {
				groupRoute[oldId].Group = append(groupRoute[oldId].Group, route)
			}
		} else {
			if len(arName) != 0 {
				var newR []models.ListRoute
				newR = append(newR, route)
				groupRoute = append(groupRoute, models.GroupRoute{Name: arName[0], Group: newR})
			}
		}
	}
	for i, v := range groupRoute {
		groupRoute[i].Name = utils.GetNameUp(v.Name)
	}
	return groupRoute
}

func GetIdGroupRout(list []models.GroupRoute, name string) int {
	for i, v := range list {
		if v.Name == name {
			return i
		}
	}
	return -1
}

func GetRoute(list []models.ListCatalog) []models.ListRoute {
	var AllRoute []models.ListRoute
	for _, v := range list {
		for _, r := range v.Files {
			input, err := ioutil.ReadFile(v.Patch + "/" + r)
			if err != nil {
				log.Fatalln(err)
			}
			lines := strings.Split(string(input), "\n")
			for _, line := range lines {
				if strings.Contains(line, "router.HandleFunc") {
					AllRoute = append(AllRoute, GetRoutReg(line))
				}
			}
		}
	}
	return AllRoute
}

func GetRoutReg(line string) models.ListRoute {
	new := models.ListRoute{}
	new.Origin = line
	re := regexp.MustCompile("\\\"(.*?)\\\"")
	match := re.FindAllStringSubmatch(line, -1)
	//fmt.Println(match)
	new.Link = match[0][1]
	//fmt.Println(len(match))
	if len(match) >= 2 {
		new.MetodFunc = match[1][1]
	}
	arName := utils.ClearArrString(strings.Split(new.Link, "/"))
	nameRoute := ""
	if len(arName) > 1 {
		for i, v := range arName {
			if i == 0 {
				continue
			}
			nameRoute = nameRoute + utils.GetNameUp(v)
		}
	} else {
		if len(arName) > 0 {
			nameRoute = utils.GetNameUp(arName[0])
		}
	}
	new.Name = nameRoute
	new.Folder, new.Function = getFuncRoute(line)

	return new
}

func getFuncRoute(line string) (string, string) {
	re := regexp.MustCompile("\\((.*?)\\)")
	match := re.FindAllStringSubmatch(line, -1)
	//fmt.Println(match)

	arName := utils.ClearArrString(strings.Split(match[0][1], ","))

	//fmt.Println(arName)
	arFunc := utils.ClearArrString(strings.Split(arName[1], "."))
	if len(arFunc) > 1 {
		return arFunc[0], arFunc[1]
	} else {
		return "/", arFunc[0]
	}
}
