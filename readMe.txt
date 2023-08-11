
packageOne - has go package details. 
proto      - has PB definition files. 
protogen   - pb generated go files. file extension will be  filename.pb.go 
root filder- files and modules  

## Install protolint https://github.com/yoheimuta/protolint 
brew tap yoheimuta/protolint
brew install protolint


# install compiler 
brew install protobuf 

# install go lang lib 
brew install protoc-gen-go


# generate hello.proto file using following command - It will create go code under 
 cmd : protoc --go_out=. ./proto/basic/*.proto 