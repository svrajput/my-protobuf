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
		Gender:   basic.Gender_GENDER_MALE,
		Emails:   []string{"abc@test.com", "back@test.com"},
	}

	log.Println(&u)
}
