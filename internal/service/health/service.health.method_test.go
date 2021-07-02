package health

import (
	"context"
	"reflect"
	"testing"

	"github.com/golang/mock/gomock"
	model "github.com/kecci/go-gql-microservice/internal/model/health"
)

func TestService_CheckHealth(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	healthMock := NewMockhealthRepo(ctrl)
	ctx := context.Background()

	type fields struct {
		healthRepo healthRepo
	}
	tests := []struct {
		name    string
		fields  fields
		mockFn  func()
		want    *model.Health
		wantErr bool
	}{
		{
			"test function success",
			fields{
				healthRepo: healthMock,
			},
			func() {
				healthMock.EXPECT().CheckHealth(ctx).Return(&model.Health{
					Message: "SERVED",
				}, nil)
			},
			&model.Health{
				Message: "SERVED",
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockFn()
			s := &Service{
				healthRepo: tt.fields.healthRepo,
			}
			got, err := s.CheckHealth()
			if (err != nil) != tt.wantErr {
				t.Errorf("Service.CheckHealth() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Service.CheckHealth() = %v, want %v", got, tt.want)
			}
		})
	}
}
