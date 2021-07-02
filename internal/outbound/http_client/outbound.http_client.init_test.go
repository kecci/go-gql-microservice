package http_client

import (
	"net/http"
	"reflect"
	"testing"

	"github.com/golang/mock/gomock"
)

func TestInitHttpClient(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	httpClientMock := &http.Client{}
	type args struct {
		httpClientOutbound *http.Client
	}
	tests := []struct {
		name string
		args args
		want *HttpClientOutbound
	}{
		{
			"test success",
			args{
				httpClientOutbound: httpClientMock,
			},
			&HttpClientOutbound{
				httpClientOutbound: httpClientMock,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewHttpClientOutbound(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewHttpClientOutbound() = %v, want %v", got, tt.want)
			}
		})
	}
}
