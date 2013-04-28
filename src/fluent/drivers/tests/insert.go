package tests

import ns "fluent"
import "time"
import "testing"

/*============================================================================*
 * {{{ Tests
 *============================================================================*/

func (self *sqlTests) Test_can_insert_records(runner SqlTestRunner, T *testing.T) {
	if !runner.Run() { return; }
  var a, i = runner.Setup(T)

	var r1 = map[string]interface{} {
	  "string_value" : "value",
		"int_value" : 100,
		"long_value" : 21312323,
		"double_value" : 324234.23,
		"bool_value" : true,
		"datetime_value" : time.Now(),
	}
	
	var r2 = map[string]interface{} {
	  "string_value" : "valuer2",
		"int_value" : 1010,
		"long_value" : 913323,
		"double_value" : 624234.23,
		"bool_value" : false,
		"datetime_value" : time.Now(),
	}
	
	var insert_query = ns.New.Query(ns.INSERT, runner.Table())
	insert_query.Select("id")
	insert_query.Row(r1)
	insert_query.Row(r2)
	
	var output, err = i.Execute(insert_query)
		
  a.NotNil(output, "No response from insert query")
  a.Nil(err, "Error on record insert")
  
  for i := 0; i < 2; i++ {
    var raw = output.Next()
  	var id = raw["id"].(int64)
  	a.True(id >= 0, "Invalid id returned")
  	a.Log("- inserted id was: %d", id)
  }
  output.Close()
  _ = output

  runner.Teardown(i)
}

/*============================================================================*
 * }}}
 *============================================================================*/
