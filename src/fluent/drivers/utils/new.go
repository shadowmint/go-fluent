package utils

import "database/sql"

/*============================================================================*
 * {{{ Internal api 
 *============================================================================*/

type factory struct {
	Rowset func(rows *sql.Rows, schema RowsetSchema) *Rowset
  Str func() Str
}

var New factory = factory {
	Rowset : newRowset,
  Str : newStr,
}

/*============================================================================*
 * }}}
 *============================================================================*/
