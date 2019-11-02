.PHONY: gen test start monitoring

gen:
	protoc --proto_path pkg/grpc --go_out=plugins=grpc:pkg/grpc calendar.proto
test:
	set -e;\
	docker-compose -f docker-compose.test.yml up -d;\
	docker-compose -f docker-compose.test.yml run integration_test bash -c "cd ./integration-test && go test --mod=vendor";\
	docker-compose -f docker-compose.test.yml down;

start:
	docker-compose up

monitoring:
	docker-compose -f docker-compose.monitoring.yml up

event:
	go run services/api/client/client.go