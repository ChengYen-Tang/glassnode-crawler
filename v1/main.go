package main

import (
	"container/list"
	"fmt"
	"strings"

	"github.com/ChengYen-Tang/glassnode-crawler/api"
	"github.com/ChengYen-Tang/glassnode-crawler/models"
	"github.com/ChengYen-Tang/glassnode-crawler/modules"
	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigFile("./config.json")

	err := viper.ReadInConfig()
	if err != nil {
		panic("讀取設定檔錯誤: " + err.Error())
	}

	rootFolder := viper.GetString("rootFolder")
	cookieString := viper.GetString("cookieString")
	getter := modules.NewGetter(&cookieString)

	apiInfos := modules.NewTaskController(list.New())
	api.SetAddressApi(apiInfos, &rootFolder)
	api.SetBridgesApi(apiInfos, &rootFolder)
	api.SetBlockchainApi(apiInfos, &rootFolder)
	api.SetDefiApi(apiInfos, &rootFolder)
	api.SetDerivativesApi(apiInfos, &rootFolder)
	api.SetDistributionApi(apiInfos, &rootFolder)
	api.SetEntitiesApi(apiInfos, &rootFolder)
	api.SetEth20Api(apiInfos, &rootFolder)
	api.SetFeesApi(apiInfos, &rootFolder)
	api.SetIndicatorsApi(apiInfos, &rootFolder)
	api.SetInstitutionsApi(apiInfos, &rootFolder)
	api.SetLightningApi(apiInfos, &rootFolder)
	api.SetMarketApi(apiInfos, &rootFolder)
	api.SetMempoolApi(apiInfos, &rootFolder)
	api.SetMiningApi(apiInfos, &rootFolder)
	api.SetPointInTimeApi(apiInfos, &rootFolder)

	fmt.Println("API 數量:", apiInfos.ApiInfos.Len())

	for e := apiInfos.ApiInfos.Front(); e != nil; e = e.Next() {
		apiInfo := e.Value.(*models.APIInfo)
		err := getter.Download(&apiInfo.APIUrl, &apiInfo.SaveFolder, &apiInfo.FileName)
		if err != nil {
			fmt.Println(apiInfo.APIUrl)
			if strings.Contains(err.Error(), "professional") {
				fmt.Println(err.Error())
			} else {
				panic(err)
			}
		}
	}
}
