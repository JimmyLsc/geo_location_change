package core

import (
	"geo_location_change/utils"
	"log"
)

func TencentURLToBaiduAPI(tencentURL string) string {
	location := utils.GetLocFromStr(tencentURL)
	apiURL := TransformBaiduAPI(location)
	log.Println(apiURL + "\n")
	return apiURL
}
