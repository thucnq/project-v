package sqlcbuilder

import (
	"database/sql"
)

var _ DBTX = (*DB)(nil)

type DB struct {
	*sql.DB
}
