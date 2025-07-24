package controller

import (
	"github.com/teo2022/go.generate_postm/models"
	"strings"
)

// Находим роуты который принимают боду
func GetRawBodyChi(group []models.GroupRoute, model []models.ListModel, AllFolder []models.ListCatalog) []models.GroupRoute {
	for i, v := range group {
		schemModel := ""
		for _, v2 := range model {
			if v2.Name == v.Name {
				schemModel = v2.Jsonschema
				break
			}
		}
		schemFilter := ""
		for _, v2 := range model {
			if v2.Name == "Filter" {
				schemFilter = v2.Jsonschema
				break
			}
		}
		for i2, v3 := range group[i].Group {
			if v3.MetodFunc == "POST" || v3.MetodFunc == "PUT" {
				if strings.Contains(v3.Link, "get-list") {
					group[i].Group[i2].RawBody = schemFilter
				} else {
					group[i].Group[i2].RawBody = schemModel
				}
			} else {
				group[i].Group[i2].RawBody = ""
			}
		}
	}
	return group
}
