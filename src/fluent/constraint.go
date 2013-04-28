package fluent

/*============================================================================*
 * {{{ Public api
 *============================================================================*/

// Query constraints
type Constraint struct {

	// The native query
	Where string
	
	// An ordered set of values to bind to the where query
	Values []interface{}
	
	// Limit for the query
	Limit int
	
	// Offset for the query
	Offset int
}

/*============================================================================*
 * }}}
 *============================================================================*/
