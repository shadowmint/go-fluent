package pq

import "fluent/drivers/utils"
import "fmt"
import "strings"

/*============================================================================*
 * {{{ Internal api
 *============================================================================*/
 
type sqlFactory struct {
}

func newSqlFactory() *sqlFactory {
	return &sqlFactory{}
}

// SELECT [Columns[0], ...] FROM [Table] WHERE [Where] OFFSET [Offset] LIMIT [Limit]
// @return SQL statement, params for prepared execution or nil
func (self *sqlFactory) Select(stmt *sqlFactorySelect) (string, []interface{}) {
  if len(stmt.Columns) == 0 { return "", nil; }
  
  var cols = self.Columns(stmt.Columns)
  
  var limits = utils.New.Str()
  if stmt.Offset > 0 { limits.S(" OFFSET %d", stmt.Offset) }
  if stmt.Limit >= 0 { limits.S(" LIMIT %d", stmt.Limit) }
  
  var rtn = utils.New.Str()
  var whr = self.adjustWhere(stmt.Where)
  rtn.S("SELECT %s FROM %s WHERE %s%s", cols, stmt.Table, whr, limits.String())
  
	return rtn.String(), stmt.Values
}

// INSERT INTO [Table] ([Columns[0], Columns[1], ...) VALUES (?, ?)
// @return SQL statement, params for prepared execution or nil
func (self *sqlFactory) Insert(stmt *sqlFactoryInsert) (string, []interface{}) {
  if len(stmt.Columns) == 0 { return "", nil; }

  var cols = self.Columns(stmt.Columns)
  
  var size = len(stmt.Columns)
  var vals = utils.New.Str()
  for i := 0; i < size; i++ {
    vals.S("$%d", i + 1)
    if i != (size - 1) {
    	vals.S(", ")
    }
  }

  var returning = ""
  if len(stmt.Requested) > 0 {
    returning = " RETURNING " + self.Columns(stmt.Requested)
  }

  var rtn = utils.New.Str()
  rtn.S("INSERT INTO %s (%s) VALUES (%s)%s", stmt.Table, cols, vals.String(), returning)

  return rtn.String(), stmt.Values
}

// Escape a column name
func (self *sqlFactory) Column(name string) string {
	return fmt.Sprintf("\"%s\"", name)
}

// Return name, name, name for a set of columns
func (self *sqlFactory) Columns(columns []string) string {
  var cols = utils.New.Str()
  var size = len(columns)
  for i := 0; i < size; i++ {
    cols.S("%s", self.Column(columns[i]))
    if i != (size - 1) {
    	cols.S(", ")
    }
  }
  return cols.String()
}

// Postgres does not support the ? syntax, only :1, :2, etc.
// Correct this so that each ? is replaced by a :N. This is a safe
// operation because the '?' symbol is NEVER valid in the where clause
// other than as a prepared statement element.
func (self *sqlFactory) adjustWhere(whr string) string {
  var root = 1
  var rtn = whr
  for strings.Contains(rtn, "?") {
    var marker = fmt.Sprintf("$%d", root)
    rtn = strings.Replace(rtn, "?", marker, 1)
    root = root + 1
  }
  return rtn
}

/*============================================================================*
 * }}}
 *============================================================================*/
 
/*============================================================================*
 * {{{ Query types
 *============================================================================*/
 
type sqlFactoryInsert struct {
	Table string
	Columns []string
	Values []interface{}
	Requested []string
}

type sqlFactorySelect struct {
	Table string
	Where string
	Columns []string
	Values []interface{}
	Limit int
	Offset int
}
 
/*============================================================================*
 * }}}
 *============================================================================*/
