package controller

import (
	"github.com/teo2022/go.generate_postm/models"
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
		group[i].Group[0].RawBody = schemModel
		group[i].Group[1].RawBody = schemModel
		group[i].Group[2].RawBody = schemModel
		group[i].Group[3].RawBody = schemFilter
		group[i].Group[4].RawBody = schemModel
	}
	return group
}
