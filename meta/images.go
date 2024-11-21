package meta

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"sync"
)

func (data Data) SetImgs(showLogs bool, imagesThreadsNumber int) bool {
	var wasThereAndError bool
	chanIndex := make(chan int)
	var locker sync.Mutex
	var progress int
	var wg sync.WaitGroup
	for i := 0; i < imagesThreadsNumber; i++ {
		go func() {
			for index := range chanIndex {
				order := data.Orders[index]
				_, imgRaw, err := order.GetImg()
				if err != nil {
					log.Println("err getting image: ", order.URL, err)
					wasThereAndError = true
				}
				convertedImg, err := order.Convert(imgRaw)
				if err == nil {
					data.Orders[index].ImgRaw = convertedImg
				} else {
					log.Println("err converting image: ", order.URL, err)
					wasThereAndError = true
				}
				data.Orders[index].Scramble = nil
				locker.Lock()
				progress++
				if showLogs {
					log.Printf("Progress getting img from id: %d %d/%d image len: %d-%d\n", data.Book.Baid, progress, len(data.Orders), len(imgRaw), len(convertedImg))
				}
				locker.Unlock()
				wg.Done()
			}
		}()
	}
	for i := range data.Orders {
		chanIndex <- i
		wg.Add(1)
	}
	wg.Wait()
	close(chanIndex)
	return wasThereAndError
}
func (Order Order) GetImg() (http.Header, []byte, error) {
	req, err := http.NewRequest("GET", Order.URL, nil)
	if err != nil {
		return nil, nil, err
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
		return nil, nil, err
	}
	if response.StatusCode != 200 {
		err := fmt.Errorf("not a correct status code: %d headers: %+v", response.StatusCode, response.Header)
		return nil, nil, err
	}
	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, nil, err
	}
	return response.Header, body, nil
}
