package main

import (
	"encoding/json"
	"fmt"
	"log"
	"mangazdl/meta"
	"os"
	"strconv"
)

func Start() {
	if len(os.Args) < 2 {
		log.Printf("Usage: go run . <manga id>")
		return
	}
	mangaIdToDownload, _ := strconv.Atoi(os.Args[1])

	var fName string
	showImgsLogs := true
	folderNameImgs := "./result_imgs"
	folderNameMetadata := "./result_meta"
	imagesThreadsNumber := 30
	if err := os.MkdirAll(folderNameImgs, 0644); err != nil {
		log.Fatalln("error while creating a folder err: ", err)
	}
	if err := os.MkdirAll(folderNameMetadata, 0644); err != nil {
		log.Fatalln("error while creating a folder err: ", err)
	}

	fName = fmt.Sprintf("%s/%d.json", folderNameMetadata, mangaIdToDownload)
	_, err := os.Stat(fName)
	if err == nil {
		log.Printf("Manga ID %d has already downloaded", mangaIdToDownload)
		return
	}
	if !os.IsNotExist(err) {
		log.Println("Issue reading file: ", err)
	}
	data, err := meta.GetFromID(mangaIdToDownload)
	if err != nil {
		log.Println("err: ", err)
		return
	}
	if data.Book.Baid == 0 {
		log.Printf("Manga ID %d not found", mangaIdToDownload)
		return
	}
	wasThereAndError := data.SetImgs(showImgsLogs, imagesThreadsNumber)
	if wasThereAndError { //
		return
	}
	mainFolderName := fmt.Sprintf("%s/%d", folderNameImgs, mangaIdToDownload)
	if err := os.Mkdir(mainFolderName, 0644); err != nil {
		log.Println("error while creating a folder err: ", err)
		return
	}
	var issueSavingImg bool
	for _, order := range data.Orders {
		if len(order.ImgRaw) == 0 {
			continue
		}
		fName = fmt.Sprintf("%s/%d.png", mainFolderName, order.No)
		if err := os.WriteFile(fName, order.ImgRaw, 0644); err != nil {
			log.Println("err saving img err: ", err)
			issueSavingImg = true
		}
	}
	if issueSavingImg {
		return
	}
	fName = fmt.Sprintf("%s/%d.json", folderNameMetadata, mangaIdToDownload)
	f, err := os.Create(fName)
	if err != nil {
		log.Println("err saving json data err: ", err)
		return
	}
	if err := json.NewEncoder(f).Encode(data); err != nil {
		log.Println("err saving json data err: ", err)
		return
	}
	if err := f.Close(); err != nil {
		log.Println("err closing json data err: ", err)
		return
	}
	log.Printf("Manga ID %d successfully saved", mangaIdToDownload)
}
