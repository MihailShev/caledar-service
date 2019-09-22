gen:
	protoc --proto_path internal/grpc/ --go_out=plugins=grpc:internal/grpc calendar.proto