package mysql

import nsql "fluent"
import sqltests "fluent/drivers/tests"
import "fluent/drivers/tests"
import "testing"

/*============================================================================*
 * {{{ Test constants
 *============================================================================*/
 
 // The URI to connect to a testing database
const TEST_ROOT_URI = "root:password@tcp(localhost:3306)/?charset=utf8&autocommit=true"
const TEST_URI = "root:password@tcp(localhost:3306)/testing?charset=utf8&autocommit=true"

// Disable tests if no mysql instance is available.
const TEST_RUN = true
 
/*============================================================================*
 * }}}
 *============================================================================*/
 
/*============================================================================*
 * {{{ Tests
 *============================================================================*/
 
func Test_sql_can_create_instance(T *testing.T) {
	if !TEST_RUN { return; }
	var i = New.Sql(TEST_URI)
  var a = tests.New.Assert(T)
  a.NotNil(i, "Unable to create instance")
  i.Close()
}

func Test_sql_driver_tests(T *testing.T) {
  var assert = tests.New.Assert(T)
	var sql_tests = sqltests.New.SqlTests()
	var runner = &sqlTestRunner{}
	var result = sql_tests.Run(runner, T)
	assert.True(result, "Sql driver tests failed")
}

/*============================================================================*
 * }}}
 *============================================================================*/
 
/*============================================================================*
 * {{{ Helpers
 *============================================================================*/

type sqlTestRunner struct {
}

func (self *sqlTestRunner) Run() bool {
	return TEST_RUN
}

func (self *sqlTestRunner) Table() string {
	return "mysql_tests"
}

func (self *sqlTestRunner) Setup(T *testing.T) (tests.Assert, nsql.Sql) {
  var assert = tests.New.Assert(T)
  
  var instance = New.Sql(TEST_ROOT_URI)
  instance.Raw("CREATE DATABASE IF NOT EXISTS testing")
  instance.Close()
  
  instance = New.Sql(TEST_URI)
  instance.Raw("DROP TABLE IF EXISTS mysql_tests")
  instance.Raw(`CREATE TABLE mysql_tests (
  	id INT NOT NULL AUTO_INCREMENT PRIMARY KEY, 
  	string_value VARCHAR(100), 
  	int_value INT, 
  	long_value INT, 
  	double_value DOUBLE, 
  	bool_value BOOL,
  	datetime_value DATETIME,
  	text_value TEXT
  )`)
  
  return assert, instance
}

func (self *sqlTestRunner) Teardown(instance nsql.Sql) {
  instance.Raw("DROP TABLE `mysql_tests`")
	instance.Close()
}

/*============================================================================*
 * }}}
 *============================================================================*/
 
