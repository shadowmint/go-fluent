package tests

import "reflect"
import "testing"
import "fmt"

/*============================================================================*
 * {{{ Public api 
 *============================================================================*/

type Assert interface {
  
  // Truth assert functions
  True(value bool, message string, args ...interface{})
  False(value bool, message string, args ...interface{})
  NotNil(value interface{}, message string, args ...interface{})
  Nil(value interface{}, message string, args ...interface{})
  Equals(value interface{}, expected interface{}, message string, args ...interface{})

  // Display a debug message
  Log(message string, args ...interface{})
}

func newAssert(t *testing.T) Assert {
  var rtn = assert { T : t }
  return &rtn
}

/*============================================================================*
 * }}} 
 *============================================================================*/

/*============================================================================*
 * {{{ Implementation
 *============================================================================*/

type assert struct {
  T *testing.T
}

func (self *assert) Log(message string, args ...interface{}) {
	var msg = message
	if len(args) > 0 {
	  msg = fmt.Sprintf(message, args...)
  }
  fmt.Printf("%s\n", msg)
}

func (self *assert) True(value bool, message string, args ...interface{}) {
  if !value {
    var msg = fmt.Sprintf(message, args...)
    msg = fmt.Sprintf("%s (expected TRUE)\n", msg)
    panic(msg)
  }
}

func (self *assert) False(value bool, message string, args ...interface{}) {
  if value {
    var msg = fmt.Sprintf(message, args...)
    msg = fmt.Sprintf("%s (expected FALSE)\n", msg)
    panic(msg)
  }
}

func (self *assert) Equals(value interface{}, expected interface{}, message string, args ...interface{}) {
  if value != expected {
    var msg = fmt.Sprintf(message, args...)
    msg = fmt.Sprintf("%s: '%+v' != '%+v'\n", msg, value, expected)
    panic(msg)
  }
}

func (self *assert) Nil(value interface{}, message string, args ...interface{}) {
  if !self.isNil(value) {
    var msg = fmt.Sprintf(message, args...)
    msg = fmt.Sprintf("%s: '%s' != nil\n", msg, value)
    panic(msg)
  }
}

func (self *assert) NotNil(value interface{}, message string, args ...interface{}) {
  if self.isNil(value) {
    var msg = fmt.Sprintf(message, args...)
    msg = fmt.Sprintf("%s: value(%s) == nil\n", msg, value)
    panic(msg)
  }
}

/*============================================================================*
 * }}} 
 *============================================================================*/

/*============================================================================*
 * {{{ Internal functions 
 *============================================================================*/
 
// Determin 'nil-ness' of a value
func (self *assert) isNil(value interface{}) bool {
	var rtn = true
	var v = reflect.ValueOf(value)
	var vi = reflect.Indirect(v)
	
	// Direct test
	if value == nil {
		rtn = true
	
	// If indirect cannot resolve to a sub-item, this is a value, not nil.
	} else if v == vi {
		rtn = false
		
	// Otherwise, it's a pointer, check for nil
	} else {
		rtn = v.IsNil()
  }
  
	return rtn
}
 
/*============================================================================*
 * }}} 
 *============================================================================*/
