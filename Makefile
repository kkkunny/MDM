SVR_ROOT_PATH = $(realpath .)
WEB_ROOT_PATH = $(realpath ./web)

install_deps:
	go install github.com/favadi/protoc-go-inject-tag@latest

gen_idl: model/idl
	@echo "Updating idl ..."
	@-rm -rf model/vo
	@protoc -I=model/idl/vo --go_out=model model/idl/vo/task.proto
	@protoc-go-inject-tag -input=model/vo/*.go
	@-rm -rf dal/db/po
	@-mdkir -p dal/db/po
	@-rm -rf dal/db/query
	@-mdkir -p dal/db/query
	@go run cmd/main.go
	@echo "Done."

gen_web: $(WEB_ROOT_PATH)
	@echo "Updating web ..."
	@-rm -rf $(SVR_ROOT_PATH)/static
	cd $(WEB_ROOT_PATH) && flutter pub get && flutter build web && mv build/web $(SVR_ROOT_PATH)/static
	@echo "Done."
