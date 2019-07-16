package asteria_test

import (
	"testing"

	"github.com/mylxsw/asteria"
)

func TestCallGraph(t *testing.T) {
	cg := asteria.CallGraph(1)
	if cg.PackageName != "github.com/mylxsw/asteria_test" {
		t.Error("test failed")
	}

	if cg.FuncName != "TestCallGraph" {
		t.Error("test failed")
	}
}
