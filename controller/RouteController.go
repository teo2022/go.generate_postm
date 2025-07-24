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

func GetRouteChi(list []models.ListCatalog) []models.GroupRoute {
	var AllRoute []models.GroupRoute
	for _, v := range list {
		for _, r := range v.Files {
			input, err := ioutil.ReadFile(v.Patch + "/" + r)
			if err != nil {
				log.Fatalln(err)
			}
			lines := strings.Split(string(input), "\n")
			isStart := false
			listLine := []string{}
			for _, line := range lines {
				if strings.Contains(line, "r.Route") {
					isStart = true
				}
				if isStart {
					listLine = append(listLine, line)
				}
				if strings.Contains(line, "})") && isStart {
					isStart = false
					AllRoute = append(AllRoute, GetRoutRegChi(listLine))
					listLine = []string{}
				}
			}
		}
	}
	return AllRoute
}

func GetRoutRegChi(list []string) models.GroupRoute {
	var groupResult models.GroupRoute

	// Регулярное выражение для извлечения методов и путей
	// Разбиваем входную строку на логические части
	matches := splitIntoLogicalLines(strings.Join(list, "\n"))

	// Извлекаем имя группы из первой строки
	groupPath := strings.Trim(strings.ReplaceAll(matches[0], "r.Route(\"", ""), "\")")
	groupName := strings.Trim(groupPath, "/")
	groupResult.Name = utils.GetNameUp(groupName)

	// Проходим по всем строкам и извлекаем маршруты
	for _, line := range matches[1:] { // Пропускаем первую строку (группу)
		parts := strings.SplitN(line, "\"", 2) // Разделяем строку на метод и путь
		if len(parts) < 2 {
			continue // Пропускаем некорректные строки
		}

		// Извлекаем метод
		methodAndPath := parts[0]
		path := parts[1]

		// Разделяем метод и путь
		methodParts := strings.Split(methodAndPath, "(")
		if len(methodParts) < 2 {
			continue // Пропускаем некорректные строки
		}
		method := strings.ToUpper(strings.TrimSpace(methodParts[0][2:])) // Удаляем "r." и пробелы
		path = strings.TrimSpace(path)
		path = strings.Replace(path, "\"", "", -1)
		// Формируем название маршрута
		name := getNameFromPath(path)

		// Добавляем маршрут в группу
		groupResult.Group = append(groupResult.Group, models.ListRoute{
			Name:      name,
			Link:      fmt.Sprintf("/%v%v", groupName, path),
			MetodFunc: method,
		})
	}
	return groupResult
}

// splitIntoLogicalLines - разбивает входную строку на логические строки
func splitIntoLogicalLines(input string) []string {
	// Регулярное выражение для извлечения метода и пути
	re := regexp.MustCompile(`r\.\w+\(\s*"[^"]+"`)

	// Находим все совпадения
	matches := re.FindAllString(input, -1)

	// Убираем лишние пробелы
	for i, match := range matches {
		match = strings.Replace(match, "\t", "", -1)
		match = strings.Replace(match, "\n", "", -1)
		matches[i] = strings.TrimSpace(match)
	}
	return matches
}

// getIndentLevel - определяет уровень вложенности строки по количеству пробелов или табуляций
func getIndentLevel(line string) int {
	trimmed := strings.TrimLeft(line, "\t")
	return len(line) - len(trimmed)
}

// getNameFromPath - генерирует название маршрута на основе пути
func getNameFromPath(path string) string {
	parts := strings.Split(strings.Trim(path, "/"), "/")
	if len(parts) == 0 {
		return "Unknown"
	}
	return strings.Title(strings.ReplaceAll(parts[0], "-", " "))
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
