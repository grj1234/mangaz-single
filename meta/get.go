package meta

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

func GetFromID(id int) (Data, error) {
	urlToUse := fmt.Sprintf(ep, id)
	req, err := http.NewRequest("GET", urlToUse, nil)
	if err != nil {
		return Data{}, err
	}
	req.Header.Set("accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9")
	req.Header.Set("accept-language", "en")
	req.Header.Set("cache-control", "no-cache")
	req.Header.Set("pragma", "no-cache")
	req.Header.Set("sec-ch-ua", `"Chromium";v="92", " Not A;Brand";v="99", "Google Chrome";v="92"`)
	req.Header.Set("sec-ch-ua-mobile", "?0")
	req.Header.Set("sec-fetch-dest", "document")
	req.Header.Set("sec-fetch-mode", "navigate")
	req.Header.Set("sec-fetch-site", "same-origin")
	req.Header.Set("sec-fetch-user", "?1")
	req.Header.Set("upgrade-insecure-requests", "1")
	req.Header.Set("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/92.0.4515.159 Safari/537.36")
	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}
	response, err := client.Do(req)
	if err != nil {
		return Data{}, err
	}
	if response.StatusCode == 404 {
		return Data{}, nil
	}
	if response.StatusCode != 200 {
		err := fmt.Errorf("not a correct status code: %d headers: %+v", response.StatusCode, response.Header)
		return Data{}, err
	}
	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return Data{}, err
	}
	reader := bytes.NewReader(body)
	doc, err := goquery.NewDocumentFromReader(reader)
	if err != nil {
		return Data{}, err
	}
	dataRaw := doc.Find("#doc").Text()
	decodedBytes, err := base64.StdEncoding.DecodeString(string(dataRaw))
	if err != nil {
		return Data{}, err
	}
	var data Data
	if err := json.Unmarshal(decodedBytes, &data); err != nil {
		return Data{}, err
	}
	for i, order := range data.Orders {
		urlToUse := fmt.Sprintf("%s%s/%s", data.Location.Base, data.Location.Scramble_dir, order.Name)
		data.Orders[i].URL = urlToUse
	}
	return data, nil
}
