package modules

import (
	"container/list"
	"fmt"
	"path/filepath"

	"github.com/ChengYen-Tang/glassnode-crawler/models"
)

type TaskController struct {
	ApiInfos *list.List
}

func NewTaskController(apiInfos *list.List) *TaskController {
	return &TaskController{
		ApiInfos: apiInfos,
	}
}

func (controller *TaskController) GetAPIInfoDefault(apiUrl string, folder *string, setCurrency bool) {
	actionName, err := GetActionName(&apiUrl)
	if err != nil {
		panic(err)
	}

	if setCurrency {
		currencyUrl := apiUrl + "&c=usd"
		fileName := actionName + "-usd" + fileExtension
		controller.ApiInfos.PushBack(&models.APIInfo{
			APIUrl:     currencyUrl,
			SaveFolder: *folder,
			FileName:   fileName,
		})
		currencyUrl = apiUrl + "&c=native"
		fileName = actionName + "-native" + fileExtension
		controller.ApiInfos.PushBack(&models.APIInfo{
			APIUrl:     currencyUrl,
			SaveFolder: *folder,
			FileName:   fileName,
		})
	} else {
		controller.ApiInfos.PushBack(&models.APIInfo{
			APIUrl:     apiUrl,
			SaveFolder: *folder,
			FileName:   actionName + fileExtension,
		})
	}
}

func (controller *TaskController) GetAPIInfoTag(apiUrl string, folder *string, tag *string, setCurrency bool) {
	newUrl := fmt.Sprintf(apiUrl, *tag)
	actionName, err := GetActionName(&newUrl)
	if err != nil {
		panic(err)
	}

	if setCurrency {
		currencyUrl := newUrl + "&c=usd"
		fileName := *tag + "-usd" + fileExtension
		controller.ApiInfos.PushBack(&models.APIInfo{
			APIUrl:     currencyUrl,
			SaveFolder: filepath.Join(*folder, actionName),
			FileName:   fileName,
		})
		currencyUrl = newUrl + "&c=native"
		fileName = *tag + "-native" + fileExtension
		controller.ApiInfos.PushBack(&models.APIInfo{
			APIUrl:     currencyUrl,
			SaveFolder: filepath.Join(*folder, actionName),
			FileName:   fileName,
		})
	} else {
		controller.ApiInfos.PushBack(&models.APIInfo{
			APIUrl:     newUrl,
			SaveFolder: filepath.Join(*folder, actionName),
			FileName:   *tag + fileExtension,
		})
	}
}

func (controller *TaskController) GetAPIInfoTwoTag(apiUrl string, folder *string, tag1 *string, tag2 *string, setCurrency bool) {
	newUrl := fmt.Sprintf(apiUrl, *tag1, *tag2)
	actionName, err := GetActionName(&newUrl)
	if err != nil {
		panic(err)
	}

	if setCurrency {
		currencyUrl := newUrl + "&c=usd"
		fileName := *tag2 + "-usd" + fileExtension
		controller.ApiInfos.PushBack(&models.APIInfo{
			APIUrl:     currencyUrl,
			SaveFolder: filepath.Join(*folder, actionName, *tag1),
			FileName:   fileName,
		})
		currencyUrl = newUrl + "&c=native"
		fileName = *tag2 + "-native" + fileExtension
		controller.ApiInfos.PushBack(&models.APIInfo{
			APIUrl:     currencyUrl,
			SaveFolder: filepath.Join(*folder, actionName, *tag1),
			FileName:   fileName,
		})
	} else {
		controller.ApiInfos.PushBack(&models.APIInfo{
			APIUrl:     newUrl,
			SaveFolder: filepath.Join(*folder, actionName, *tag1),
			FileName:   *tag2 + fileExtension,
		})
	}
}
