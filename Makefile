SVR_ROOT_PATH = $(realpath .)
WEB_ROOT_PATH = $(realpath ../../../../dart/mdm)

install_deps:
	go install github.com/favadi/protoc-go-inject-tag@latest

gen_idl: model/idl
	-rm -rf model/vo
	protoc -I=model/idl/vo --go_out=model model/idl/vo/task.proto
	protoc-go-inject-tag -input=model/vo/*.go

gen_web: $(WEB_ROOT_PATH)
	-rm -rf $(SVR_ROOT_PATH)/static
	cd $(WEB_ROOT_PATH) && flutter build web && mv build/web $(SVR_ROOT_PATH)/static