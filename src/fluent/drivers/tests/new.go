package tests

import "testing"

/*============================================================================*
 * {{{ Internal api 
 *============================================================================*/

type factory struct {
	SqlTests func() SqlTests
  Assert func(t *testing.T) Assert
}

var New factory = factory {
	SqlTests : newSqlTests,
  Assert : newAssert,
}

/*============================================================================*
 * }}}
 *============================================================================*/
