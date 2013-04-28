package utils

import "fmt"
import "bytes"

/*============================================================================*
 * {{{ Public api 
 *============================================================================*/

// A simple string builder wrapper, because the standard one is 
// too fucking hard to remember how to use.
type Str interface {

  // Append to this string (shortcut)
  S(msg string, params ...interface{})

  // Appent to this string
  Print(msg string, params ...interface{})

  // Clear
  Clear()

  // As a string
  String() string
}

/*============================================================================*
 * }}} 
 *============================================================================*/

/*============================================================================*
 * {{{ Implementation
 *============================================================================*/

type str struct {
  buffer *bytes.Buffer
}

func newStr() Str {
  var rtn = &str {
    buffer : &bytes.Buffer{},
  }
  return rtn
}

// Append to this string
func (self *str) S(msg string, params ...interface{}) {
  msg = fmt.Sprintf(msg, params...)
  self.buffer.WriteString(msg)
}

// Append to this string
func (self *str) Print(msg string, params ...interface{}) {
  msg = fmt.Sprintf(msg, params...)
  self.buffer.WriteString(msg)
}

// Clear
func (self *str) Clear() {
  self.buffer = &bytes.Buffer{}
}

// As a string
func (self *str) String() string {
  return self.buffer.String()
}

/*============================================================================*
 * }}} 
 *============================================================================*/
