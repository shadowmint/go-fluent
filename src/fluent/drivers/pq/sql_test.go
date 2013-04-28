package pq

import nsql "fluent"
import sqltests "fluent/drivers/tests"
import "fluent/drivers/tests"
import "testing"
import "fmt"

/*============================================================================*
 * {{{ Test constants
 *============================================================================*/
 
// The URI to connect to a testing database
const TEST_ROOT_URI = "user=root dbname=template1 sslmode=disable"
const TEST_URI = "user=root dbname=testing sslmode=disable"

// Disable tests if no instance is available.
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
	return "pq_tests"
}

func (self *sqlTestRunner) Setup(T *testing.T) (tests.Assert, nsql.Sql) {
  var assert = tests.New.Assert(T)
  
  var instance = New.Sql(TEST_ROOT_URI)
  var _, err = instance.Raw("CREATE DATABASE testing")
  if err != nil {
    // fmt.Printf("Failed to create database: %s", err.Error())
  }
  instance.Close()
  
  instance = New.Sql(TEST_URI)
  instance.Raw("DROP TABLE IF EXISTS " + self.Table())
  _, err = instance.Raw("CREATE TABLE " + self.Table() + ` (
  	id SERIAL,
  	string_value VARCHAR(100), 
  	int_value INT, 
  	long_value INT, 
  	double_value DOUBLE PRECISION, 
  	bool_value BOOL,
  	datetime_value TIMESTAMP,
  	text_value TEXT,
    PRIMARY KEY (id)
  )`)
  if err != nil {
    fmt.Printf("Failed to create table: %s", err.Error())
  }
  
  return assert, instance
}

func (self *sqlTestRunner) Teardown(instance nsql.Sql) {
  instance.Raw("DROP TABLE " + self.Table())
	instance.Close()
}

/*============================================================================*
 * }}}
 *============================================================================*/
 
