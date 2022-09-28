gen: gen-go gen-python

# Go code can use protoc, like most languages
gen-go:
	@protoc \
		--go_out=../pb \
		--go_opt=paths=source_relative \
		--go-grpc_out=../pb \
		--go-grpc_opt=paths=source_relative \
		--go-grpc_opt=require_unimplemented_servers=false \
		./example/v1/example.proto

# ... but Python seems to be an angry elf, so we have to use a protoc distro
# that's in the grpcio_rools Python module
gen-python:
	@python3 -m grpc_tools.protoc \
		-I=. \
		--python_out=../pb \
		--grpc_python_out=../pb \
		./example/v1/example.proto
