package main

import (
	"fmt"
	"log"
	"runtime"
	"runtime/debug"
)

func runtime_mod() {
	bi, ok := debug.ReadBuildInfo()
	if !ok {
		log.Fatal("Error reading build info")
	}

	fmt.Println(bi)

	fmt.Println("\nNumCPU:", runtime.NumCPU())

	debug.PrintStack()
	fmt.Println("")

	pc, file, line, ok := runtime.Caller(0)
	fmt.Println(pc, file, line, ok)

	f := runtime.FuncForPC(pc)
	fmt.Println(f.Entry(), f.Name())

	/* pcs := make([]uintptr, 0, 16)
	fmt.Println(runtime.Callers(0, pcs))
	frames := runtime.CallersFrames(pcs)
	for {
		frame, more := frames.Next()
		fmt.Println(frame.Func.Name())
		if !more {
			break
		}
	} */
}
