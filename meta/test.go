package meta

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

func Test() {
	test1()
}
func test0() {
	id := 101
	data, err := GetFromID(id)
	if err != nil {
		log.Println("err: ", err)
	}
	data.SetImgs(true, 20)
	os.MkdirAll("./test", 0644)

	_, imgRaw, err := data.Orders[0].GetImg()
	if err != nil {
		log.Println("err: ", err)
	}
	convertedImg, err := data.Orders[0].Convert(imgRaw)

	if err != nil {
		log.Println("err: ", err)
	}
	os.WriteFile("./input.jpg", imgRaw, 0644)
	os.WriteFile("./output.png", convertedImg, 0644)
	f, _ := os.Create("./data.json")
	json.NewEncoder(f).Encode(data)
}

func test1() {
	id := 101
	data, err := GetFromID(id)
	if err != nil {
		log.Println("err: ", err)
	}
	data.SetImgs(true, 20)
	os.MkdirAll("./test", 0644)
	for i, order := range data.Orders {
		if len(order.ImgRaw) == 0 {
			continue
		}
		fName := fmt.Sprintf("./test/%d.png", i)
		os.WriteFile(fName, order.ImgRaw, 0644)
	}
	f, _ := os.Create("./data.json")
	json.NewEncoder(f).Encode(data)
}
