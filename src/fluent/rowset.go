package fluent

/*============================================================================*
 * {{{ Public api
 *============================================================================*/

// Interface for working with rows
type Rowset interface {

  // Define a column type mapping for Next() to use
  // Note: You MUST do this for Next() will return an empty map!
  Map(name string, tvalue int) Rowset

	// Get the next row, return nil when none left
	Next() map[string]interface{}
	
	// Close this rowset and discard it
	Close()
}

/*============================================================================*
 * }}}
 *============================================================================*/
