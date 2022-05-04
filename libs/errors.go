package libs

import "log"

func CatchPanic() {
	if err := recover(); err != nil {
		log.Fatalf("Occured panic: [%v]", err)
	}
}
