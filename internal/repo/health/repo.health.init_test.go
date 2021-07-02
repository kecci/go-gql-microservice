package health

import (
	"testing"

	sqlmock "github.com/DATA-DOG/go-sqlmock"
	"github.com/kecci/go-toolkit/lib/sql"
	"github.com/stretchr/testify/require"
)

func TestNew(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer mockDB.Close()

	db := sql.NewFromDB(mockDB, mockDB, "postgres")

	type args struct {
		db *sql.DB
	}
	tests := []struct {
		name    string
		args    args
		fnMock  func()
		want    *Repo
		wantErr bool
	}{
		{
			"test initialize success",
			args{
				db: db,
			},
			func() {
				mock.ExpectPrepare("SELECT 'active' status")
			},
			&Repo{
				db: db,
			},
			false,
		},
		{
			"test initialize fail",
			args{
				db: db,
			},
			func() {
				// mock.ExpectPrepare("SELECT 'active' status")
			},
			&Repo{
				db: db,
			},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.fnMock()
			_, err := New(tt.args.db)
			if (err != nil) != tt.wantErr {
				t.Errorf("New() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
