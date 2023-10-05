.PHONY: dev build down proto mock test

dev:
	docker-compose up dev
build:
	docker-compose up --build -d app 
down:
	docker-compose down --rmi all
proto:
	cd proto && protoc --go_out=../pkg/grpc --go_opt=paths=source_relative \
	--go-grpc_out=../pkg/grpc --go-grpc_opt=paths=source_relative \
	separate_commit_service.proto
mock:
	go generate ./...
test:
	go test ./pkg/usecase/.