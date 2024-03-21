package modules

import (
	"errors"
	"fmt"
	"net/url"
	"path/filepath"
	"strings"

	"github.com/ChengYen-Tang/glassnode-crawler/models"
)

const fileExtension string = ".csv"

func GetActionName(apiUrl *string) (string, error) {
	// 解析 URL 以獲取 API 名稱
	parsedURL, err := url.Parse(*apiUrl)
	if err != nil {
		return "", errors.New("解析 URL 時發生錯誤: " + err.Error())
	}

	// 提取 URL path 的最後一部分作為 action name
	path := parsedURL.Path
	segments := strings.Split(path, "/")

	// 從 URL 路徑中取得 API 名稱
	return segments[len(segments)-1], nil
}

func GetAPIInfoDefault(apiUrl string, folder *string) *models.APIInfo {
	actionName, err := GetActionName(&apiUrl)
	if err != nil {
		panic(err)
	}
	return &models.APIInfo{
		APIUrl:     apiUrl,
		SaveFolder: *folder,
		FileName:   actionName + fileExtension,
	}
}

func GetAPIInfoTag(apiUrl string, folder *string, tag *string) *models.APIInfo {
	newUrl := fmt.Sprintf(apiUrl, *tag)
	actionName, err := GetActionName(&newUrl)
	if err != nil {
		panic(err)
	}
	return &models.APIInfo{
		APIUrl:     newUrl,
		SaveFolder: filepath.Join(*folder, actionName),
		FileName:   *tag + fileExtension,
	}
}
