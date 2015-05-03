// base_test
package commd

import "testing"

func TestCommand(t *testing.T) {
	c := newStringCmd("set")
	c.AddCmd("mykey")
	c.AddCmd("mvalue")
	t.Logf("Logtext:"+c.toString())
}
