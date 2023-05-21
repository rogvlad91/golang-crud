//go:build integration
// +build integration

package postgres

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	db "golang-crud/pkg/db/pg"
	"log"
	"strings"
	"sync"
	"testing"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "test"
	password = "test"
	dbname   = "test"
)

type TDB struct {
	sync.Mutex
	Db *db.Database
}

func NewDB(ctx context.Context) (*TDB, error) {
	dsn := generateDsn()
	pool, err := pgxpool.Connect(ctx, dsn)
	if err != nil {
		return nil, err
	}

	db := db.NewDatabase(pool)
	return &TDB{Db: db}, err
}

func generateDsn() string {
	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
}

func (d *TDB) SetUp(t *testing.T) {
	t.Helper()
	ctx := context.Background()

	d.Lock()
	d.Truncate(ctx)
}

func (d *TDB) TearDown(t *testing.T) {
	defer d.Unlock()
	ctx := context.Background()
	d.Truncate(ctx)
}

func (d *TDB) Truncate(ctx context.Context) {
	var tables []string

	err := d.Db.Select(ctx, &tables, `SELECT table_name FROM information_schema.tables WHERE table_schema='public' AND table_type = 'BASE TABLE' AND table_name != 'goose_db_version'`)
	if err != nil {
		log.Fatalf(err.Error())
	}

	if len(tables) == 0 {
		panic("run migrations")
	}

	q := fmt.Sprintf("truncate table %s", strings.Join(tables, ","))

	if _, err := d.Db.Exec(ctx, q); err != nil {
		panic(err)
	}
}
