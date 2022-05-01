package utils

import (
	"geo_location_change/model"
	"reflect"
	"testing"
)

func TestGetLocFromStr(t *testing.T) {
	type args struct {
		tencentLocStr string
	}
	tests := []struct {
		name string
		args args
		want model.Location
	}{
		{
			name: "test",
			args: args{
				tencentLocStr: "https://map.qq.com/?type=marker&isopeninfowin=1&markertype=1&pointx=114.075&pointy=32.1445&name=%E5%BA%9C%E9%83%BD%E8%8A%B1%E5%9B%AD2%E6%9C%9F(%E4%BF%A1%E9%98%B3%E5%B8%82%E5%B9%B3%E6%A1%A5%E5%8C%BA)&addr=%E5%B9%B3%E6%A1%A5%E5%8C%BA%E7%99%BD%E9%AB%98%E5%BA%99%E8%B7%AF&ref=WeChat\n",
			},
			want: model.Location{
				Latitude:  "32.1445",
				Longitude: "114.075",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetLocFromStr(tt.args.tencentLocStr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetLocFromStr() = %v, want %v", got, tt.want)
			}
		})
	}
}
