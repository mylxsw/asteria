package misc_test

import (
	"testing"

	"github.com/mylxsw/asteria/misc"
)

func TestCallGraph(t *testing.T) {
	cg := misc.CallGraph(1)
	if cg.PackageName != "github.com/mylxsw/asteria/misc_test" {
		t.Error("test failed")
	}

	if cg.FuncName != "TestCallGraph" {
		t.Error("test failed")
	}
}
