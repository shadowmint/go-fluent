package fluent

/*============================================================================*
 * {{{ Public api
 *============================================================================*/

// Generic SQL representation
type Query struct {

	// Query type constant
	Type int
	
	// The target table
	Table string
	
	// The requested output columns
	Requested []string
	
	// The constraints on the data
	Constraints Constraint
	
	// If this is an insert or update operation, the data 
	Rows []map[string]interface{}
}

// Select a particular column to read in the result set
func (self *Query) Select(column string) *Query {
	var size = len(self.Requested) + 1
	var tmp = make([]string, size, size)
	copy(tmp, self.Requested)
	self.Requested = tmp
	self.Requested[size - 1] = column
	return self
}

// Add a row to this query 
func (self *Query) Row(row map[string] interface{}) *Query {
	var size = len(self.Rows) + 1
	var tmp = make([]map[string]interface{}, size, size)
	copy(tmp, self.Rows)
	self.Rows = tmp
	self.Rows[size - 1] = row
	return self
}

// Set the where constraint on this query
func (self *Query) Where(where string, values ...interface{}) *Query {
	self.Constraints.Where = where
	self.Constraints.Values = values
	return self
}

// Set the limit constraint on this query
func (self *Query) Offset(offset int) *Query {
	self.Constraints.Offset = offset
	return self
}

// Set the offset constraint on this query
func (self *Query) Limit(count int) *Query {
	self.Constraints.Limit = count
	return self
}

/*============================================================================*
 * }}}
 *============================================================================*/
 
/*============================================================================*
 * {{{ Internal api
 *============================================================================*/
 
// Return a new query for the given table and operation
func newQuery(action int, table string) *Query {
	var rtn = &Query {
		Type : action,
		Table : table,
		Requested : []string{},
		Rows : []map[string]interface{} {},
		Constraints : Constraint {
			Where : "",
			Values : []interface{} {},
			Limit : -1,
			Offset : -1,
		},
	}
	return rtn
}
 
/*============================================================================*
 * }}}
 *============================================================================*/
