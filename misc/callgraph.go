package misc

import (
	"runtime"
	"strings"
)

type CallGraphInfo struct {
	PackageName string
	FuncName    string
	FileName    string
	Line        int
}

// CallGraph https://stackoverflow.com/a/25265493/2429469
func CallGraph(skip int) CallGraphInfo {
	pc, f, line, _ := runtime.Caller(skip)
	parts := strings.Split(runtime.FuncForPC(pc).Name(), ".")
	pl := len(parts)
	packageName := ""
	funcName := parts[pl-1]

	if parts[pl-2][0] == '(' {
		funcName = parts[pl-2] + "." + funcName
		packageName = strings.Join(parts[0:pl-2], ".")
	} else {
		packageName = strings.Join(parts[0:pl-1], ".")
	}

	return CallGraphInfo{
		PackageName: packageName,
		FileName:    f,
		Line:        line,
		FuncName:    funcName,
	}
}
