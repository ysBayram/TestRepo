package utils

import (
	"fmt"
	"net/http"
	"runtime"

	"github.com/gin-gonic/gin"

	"github.com/sirupsen/logrus"
)

func GetStackFrame(beginningLine, lineCount int) (stackFrame string) {
	pcList := make([]uintptr, lineCount)

	runtime.Callers(beginningLine, pcList)
	for i, pc := range pcList {
		f := runtime.FuncForPC(pc)

		if f != nil {
			file, line := f.FileLine(pc)
			stackFrame += fmt.Sprintf("Level:%d, File:%s, Line:%d, Function:%s\n", i, file, line, f.Name())
		} else {
			stackFrame += fmt.Sprintf("Can not infer stackFrame for level:%d\n", i)
		}
	}

	return
}

func ErrorFuncForDefer(ctx *gin.Context) {
	if r := recover(); r != nil {
		logrus.Errorf("Error Detail : \r\n %v \r\n StackFrame : \r\n %v ", r, GetStackFrame(6, 6))
		ctx.JSON(http.StatusInternalServerError, r)
	}
}
