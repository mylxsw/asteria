package misc_test

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/mylxsw/asteria/misc"
	"github.com/stretchr/testify/assert"
)

func TestCallGraph(t *testing.T) {
	cg := misc.CallGraph(1)

	fmt.Println(cg)

	assert.Equal(t, "github.com/mylxsw/asteria/misc_test", cg.PackageName)
	assert.Equal(t, "TestCallGraph", cg.FuncName)
	assert.Regexp(t, regexp.MustCompile("misc/callgraph_test\\.go$"), cg.FileName)
}
