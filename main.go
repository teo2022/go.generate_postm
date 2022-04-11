package generate

import (
	"github.com/teo2022/go.generate_postm/controller"
)

//func main() {
//	TeoStartGenerate("localhost","9008","/Users/alex/GolandProjects/go.store", "store")
//}

func TeoStartGenerate(constUrl string, constPort string, patch string, name string) {
	AllFolder := controller.GetFolders(patch)
	allRoute := controller.GetRoute(AllFolder)
	groupRoute := controller.GroupRoute(allRoute)
	allStruct := controller.GetStruct(AllFolder)
	groupRoute = controller.GetRawBody(groupRoute, allStruct, AllFolder)
	controller.GeneratePostman(patch, groupRoute, constUrl, constPort, name)
}
