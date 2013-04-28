package utils

import "testing"
import "fluent/drivers/tests"

/*============================================================================*
 * {{{ Tests
 *============================================================================*/

func Test_str_can_create_instance(T *testing.T) {
  var a = tests.New.Assert(T)
  var i = setup_str()

  a.NotNil(i, "Failed to create instance")

  teardown_str(i)
}

func Test_str_can_print(T *testing.T) {
  var a = tests.New.Assert(T)
  var i = setup_str()

  i.Print("Hello")
  i.S("World")
  var out = i.String()
  a.Equals(out, "HelloWorld", "Print test failed")

  i.Clear()
  i.S("%s %d %s", "Hello", 10, "World")
  i.S(" ")
  i.S("%s %d %s", "Hello", 10, "World")
  out = i.String()
  a.Equals(out, "Hello 10 World Hello 10 World", "Printf test failed")
  
  teardown_str(i)
}

/*============================================================================*
 * }}}
 *============================================================================*/

/*============================================================================*
 * {{{ Tests
 *============================================================================*/

func setup_str() Str {
  var rtn = newStr()
  return rtn
}

func teardown_str(i Str) {
}

/*============================================================================*
 * }}}
 *============================================================================*/
