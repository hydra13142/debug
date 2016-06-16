package debug

import (
	"fmt"
	"os"
)

func Catch() {
	if err := recover(); err != nil {
		fmt.Println(err)
	}
}

func CatchAndExit() {
	if err := recover(); err != nil {
		fmt.Println(err)
		os.Exit()
	}
}

func PanicError(err error) {
	if err != nil {
		_, file, line, ok := runtime.Caller(1)
		if ok {
			panic(fmt.Errorf("%s: line %d:\n%v\n",file, line, err))
		} else {
			panic(err)
		}
	}
}

func PanicAssist(ok bool) {
	if !ok {
		_, file, line, ok := runtime.Caller(1)
		if ok {
			panic(fmt.Errorf("%s: line %d: Assist failed\n",file, line))
		} else {
			panic("Assist failed\n")
		}
	}
}
