package health

import (
	"context"

	"github.com/kecci/go-toolkit/lib/sql"
)

// Repo structure repository for health
type Repo struct {
	db    *sql.DB
	stmts statements
}

type statements struct {
	check sql.ReadStatement
}

// New initialize repo Health
func New(db *sql.DB) (*Repo, error) {
	var (
		ctx = context.Background()
	)

	// create check statement
	getQuery := db.Rebind("SELECT 'active' status")
	checkStmt, err := db.PrepareRead(ctx, getQuery)
	if err != nil {
		return nil, err
	}

	return &Repo{
		db: db,
		stmts: statements{
			check: checkStmt,
		},
	}, nil
}
