.PHONY: get test

gen:
	protoc --proto_path services/api/internal/grpc --go_out=plugins=grpc:services/api/internal/grpc calendar.proto
test:
	set -e;\
	docker-compose -f docker-compose.test.yml up -d;\
	docker-compose -f docker-compose.test.yml run integration_test bash -c "cd ./integration-test && go test --mod=vendor";\
	docker-compose -f docker-compose.test.yml down;
