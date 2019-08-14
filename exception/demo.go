package main

import (
	"errors"
	"fmt"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
)

func formatError(formating string, args ...interface{}) error {
	filename, line, funcname := "???", 0, "???"
	pc, filename, line, ok := runtime.Caller(2)
	if ok {
		funcname = runtime.FuncForPC(pc).Name()      // main.(*MyStruct).foo
		funcname = filepath.Ext(funcname)            // .foo
		funcname = strings.TrimPrefix(funcname, ".") // foo
		filename = filepath.Base(filename)           // /full/path/basename.go => basename.go
	}
	argMsg := ""
	if len(formating) > 0 {
		argMsg = fmt.Sprintf(formating, args...)
	}
	return fmt.Errorf("%s:%d:%s: %s", filename, line, funcname, argMsg)
}

func FormatError(cause error, formating string, args ...interface{}) error {
	if cause == nil {
		return formatError(formating, args...)
	}
	msg := formatError(formating, args...).Error()
	msg += "\ncaused by " + cause.Error()
	return errors.New(msg)
}

func doStep(i int) error {
	fmt.Println("do step " + strconv.Itoa(i))

	if i == 1 {
		return FormatError(nil, "%d", i)
	}
	return nil
}

func workflow(name string) error {
	fmt.Println("starting workflow: " + name)
	for i := 0; i < 3; i++ {
		err := doStep(i)
		if err != nil {
			return FormatError(err, "%s", name)
		}
	}
	fmt.Println("ending workflow: " + name)
	return nil
}

func program() error {
	err := workflow("design")
	if err != nil {
		return FormatError(err, "%s %s", "hello", "world")
	}
	workflow("execute")
	if err != nil {
		return fmt.Errorf("program: %v", err)
	}
	return nil
}

func main() {
	err := program()
	if err != nil {
		fmt.Printf("%v", err)
	}
}
