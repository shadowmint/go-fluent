package tests

import ns "fluent"
import "time"
import "testing"

/*============================================================================*
 * {{{ Tests
 *============================================================================*/

func (self *sqlTests) Test_can_select_records(runner SqlTestRunner, T *testing.T) {
	if !runner.Run() { return; }
  var a, i = runner.Setup(T)

	var r1 = map[string]interface{} {
	  "string_value" : "value",
		"int_value" : 0,
		"long_value" : 21312323,
		"double_value" : 324234.23,
		"bool_value" : true,
		"datetime_value" : time.Now(),
	}
	
	var insert_query = ns.New.Query(ns.INSERT, runner.Table())
	for i := 0; i < 15; i++ {
		var row = map[string]interface{} {}
		for k, v := range r1 {
	    row[k] = v
    }
		insert_query.Row(row)
	}
	
	var _, err = i.Execute(insert_query)
  a.Nil(err, "Error on record insert")
  
  var select_query = ns.New.Query(ns.SELECT, runner.Table()). 
    Select("int_value").
    Select("string_value").
    Where("int_value > ? AND int_value < ?", 0, 5).
    Limit(5)

  var result, serr = i.Execute(select_query)
  
  a.Nil(serr, "Error on record select")
  // TODO: Check result
  _ = result

  runner.Teardown(i)
}

/*============================================================================*
 * }}}
 *============================================================================*/
