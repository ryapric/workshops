PROTO_PACKAGE := example
PROTO_PACKAGE_VERSION := v1

gen: gen-go gen-python

# Go code can use protoc, like most languages
gen-go:
	@protoc \
		--go_out=../pb \
		--go_opt=paths=source_relative \
		--go-grpc_out=../pb \
		--go-grpc_opt=paths=source_relative \
		--go-grpc_opt=require_unimplemented_servers=false \
		./$(PROTO_PACKAGE)/$(PROTO_PACKAGE_VERSION)/*.proto

# ... but Python seems to be an angry elf, so we have to use a protoc distro
# that's in the grpcio_rools Python module
gen-python:
	@python3 -m grpc_tools.protoc \
		-I=. \
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
