package assert

import (
	"fmt"
	"path"
	"runtime"
)

func Assert(condition bool) {
	if !condition {
		pc, file, lineNo, ok := runtime.Caller(1)
		if !ok {
			panic("get caller fail")
		}
		funcName := runtime.FuncForPC(pc).Name()
		fileName := path.Base(file)
		panic(fmt.Sprintf("file: %v:%v, func: %v", fileName, lineNo, funcName))
	}
}
