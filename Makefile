SVR_ROOT_PATH = $(realpath .)
WEB_ROOT_PATH = $(realpath ./web)

install_deps:
	go install github.com/favadi/protoc-go-inject-tag@latest

gen_idl_for_go: $(SVR_ROOT_PATH)/model/idl
	@echo "Updating idl ..."
	@-rm -rf $(SVR_ROOT_PATH)/model/vo
	@protoc -I=$(SVR_ROOT_PATH)/model/idl/vo --go_out=model $(SVR_ROOT_PATH)/model/idl/vo/task.proto
	@protoc-go-inject-tag -input=$(SVR_ROOT_PATH)/model/vo/*.go
	@-rm -rf $(SVR_ROOT_PATH)/dal/db/po
	@-mkdir -p $(SVR_ROOT_PATH)/dal/db/po
	@-rm -rf $(SVR_ROOT_PATH)/dal/db/query
	@-mkdir -p $(SVR_ROOT_PATH)/dal/db/query
	@go run $(SVR_ROOT_PATH)/cmd/main.go
	@echo "Done."

gen_idl_for_dart: $(SVR_ROOT_PATH)/model/idl
	-rm -rf $(WEB_ROOT_PATH)/lib/models/vo
	mkdir -p $(WEB_ROOT_PATH)/lib/models/vo
	protoc -I=$(SVR_ROOT_PATH)/model/idl/vo --dart_out=$(WEB_ROOT_PATH)/lib/models/vo $(SVR_ROOT_PATH)/model/idl/vo/task.proto
	find $(WEB_ROOT_PATH)/lib/models/vo -type f -name "*.dart" -print0 | xargs -0 perl -pi -e 's/\$$pb\.PbList<([^>]+)>\(\)/[] as \$$pb.PbList<\1>/g'

gen_idl: gen_idl_for_go gen_idl_for_dart

gen_web: $(WEB_ROOT_PATH)
	@echo "Updating web ..."
	@-rm -rf $(SVR_ROOT_PATH)/static
	cd $(WEB_ROOT_PATH) && flutter pub get && flutter build web && mv build/web $(SVR_ROOT_PATH)/static
	@echo "Done."
