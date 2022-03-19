package controller

import (
	"fmt"
	"github.com/teo2022/go.generate_postm/models"
	"io/ioutil"
	"log"
	"strings"
)

func GetFolders(patch string) []models.ListCatalog {
	files, err := ioutil.ReadDir(patch)
	if err != nil {
		log.Fatal(err)
	}
	var allListCatalog  []models.ListCatalog
	lCatalog := models.ListCatalog{}
	lCatalog.Catalog = "/"
	lCatalog.Patch = patch + "/"
	for _, file := range files {
		if file.IsDir() {
			allListCatalog = append(allListCatalog, models.ListCatalog{Catalog: file.Name()})
		} else {
			lCatalog.Files = append(lCatalog.Files, file.Name())
		}
	}

	var finListCatalog  []models.ListCatalog
	finListCatalog = RecursionFolder(finListCatalog, allListCatalog, patch)
	finListCatalog = append(finListCatalog, lCatalog)
	fmt.Println(finListCatalog)
	return finListCatalog
}

func RecursionFolder(finListCatalog []models.ListCatalog, allListCatalog []models.ListCatalog, patch string) []models.ListCatalog {
	for _,v := range allListCatalog {
		patch2 := patch + "/" + v.Catalog
		if strings.Contains(v.Catalog, ".") {
			continue
		}
		files, err := ioutil.ReadDir(patch2)
		if err != nil {
			log.Fatal(err)
		}
		lCatalog := models.ListCatalog{}
		lCatalog.Catalog = v.Catalog
		lCatalog.Patch = patch2
		for _, file := range files {
			if file.IsDir() {
				allListCatalog = append(allListCatalog, models.ListCatalog{Catalog: file.Name()})
			} else {
				lCatalog.Files = append(lCatalog.Files, file.Name())
			}
		}
		finListCatalog = append(finListCatalog, lCatalog)
	}
	return finListCatalog
}
