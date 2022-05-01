package core

import (
	"fmt"
	"geo_location_change/utils"
)

func TencentURLToBaiduAPI(tencentURL string) string {
	location := utils.GetLocFromStr(tencentURL)
	apiURL := TransformBaiduAPI(location)
	fmt.Println(apiURL)
	return apiURL
}
