include .env

.PHONY: proto_gen test-db

proto_gen:
	protoc -I api api/user_service_v1/user_service.proto --go_out=./pkg --go_opt=paths=source_relative --go-grpc_out=./pkg --go-grpc_opt=paths=source_relative

test-db:
	docker run --name=user-service -e POSTGRES_PASSWORD=$(POSTGRES_PASSWORD) -e POSTGRES_DB=$(POSTGRES_DB) -p 5432:5432 -d postgres