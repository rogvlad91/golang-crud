//go:build integration
// +build integration

package test

import (
	"context"
	"homework-7/test/postgres"
)

var Database *postgres.TDB

func init() {
	Database, _ = postgres.NewDB(context.Background())
}
