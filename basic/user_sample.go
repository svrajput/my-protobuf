package basic

import (
	"log"
	"testCode/Projects/my-protobuf/protogen/basic"
)

// function which rceives information.
func BasicUser() {
	u := basic.User{
		Id:       10,
		Username: "sunil_rajput",
		IsActive: true,
		Password: []byte("testpasswd"),
	}

	log.Println(&u)
}
