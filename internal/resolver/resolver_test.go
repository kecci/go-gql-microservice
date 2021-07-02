package resolver

import (
	"reflect"
	"testing"

	gomock "github.com/golang/mock/gomock"
)

func TestNewResolver(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	healthSrvMock := NewMockhealthServiceInterface(ctrl)

	type args struct {
		healthService healthServiceInterface
	}
	tests := []struct {
		name string
		args args
		want *Resolver
	}{
		{
			"test new initialize success",
			args{
				healthService: healthSrvMock,
			},
			&Resolver{
				healthService: healthSrvMock,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewResolver(tt.args.healthService); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewResolver() = %v, want %v", got, tt.want)
			}
		})
	}
}
