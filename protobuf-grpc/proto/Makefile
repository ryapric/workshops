PROTO_PACKAGE := example
PROTO_PACKAGE_VERSION := v1

gen: gen-go gen-python

proto_deps = https://github.com/googleapis/googleapis \
						 https://github.com/google/protobuf

proto_deps_import_flags = -I=. \
													-I=/tmp/googleapis \
													-I=/tmp/protobuf/src

get-deps:
	@for dep in $(proto_deps) ; do \
		printf 'Getting/checking for proto dep "%s"...\n' "$${dep}" && \
		git clone --depth=1 "$${dep}.git" /tmp/"$$(basename $${dep})" > /dev/null 2>&1 || true ; \
	done

# Go code can use protoc, like most languages
gen-go: get-deps
	@protoc \
		$(proto_deps_import_flags) \
		--go_out=../pb \
		--go_opt=paths=source_relative \
		--go-grpc_out=../pb \
		--go-grpc_opt=paths=source_relative \
		--go-grpc_opt=require_unimplemented_servers=false \
		./$(PROTO_PACKAGE)/$(PROTO_PACKAGE_VERSION)/*.proto

# ... but Python seems to be an angry elf, so we have to use a protoc distro
# that's in the grpcio_rools Python module
gen-python: get-deps
	@python3 -m grpc_tools.protoc \
		$(proto_deps_import_flags) \
		--python_out=../pb \
		--grpc_python_out=../pb \
		./$(PROTO_PACKAGE)/$(PROTO_PACKAGE_VERSION)/*.proto
# AND ALSO we need to modify an import line in the generated code, or else that
# code's module pathing won't work as-generated. It would work if the proto
# file(s) AND the generated output were in the same directory, but we don't have
# our tree set up that way
	@sed -i \
		's/from $(PROTO_PACKAGE).$(PROTO_PACKAGE_VERSION)/from ./g' \
		../pb/$(PROTO_PACKAGE)/$(PROTO_PACKAGE_VERSION)/$(PROTO_PACKAGE)_pb2_grpc.py
