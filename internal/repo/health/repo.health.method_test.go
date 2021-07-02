package health

import (
	"context"
	"reflect"
	"testing"

	sqlmock "github.com/DATA-DOG/go-sqlmock"
	modelHealth "github.com/kecci/go-gql-microservice/internal/model/health"
	"github.com/kecci/go-toolkit/lib/sql"
	"github.com/stretchr/testify/require"
)

func TestRepo_CheckHealth(t *testing.T) {
	mockDB, _, err := sqlmock.New()
	require.NoError(t, err)
	defer mockDB.Close()

	db := sql.NewFromDB(mockDB, mockDB, "postgres")

	ctx := context.Background()

	type fields struct {
		db    *sql.DB
		stmts statements
	}
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *modelHealth.Health
		wantErr bool
	}{
		{
			"test function success",
			fields{
				db: db,
			},
			args{
				ctx: ctx,
			},
			&modelHealth.Health{
				Message: "SERVED",
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Repo{
				db:    tt.fields.db,
				stmts: tt.fields.stmts,
			}
			got, err := r.CheckHealth(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("Repo.CheckHealth() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Repo.CheckHealth() = %v, want %v", got, tt.want)
			}
		})
	}
}
