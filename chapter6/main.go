package main

import (
	_ "net/http/pprof"
	"os"
	"runtime/trace"
)

func main() {
	trace.Start(os.Stderr)
	defer trace.Stop()

	ch := make(chan string)
	go func() {
		ch <- "assasasasas"
	}()
	<-ch
}



