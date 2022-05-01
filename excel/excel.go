package excel

import (
	"fmt"
	"geo_location_change/core"
	"github.com/xuri/excelize/v2"
)

func AddBaiduLink(filename string) (int, error) {
	file, err := excelize.OpenFile(filename)
	if err != nil {
		return 0, err
	}
	count := 0
	for row := 2; ; row++ {
		tencentLink, err := file.GetCellValue("Sheet1", "B"+fmt.Sprintf("%d", row))
		if err != nil {
			return 0, err
		}
		if tencentLink == "" {
			break
		}
		baiduLink := core.TencentURLToBaiduAPI(tencentLink)
		err = file.SetCellValue("Sheet1", "C"+fmt.Sprintf("%d", row), baiduLink)
		if err != nil {
			return 0, err
		}
		count++
	}
	if err := file.SaveAs(filename); err != nil {
		return 0, err
	}
	return count, nil
}
