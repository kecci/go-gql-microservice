package health

import (
	"reflect"
	"testing"

	"github.com/golang/mock/gomock"
)

func TestNew(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	healthMock := NewMockhealthRepo(ctrl)
	type args struct {
		repoHealth healthRepo
	}
	tests := []struct {
		name string
		args args
		want *Service
	}{
		{
			"test success",
			args{
				repoHealth: healthMock,
			},
			&Service{
				healthRepo: healthMock,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(tt.args.repoHealth); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}
