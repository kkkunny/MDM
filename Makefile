ROOT_PATH = $(realpath .)
WEB_ROOT_PATH = $(realpath ../../../../dart/mdm)

gen_idl: model/idl
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	rm -rf model/vo
	protoc -I=model/idl/vo --go_out=model model/idl/vo/task.proto

gen_web: $(WEB_ROOT_PATH)
	-rm -rf $(ROOT_PATH)/static
	cd $(WEB_ROOT_PATH) && flutter build web && mv build/web $(ROOT_PATH)/static