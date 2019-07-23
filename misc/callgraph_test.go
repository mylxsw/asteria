package misc_test

import (
	"regexp"
	"sync"
	"testing"

	"github.com/mylxsw/asteria/misc"
	"github.com/stretchr/testify/assert"
)

func TestCallGraph(t *testing.T) {
	cg := misc.CallGraph(1)

	assert.Equal(t, "github.com/mylxsw/asteria/misc_test", cg.PackageName)
	assert.Regexp(t, regexp.MustCompile("misc/callgraph_test\\.go$"), cg.FileName)


	func() {

		var wg sync.WaitGroup
		wg.Add(1)
		go func() {
			defer wg.Done()

			cg := misc.CallGraph(1)

			assert.Equal(t, "github.com/mylxsw/asteria/misc_test", cg.PackageName)
			assert.Regexp(t, regexp.MustCompile("misc/callgraph_test\\.go$"), cg.FileName)
		}()

		wg.Wait()
	}()

}
