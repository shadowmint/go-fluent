package utils

import "fmt"
import "reflect"
import "database/sql"

/*============================================================================*
 * {{{ Public api
 *============================================================================*/

// The sql api isn't very friendly; this function reads a single row from
// the database and converts it into native driver types.
//
// You may want to extend this to provide complex type mapping like 
// DATETIME for sqlite, etc.
//
// Pass in a rows() that already has had next called on it; this function
// does not invoke Next().
//
// If unable to process the given rowset, the return is nil.
func Scan(rows *sql.Rows) map[string]interface{} {
	var rtn map[string]interface{} = nil
	if rows != nil {
		var cols, cerr = rows.Columns()
		if cerr != nil {
			fmt.Printf("Failed to query column list on db rows: %s", cerr.Error())
		} else if (len(cols) == 0) {
			fmt.Printf("Failed to query column list on db rows: No columns in result set")
    } else {
      rtn = scan(cols, rows)
		}
	} else {
    fmt.Printf("Invalid row set: nil")
  }
  return rtn
}
	
/*============================================================================*
 * }}}
 *============================================================================*/
 
/*============================================================================*
 * {{{ Implementation
 *============================================================================*/
 
// Scan values from the database
func scan(cols []string, rows *sql.Rows) map[string]interface{} {

  var rtn = map[string]interface{} {}

  // Create scanner array
  var values []interface{} 
  var generic = reflect.TypeOf(values).Elem()
  for i := 0; i < len(cols); i++ {
    values = append(values, reflect.New(generic).Interface())
  }
  
  // Scan!
  var err = rows.Scan(values...)
  if err != nil {
    fmt.Printf("Driver failed to scan values: %s", err.Error())
    return nil;
  }

  // Convert into native types
  for i := 0; i < len(cols); i++ {
    var raw_value = *(values[i].(*interface{}))
    var raw_type = reflect.TypeOf(raw_value)
    switch {
      case raw_type == reflect.TypeOf(int64(0)):
        rtn[cols[i]] = raw_value.(int64)
    }
  }

  return rtn;
}
 
/*============================================================================*
 * }}}
 *============================================================================*/
