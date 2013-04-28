go fluent
==

This is a basic proof of concept for the idea of a fluent database
api for code. 

See the fluent/drivers/tests folder for test examples.

Example code
--

    import ns "fluent"
    import "time"
    import "testing"

      var r1 = map[string]interface{} {
        "string_value" : "value",
        "int_value" : 0,
        "long_value" : 21312323,
        "double_value" : 324234.23,
        "bool_value" : true,
        "datetime_value" : time.Now(),
      }
      
      var insert_query = ns.New.Query(ns.INSERT, "TargetTable")
      for i := 0; i < 15; i++ {
        var row = map[string]interface{} {}
        for k, v := range r1 {
          row[k] = v
        }
        insert_query.Row(row)
      }
      
      var _, err = i.Execute(insert_query)
      
      var select_query = ns.New.Query(ns.SELECT, runner.Table()). 
        Select("int_value").
        Select("string_value").
        Where("int_value > ? AND int_value < ?", 0, 5).
        Limit(5)

      var result, serr = i.Execute(select_query)
