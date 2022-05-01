package main

import (
	"geo_location_change/excel"
	"geo_location_change/utils"
	"log"
)

func main() {
	if err := utils.InitLog(); err != nil {
		panic(err)
	}
	count, err := excel.AddBaiduLink("test.xlsx")
	if err != nil {
		log.Printf("success number: %d, err: %s\n", count, err)
		return
	}
	log.Printf("success number: %d\n", count)
}
