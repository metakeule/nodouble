package nodouble

import (
	"fmt"
	"runtime"
)

type info struct {
	version string
	file    string
	line    int
}

var pkgs = map[string]info{}

func AddPackage(path, version string) {
	_, file, line, _ := runtime.Caller(0)
	i, has := pkgs[path]
	if has {
		panic(fmt.Sprintf("can't load package %s version %s as requested in %s:%v; already loaded version %s as requested from %s:%v",
			path, version, file, line,
			i.version, i.file, i.line,
		))
	}
	pkgs[path] = info{version, file, line}
}
