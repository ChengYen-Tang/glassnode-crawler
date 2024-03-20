package modules

import (
	"errors"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

type Getter struct {
	cookies *[]*http.Cookie
}

func NewGetter(cookieString *string) *Getter {
	var cookies []*http.Cookie
	cookieStrings := strings.Split(*cookieString, ";")
	for _, cookieStr := range cookieStrings {
		parts := strings.Split(strings.TrimSpace(cookieStr), "=")
		if len(parts) == 2 {
			cookie := &http.Cookie{
				Name:  parts[0],
				Value: parts[1],
			}
			cookies = append(cookies, cookie)
		}
	}

	return &Getter{cookies: &cookies}
}

func (g *Getter) Download(apiUrl *string, saveFolderPath *string, fileName *string) error {
	req, err := http.NewRequest("GET", *apiUrl, nil)
	if err != nil {
		return errors.New("建立請求時發生錯誤: " + err.Error())
	}

	// 將解析後的 Cookie 加入到請求中
	for _, cookie := range *g.cookies {
		req.AddCookie(cookie)
	}

	// 發送 HTTP 請求並取得回應
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return errors.New("發送請求時發生錯誤: " + err.Error())
	}
	defer resp.Body.Close()

	// 檢查 HTTP 狀態碼
	if resp.StatusCode != http.StatusOK {
		body := make([]byte, 1024)
		n, _ := resp.Body.Read(body)
		return errors.New("無效的 HTTP 狀態碼: " + resp.Status + ", message: " + string(body[:n]))
	}

	// 檢查目錄和檔案是否存在，如果目錄不存在則建立目錄，如果檔案存在則刪除檔案
	if _, err := os.Stat(*saveFolderPath); os.IsNotExist(err) {
		os.Mkdir(*saveFolderPath, os.ModePerm)
	}
	filePath := filepath.Join(*saveFolderPath, *fileName)
	if _, err := os.Stat(filePath); err == nil {
		os.Remove(filePath)
	}

	// 建立檔案
	out, err := os.Create(filePath)
	if err != nil {
		return errors.New("無法創建目標檔案: " + err.Error())
	}
	defer out.Close()

	// 將回應體內容寫入目標檔案
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return errors.New("寫入檔案時發生錯誤: " + err.Error())
	}

	return nil
}
