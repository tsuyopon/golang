all:
	mkdir protobuffer
	protoc --go_out=./protobuffer/ sample.proto 
	go build sample.go

# go言語のインストールやGOPATHの設定は事前に必要です。
setup:
	go get -u github.com/golang/protobuf/protoc
	go get -u github.com/golang/protobuf/protoc-gen-go

clean:
	rm -r protobuffer sample
