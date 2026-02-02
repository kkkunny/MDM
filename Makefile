gen_idl:
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	rm -rf model/vo
	protoc -I=model/idl/vo --go_out=model model/idl/vo/task.proto