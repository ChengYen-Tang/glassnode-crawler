package main

import (
	"container/list"
	"fmt"
	"strings"

	api "github.com/ChengYen-Tang/glassnode-crawler/Api"
	"github.com/ChengYen-Tang/glassnode-crawler/models"
	"github.com/ChengYen-Tang/glassnode-crawler/modules"
)

func main() {
	rootFolder := ""
	cookieString := ""
	getter := modules.NewGetter(&cookieString)

	apiInfos := list.New()
	api.SetAddressApi(apiInfos, &rootFolder)
	api.SetBridgesApi(apiInfos, &rootFolder)

	for e := apiInfos.Front(); e != nil; e = e.Next() {
		apiInfo := e.Value.(*models.APIInfo)
		err := getter.Download(&apiInfo.APIUrl, &apiInfo.SaveFolder, &apiInfo.FileName)
		if err != nil {
			fmt.Println(apiInfo.APIUrl)
			if strings.Contains(err.Error(), "無效的 HTTP 狀態碼") {
				fmt.Println(err.Error())
			} else {
				panic(err)
			}
		}
	}
}
