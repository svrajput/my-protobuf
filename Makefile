# module name
GO_MODULE := my-protobuf

.PHONY: tidy 
tidy : 	
	go mod tidy

.PHONY: clean
clean : 
	rm -rf ./protogen 

.PHONY: protoc
protoc: 
	protoc --go_opt=module=${GO_MODULE} --go_out=. ./proto/basic/*.proto

.PHONY: build
build: clean protoc tidy

.PHONY: run
run: go run main.go