package generate

import (
	"github.com/teo2022/go.generate_postm/controller"
)

func TeoStartGenerate(constUrl string, constPort string, patch string, name string, chi bool) {
	if chi {
		AllFolder := controller.GetFolders(patch)
		groupRoute := controller.GetRouteChi(AllFolder)
		allStruct := controller.GetStruct(AllFolder)
		groupRoute = controller.GetRawBodyChi(groupRoute, allStruct, AllFolder)
		controller.GeneratePostman(patch, groupRoute, constUrl, constPort, name)
	} else {
		AllFolder := controller.GetFolders(patch)
		allRoute := controller.GetRoute(AllFolder)
		groupRoute := controller.GroupRoute(allRoute)
		allStruct := controller.GetStruct(AllFolder)
		groupRoute = controller.GetRawBody(groupRoute, allStruct, AllFolder)
		controller.GeneratePostman(patch, groupRoute, constUrl, constPort, name)
	}

}
