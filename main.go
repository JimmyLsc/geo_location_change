package main

import (
	"fmt"
	"geo_location_change/excel"
)

func main() {
	count, err := excel.AddBaiduLink("test.xlsx")
	if err != nil {
		fmt.Printf("success number: %d, err: %s", count, err)
		return
	}
	fmt.Printf("success number: %d", count)
}
