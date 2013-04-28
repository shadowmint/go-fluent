package fluent

/*============================================================================*
 * {{{ Public api
 *============================================================================*/

// A generic interface for working with tables
type Sql interface {

	// Execute a query and return a rowset
	Execute(query *Query) (Rowset, error)
	
	// Execute a raw SQL query for anything not supported, eg. CREATE TABLE, etc.
	Raw(query string, values  ...[]interface{}) (Rowset, error)
	
	// Close the connection
	Close() 
}

/*============================================================================*
 * }}}
 *============================================================================*/
