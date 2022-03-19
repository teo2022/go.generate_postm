# Информация общая

Данный модуль формирует автоматическую документацию для постмана, работает с модулем `github.com/gorilla/mux`
Автоматически собирает все toure которые написанны в формате
`router.HandleFunc("adress", контролер).Methods("POST")`

Модуль так же собирает все структуры с проекта и формирует изних модели для body если модуль видет метод `POST,PUT`
то он находит нужную функцию и смотри что она принимает и формирует из это запрос! 

Так же модуль автоматически формирует папки и запросы к этой папке 

# Запуск модуля

Установите модуль к себе в проект коммандой `go get github.com/teo2022/go.generate_postm`

Пример работы 

```GO
package main

import (
	"fmt"
	"github.com/gorilla/mux"
	generate "github.com/teo2022/go.generate_postm"
	"path/filepath"
	"runtime"
)

func main() {
	_, b, _, _ := runtime.Caller(0)
	basepath   := filepath.Dir(b)

	generate.TeoStartGenerate("localhost", "9014", basepath, "apiDoc")

	router := mux.NewRouter()
	router.StrictSlash(true)

	router.HandleFunc("/", Base)

	router.HandleFunc("/config/get-list", controllers.GetListConfig).Methods("POST")
	router.HandleFunc("/config/update", controllers.UpdateConfig).Methods("POST")

	err := http.ListenAndServe("localhost:9014", router) 
	if err != nil {
		fmt.Println("Error api")
		fmt.Println(err)
	}
}
```