package utils

import "fmt"
import nsql "fluent"
import gsql "database/sql"
import "container/list"

/*============================================================================*
 * {{{ Public api
 *============================================================================*/

// Interface for applying a distinct data schema
type RowsetSchema interface {
  ApplySchema(rows *Rowset, values map[string]interface{}) error
}

// Define the schema to access row data with
func (self *Rowset) Map(name string, value int) nsql.Rowset {
  self.Schema[name] = value
  return self
}

// Get the next row, return nil when none left
func (self *Rowset) Next() map[string]interface{} {
	var rtn map[string]interface{} = nil
	if self.Rows != nil {
		var cols, cerr = self.Rows.Columns()
		if cerr != nil {
			fmt.Printf("Failed to query column list on db rows: %s", cerr.Error())
		} else if (len(cols) == 0) {
			fmt.Printf("Failed to query column list on db rows: No columns in result set")
    } else {
			if self.Rows.Next() {
        rtn = Scan(self.Rows)
        self.Worker.ApplySchema(self, rtn)
			}
		}
    return rtn
	}

  // No rows? Ok! Return a value
  var tmp = self.Values.Front()
  if tmp != nil {
    self.Values.Remove(tmp)
    rtn = tmp.Value.(map[string]interface{})
    self.Worker.ApplySchema(self, rtn)
  }

  return rtn
}

// Close this Rowset and discard it
func (self *Rowset) Close() {
	if self.Rows != nil {
		self.Rows.Close()
	}
}

// Merge another set of values into this Rowset
func (self *Rowset) MergeValues(data *Rowset, evaluate bool) {
  if !evaluate {
    self.Rows = data.Rows
  } else {
    var rows = data.Rows
    if rows != nil {

      // Handle new loaded records
      for rows.Next() {
        var record = map[string]interface{} {}
        record = Scan(rows)
        self.Values.PushBack(record)
      }
      rows.Close()
      self.Rows = nil
    } else {
      
      // Handle existing loaded records
      var record = data.Values.Front()
      for record != nil {
        var r = record.Value.(map[string]interface{})
        if len(r) > 0 {
          self.Values.PushBack(r)
        }
        record = record.Next()
      }
    }
  }
}

type Rowset struct {

	// For rows that actually come from the database
	Rows *gsql.Rows
	
	// For manually constructed row sets
	Values list.List

  // Schema
  Schema map[string]int 

  // Schema worker
  Worker RowsetSchema
}

/*============================================================================*
 * }}}
 *============================================================================*/
 
/*============================================================================*
 * {{{ Implementation
 *============================================================================*/
 
func newRowset(rows *gsql.Rows, schema RowsetSchema) *Rowset {
	return &Rowset{
		Rows : rows,
		Values : list.List{},
    Schema : map[string]int{},
    Worker : schema,
	}
}

/*============================================================================*
 * }}}
 *============================================================================*/
