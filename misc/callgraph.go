package misc

import (
	"runtime"
	"strings"
)

type CallGraphInfo struct {
	PackageName string
	FileName    string
	Line        int
}

// CallGraph https://stackoverflow.com/a/25265493/2429469
func CallGraph(skip int) CallGraphInfo {
	pc, f, line, _ := runtime.Caller(skip)

	segs := strings.Split(runtime.FuncForPC(pc).Name(), "/")
	lastSegs := strings.Split(segs[len(segs)-1], ".")

	packageName := strings.Join(append(segs[:len(segs)-1], lastSegs[0]), "/")

	return CallGraphInfo{
		PackageName: packageName,
		FileName:    f,
		Line:        line,
	}
}
