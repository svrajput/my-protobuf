package basic

import (
	"log"
	"testCode/Projects/my-protobuf/protogen/basic"
)

func BasicHello() {
	h := basic.Hello{
		Name: "Sunil Rajput",
	}

	log.Println(&h)
}
