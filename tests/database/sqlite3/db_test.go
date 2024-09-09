package sqlite_test

import (
	"github.com/jqwez/caselore/database/sqlite"
	"testing"
)

func TestSQLiteConnection(t *testing.T) {
	db := sqlite.NewSQLiteConnection(":memory:")
	defer db.Close()
}
