package pq

import "fluent/drivers/tests"
import "testing"

/*============================================================================*
 * {{{ Test constants
 *============================================================================*/
 
/*============================================================================*
 * {{{ Tests
 *============================================================================*/

func Test_sql_factory_can_create_instance(T *testing.T) {
  var a, i = sqlFactoryTestSetup(T)

  a.NotNil(i, "Unable to create instance")

  sqlFactoryTestTeardown(i)
}

func Test_sql_factory_can_handle_whr_stmts(T *testing.T) {
  var a, i = sqlFactoryTestSetup(T)

  var output = i.adjustWhere("")
  a.Equals(output, "", "Where statement adjust failed on no vars")

  output = i.adjustWhere("x = ?")
  a.Equals(output, "x = $1", "Where statement adjust failed on 1 var")

  output = i.adjustWhere("x = ? AND y = ? AND z = ?")
  a.Equals(output, "x = $1 AND y = $2 AND z = $3", "Where statement adjust failed on 3 vars")

  sqlFactoryTestTeardown(i)
}

func Test_sql_factory_can_process_insert_statement(T *testing.T) {
  var a, i = sqlFactoryTestSetup(T)

  var stmt = sqlFactoryInsert {
    Table : "HelloTable",
    Columns : []string { "date", "key", "value" },
    Values : []interface{} { 0, "Hello", "World" },
  }

  var sql, vals = i.Insert(&stmt)
  a.NotNil(vals, "Value set for insert is invalid")
  a.Equals(sql, "INSERT INTO HelloTable (\"date\", \"key\", \"value\") VALUES ($1, $2, $3)", "Invalid SQL")

  sqlFactoryTestTeardown(i)
}

func Test_sql_factory_can_process_select_statement(T *testing.T) {
  var a, i = sqlFactoryTestSetup(T)

  var stmt = sqlFactorySelect {
    Table : "HelloTable",
    Columns : []string { "date", "key" },
    Limit : 4,
    Offset : 2,
    Values : []interface{} { 0, "World" },
    Where : "date = ? AND key = ?",
  }

  var sql, vals = i.Select(&stmt)
  a.NotNil(vals, "Value set for select is invalid")
  a.Equals(sql, "SELECT \"date\", \"key\" FROM HelloTable WHERE date = $1 AND key = $2 OFFSET 2 LIMIT 4", "Invalid SQL")
  
  stmt.Limit = -1
  stmt.Offset = -1
  sql, _ = i.Select(&stmt)
  a.Equals(sql, "SELECT \"date\", \"key\" FROM HelloTable WHERE date = $1 AND key = $2", "Invalid SQL")

  stmt.Limit = 1
  stmt.Offset = -1
  sql, _ = i.Select(&stmt)
  a.Equals(sql, "SELECT \"date\", \"key\" FROM HelloTable WHERE date = $1 AND key = $2 LIMIT 1", "Invalid SQL")

  stmt.Limit = -1
  stmt.Offset = 1
  sql, _ = i.Select(&stmt)
  a.Equals(sql, "SELECT \"date\", \"key\" FROM HelloTable WHERE date = $1 AND key = $2 OFFSET 1", "Invalid SQL")

  sqlFactoryTestTeardown(i)
}

/*============================================================================*
 * }}}
 *============================================================================*/
 
/*============================================================================*
 * {{{ Helpers
 *============================================================================*/

func sqlFactoryTestSetup(T *testing.T) (tests.Assert, *sqlFactory) {
  var assert = tests.New.Assert(T)
  var instance = newSqlFactory()
  return assert, instance
}

func sqlFactoryTestTeardown(instance *sqlFactory) {
}

/*============================================================================*
 * }}}
 *============================================================================*/
 
