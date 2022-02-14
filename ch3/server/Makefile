all: setupTools grpc

grpc:
	protoc proto/order_management.proto \
		--go_opt=module=github.com/channel-io/grpc-study \
		--go_out=. \
		--go-grpc_opt=module=github.com/channel-io/grpc-study \
		--go-grpc_out=.

setupTools:
	cat tools.go | grep _ | awk -F'"' '{print $$2}' | xargs -tI % go install %

clean:
	go mod tidy
