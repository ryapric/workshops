gen: gen-go gen-python

gen-go:
	@protoc \
		--go_out=. \
		--go_opt=paths=source_relative \
		--go-grpc_out=. \
		--go-grpc_opt=paths=source_relative \
		pb/example.proto

gen-python:
	@python3 -m grpc_tools.protoc \
		-I=. \
		--python_out=. \
		--grpc_python_out=. \
		pb/example.proto
