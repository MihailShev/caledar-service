gen:
	protoc --proto_path services/api/internal/grpc --go_out=plugins=grpc:services/api/internal/grpc calendar.proto