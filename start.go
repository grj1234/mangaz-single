package main

import (
	"encoding/json"
	"fmt"
	"log"
	"mangaz/meta"
	"os"
	"sync"
)

func Start() {
	idStartFrom := 1
	showImgsLogs := true
	folderNameImgs := "./result_imgs"
	folderNameMetadata := "./result_meta"
	mainThreadsNumber := 30
	imagesThreadsNumber := 30
	if err := os.MkdirAll(folderNameImgs, 0644); err != nil {
		log.Fatalln("error while creating a folder err: ", err)
	}
	if err := os.MkdirAll(folderNameMetadata, 0644); err != nil {
		log.Fatalln("error while creating a folder err: ", err)
	}
	var locker sync.Mutex
	var alreadySavedNumber, currentSaved int
	chanNumber := make(chan int)
	for i := 0; i < mainThreadsNumber; i++ {
		go func() {
			for number := range chanNumber {
				data, err := meta.GetFromID(number)
				if err != nil {
					log.Println("err: ", err)
					continue
				}
				if data.Book.Baid == 0 {
					log.Printf("Mangas already saved: %d saved on current session: %d current number: %d status: NOT FOUND", alreadySavedNumber, currentSaved, number)
					continue
				}
				wasThereAndError := data.SetImgs(showImgsLogs, imagesThreadsNumber)
				if wasThereAndError { //
					continue
				}
				mainFolderName := fmt.Sprintf("%s/%d", folderNameImgs, number)
				if err := os.Mkdir(mainFolderName, 0644); err != nil {
					log.Println("error while creating a folder err: ", err)
					continue
				}
				var issueSavingImg bool
				for _, order := range data.Orders {
					if len(order.ImgRaw) == 0 {
						continue
					}
					fName := fmt.Sprintf("%s/%d.jpg", mainFolderName, order.No)
					if err := os.WriteFile(fName, order.ImgRaw, 0644); err != nil {
						log.Println("err saving img err: ", err)
						issueSavingImg = true
					}
				}
				if issueSavingImg {
					continue
				}
				fName := fmt.Sprintf("%s/%d.json", folderNameMetadata, number)
				f, err := os.Create(fName)
				if err != nil {
					log.Println("err saving json data err: ", err)
					continue
				}
				if err := json.NewEncoder(f).Encode(data); err != nil {
					log.Println("err saving json data err: ", err)
					continue
				}
				if err := f.Close(); err != nil {
					log.Println("err closing json data err: ", err)
					continue
				}
				locker.Lock()
				currentSaved++
				log.Printf("Mangas already saved: %d saved on current session: %d current number: %d status: FOUND", alreadySavedNumber, currentSaved, number)
				locker.Unlock()
			}
		}()
	}
	for ; ; idStartFrom++ {
		fName := fmt.Sprintf("%s/%d.json", folderNameMetadata, idStartFrom)
		_, err := os.Stat(fName)
		if err == nil {
			locker.Lock()
			alreadySavedNumber++
			locker.Unlock()
			continue
		}
		if !os.IsNotExist(err) {
			log.Println("Issue reading file: ", err)
		}
		chanNumber <- idStartFrom
	}

}
