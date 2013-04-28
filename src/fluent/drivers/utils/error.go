package utils

import "fmt"

/*============================================================================*
 * {{{ Public api 
 *============================================================================*/

// Standard error type
type Error struct {
  
  // Error code associated with this response.
  Code int

  // Error message
  Message string
}

// Generate a standard error
func Fail(code int, msg string, args ...interface{}) *Error {
  return &Error {
    Code : code,
    Message : fmt.Sprintf(msg, args...),
  }
}

/*============================================================================*
 * }}} 
 *============================================================================*/

/*============================================================================*
 * {{{ Error implementation
 *============================================================================*/

func (self *Error) Error() string {
  return fmt.Sprintf("%s (%d)", self.Message, self.Code)
}

/*============================================================================*
 * }}} 
 *============================================================================*/
