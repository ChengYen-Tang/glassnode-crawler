package modules

import (
	"errors"
	"net/url"
	"strings"
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
