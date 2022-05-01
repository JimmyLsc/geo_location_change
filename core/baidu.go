package core

import (
	"fmt"
	"geo_location_change/model"
)

func TransformBaiduAPI(location model.Location) string {
	return fmt.Sprintf(baiduAPIPattern, location.Latitude, location.Longitude)
}
