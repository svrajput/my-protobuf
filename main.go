package main

import (
	"fmt"
	"log"
	"testCode/Projects/my-protobuf/basic"
	"time"
)

type logWriter struct {
}

func (writer logWriter) Write(bytes []byte) (int, error) {
	return fmt.Print(time.Now().Format("15:10:10") + " " + string(bytes))
}

func main() {

	log.SetFlags(0)
	log.SetOutput(new(logWriter))

	basic.BasicHello()
}
