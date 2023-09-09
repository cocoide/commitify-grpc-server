.PHONY: dev build down proto mock

dev:
	docker-compose up dev
build:
	docker-compose up --build -d app 
down:
	docker-compose down --rmi all
proto:
	cd proto && protoc --go_out=../pkg/pb --go_opt=paths=source_relative \
	--go-grpc_out=../pkg/pb --go-grpc_opt=paths=source_relative \
	commit_message_service.proto
mock:
	go generate ./...