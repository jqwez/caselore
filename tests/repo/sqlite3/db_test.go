package sqlite_test

import (
	"github.com/jqwez/caselore/repo/sqlite"
	"testing"
)

func TestSQLiteConnection(t *testing.T) {
	db := sqlite.NewSQLiteConnection(":memory:")
	defer db.Close()
}
