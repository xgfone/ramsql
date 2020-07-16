package engine_test

import (
	"database/sql"
	"testing"

	_ "github.com/xgfone/ramsql/driver"
	"github.com/xgfone/ramsql/engine/log"
)

func TestDrop(t *testing.T) {
	log.UseTestLogger(t)
	db, err := sql.Open("ramsql", "TestDrop")
	if err != nil {
		t.Fatalf("%s", err)
	}
	defer db.Close()

	_, err = db.Exec("CREATE TABLE account (id INT, email TEXT)")
	if err != nil {
		t.Fatalf("%s", err)
	}

	_, err = db.Exec("DROP TABLE account")
	if err != nil {
		t.Fatalf("cannot drop table: %s", err)
	}
}