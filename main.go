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
		log.Printf("!!! please close the excel file when you use this application")
		return
	}
	log.Printf("success number: %d\n", count)
	log.Printf("Completed! You can open the file now!")
}
