package debug

import (
	"fmt"
	"os"
	"runtime"
)

func Catch() {
	if err := recover(); err != nil {
		fmt.Println(err)
	}
}

func CatchAndGoexit() {
	if err := recover(); err != nil {
		fmt.Println(err)
		runtime.Goexit()
	}
}

func CatchAndExit() {
	if err := recover(); err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
}

func PanicError(err error) {
	if err != nil {
		_, file, line, ok := runtime.Caller(1)
		if ok {
			panic(fmt.Errorf("%s: line %d:\n%v\n", file, line, err))
		} else {
			panic(err)
		}
	}
}

func PanicAssist(ok bool) {
	if !ok {
		_, file, line, ok := runtime.Caller(1)
		if ok {
			panic(fmt.Errorf("%s: line %d: Assist failed\n", file, line))
		} else {
			panic("Assist failed\n")
		}
	}
}

func PrintError(err error) bool {
	if err != nil {
		_, file, line, ok := runtime.Caller(1)
		if ok {
			fmt.Printf("%s: line %d:\n%v", file, line, err)
		} else {
			fmt.Print(err)
		}
		return true
	}
	return false
}

func PrintlnError(err error) bool {
	if err != nil {
		_, file, line, ok := runtime.Caller(1)
		if ok {
			fmt.Printf("%s: line %d:\n%v\n", file, line, err)
		} else {
			fmt.Println(err)
		}
		return true
	}
	return false
}
