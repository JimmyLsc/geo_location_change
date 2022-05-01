package core

import "testing"

func TestTencentURLToBaiduAPI(t *testing.T) {
	type args struct {
		tencentURL string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "test",
			args: args{
				tencentURL: "https://map.qq.com/?type=marker&isopeninfowin=1&markertype=1&pointx=114.075&pointy=32.1445&name=%E5%BA%9C%E9%83%BD%E8%8A%B1%E5%9B%AD2%E6%9C%9F(%E4%BF%A1%E9%98%B3%E5%B8%82%E5%B9%B3%E6%A1%A5%E5%8C%BA)&addr=%E5%B9%B3%E6%A1%A5%E5%8C%BA%E7%99%BD%E9%AB%98%E5%BA%99%E8%B7%AF&ref=WeChat\n",
			},
			want: "http://api.map.baidu.com/marker?location=32.1445,114.075&coord_type=gcj02&output=html&src=webapp.baidu.openAPIdemo",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TencentURLToBaiduAPI(tt.args.tencentURL); got != tt.want {
				t.Errorf("TencentURLToBaiduAPI() = %v, want %v", got, tt.want)
			}
		})
	}
}
