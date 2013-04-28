package sqlite3

import "fmt"
import nsql "fluent"
import gsql "database/sql"
import "fluent/drivers/utils"

/*============================================================================*
 * {{{ Implementation
 *============================================================================*/
 
type rowset struct {
  utils.Rowset
}

func newRowset(rows *gsql.Rows) nsql.Rowset {
  var rtn = &rowset{*utils.New.Rowset(rows, &schema{})}
  return rtn
}

type schema struct {
}

// Apply the schema and convert types to schema types
func (self *schema) ApplySchema(rows *utils.Rowset, rtn map[string]interface{}) error {
  fmt.Printf("Applying schema")
  return nil
}

/*============================================================================*
 * }}}
 *============================================================================*/
