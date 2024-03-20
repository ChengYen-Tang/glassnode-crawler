package main

import (
	"fmt"
	"strings"

	"github.com/ChengYen-Tang/glassnode-crawler/models"
	"github.com/ChengYen-Tang/glassnode-crawler/modules"
)

func main() {
	cookieString := ""
	getter := modules.NewGetter(&cookieString)

	apiInfos := make([]*models.APIInfo, 0)

	for _, apiInfo := range apiInfos {
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
