package utils

import (
	"fmt"
	"geo_location_change/model"
	"regexp"
)

// GetLocFromStr get location information from tencent location string
func GetLocFromStr(tencentLocStr string) model.Location {
	var location model.Location

	pointxStr := "pointx=[0-9.]*"
	pointyStr := "pointy=[0-9.]*"

	pointxReg, _ := regexp.Compile(pointxStr)
	pointyReg, _ := regexp.Compile(pointyStr)
	location.Longitude = fmt.Sprintf("%s", pointxReg.Find([]byte(tencentLocStr))[len("pointx="):])
	location.Latitude = fmt.Sprintf("%s", pointyReg.Find([]byte(tencentLocStr))[len("pointy="):])

	return location
}
