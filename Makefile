CURRENT_DIR=$(shell pwd)

-include .env

.PHONY: run
run:
	go run main.go

.PHONY: tidy
tidy:
	go mod tidy 

.PHONY: vendor
vendor:
	go mod vendor

.PHONY: lint
lint:
	golangci-lint run

.PHONY: sqlc
sqlc:
	sqlc generate

.PHONY: migrate_up
migrate_up:
	goose -dir db/migration/migrations postgres ${PSQL_URI} up

.PHONY: migrate_down
migrate_down:
	goose -dir db/migration/migrations postgres ${PSQL_URI} down

.PHONY: sync_db
sync_db:
	pg_dump ${PSQL_URI} --schema-only > db/migration/schema/1_schema.up.sql

.PHONY: proto
proto:
	rm -f generated/**/*.go
	rm -f doc/swagger/*.swagger.json
	mkdir -p generated
	protoc \
		--proto_path=protos --go_out=generated --go_opt=paths=source_relative \
		--go-grpc_out=generated --go-grpc_opt=paths=source_relative \
		--grpc-gateway_out=generated --grpc-gateway_opt=paths=source_relative \
		--openapiv2_out=doc/swagger --openapiv2_opt=allow_merge=true,merge_file_name=swagger_docs,use_allof_for_refs=true,disable_service_tags=true,json_names_for_fields=false \
		--validate_out="lang=go,paths=source_relative:generated" \
			protos/**/*.proto



##########################################################################
# Server
server-copy:
	mkdir -p server
	cp ./Makefile ./server
	cp -r ./config ./server
	cp ./.env ./server
	cp ./Dockerfile ./server
	cp ./docker-compose.yml ./server
	ssh ${SERVER_USERNAME}@${SERVER_IP} "mkdir -p ${PROJECT_NAME}"
	scp -r ./server/.env ${SERVER_USERNAME}@${SERVER_IP}:~/${PROJECT_NAME}/
	scp -r ./server/* ${SERVER_USERNAME}@${SERVER_IP}:~/${PROJECT_NAME}/

##########################################################################
# Image
push-test:
	make server-copy
	sudo docker build -t ${DOCKERHUB_REPOSITORY}/${PROJECT_NAME}:${DOCKERHUB_TAG} .
	docker push ${DOCKERHUB_REPOSITORY}/${PROJECT_NAME}:${DOCKERHUB_TAG} 
	docker rmi ${DOCKERHUB_REPOSITORY}/${PROJECT_NAME}:${DOCKERHUB_TAG}
	ssh ${SERVER_USERNAME}@${SERVER_IP} "mkdir -p ${PROJECT_NAME}"
	scp -r ./server/* ${SERVER_USERNAME}@${SERVER_IP}:~/${PROJECT_NAME}
	ssh ${SERVER_USERNAME}@${SERVER_IP} "cd ~/${PROJECT_NAME} && sudo -S docker compose down && sudo -S docker compose pull && sudo -S docker compose up -d"
