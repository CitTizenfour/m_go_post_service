CURRENT_DIR=$(shell pwd)

APP=$(shell basename ${CURRENT_DIR})

APP_CMD_DIR=${CURRENT_DIR}/cmd

REGISTRY=${REGISTRY}
TAG=latest
ENV_TAG=latest
PROJECT_NAME=${PROJECT_NAME}

copy-proto-module:
	rm -rf ${CURRENT_DIR}/protos
	rsync -rv --exclude=.git ${CURRENT_DIR}/m_protos/* ${CURRENT_DIR}/protos

rm-proto-omit-empty:
	chmod 744 ./scripts/rm_omit_empty.sh && ./scripts/rm_omit_empty.sh ${CURRENT_DIR}

build:
	CGO_ENABLED=0 GOOS=linux go build -mod=vendor -a -installsuffix cgo -o ${CURRENT_DIR}/bin/${APP} ${APP_CMD_DIR}/main.go

proto-gen:
	./scripts/gen-proto.sh  ${CURRENT_DIR}

pull-proto-module:
	git submodule update --init --recursive

update-proto-module:
	git submodule update --remote --merge

migration:
		migrate create -ext sql -dir ./migrations -seq -digits 6 ${name}


clear:
	rm -rf ${CURRENT_DIR}/bin/*

network:
	docker network create --driver=bridge ${NETWORK_NAME}

migrate-up:
	docker run --mount type=bind,source="${CURRENT_DIR}/migrations,target=/migrations" --network ${NETWORK_NAME} migrate/migrate \
		-path=/migrations/ -database=postgres://${POSTGRES_USER}:${POSTGRES_PASSWORD}@${POSTGRES_HOST}:${POSTGRES_PORT}/${POSTGRES_DB}?sslmode=disable up

migrate-down:
	docker run --mount type=bind,source="${CURRENT_DIR}/migrations,target=/migrations" --network ${NETWORK_NAME} migrate/migrate \
		-path=/migrations/ -database=postgres://${POSTGRES_USER}:${POSTGRES_PASSWORD}@${POSTGRES_HOST}:${POSTGRES_PORT}/${POSTGRES_DB}?sslmode=disable down

mark-as-production-image:
	docker tag ${REGISTRY}/${APP}:${TAG} ${REGISTRY}/${APP}:production
	docker push ${REGISTRY}/${APP}:production

build-image:
	docker build --rm -t ${REGISTRY}/${PROJECT_NAME}/${APP}:${TAG} .
	docker tag ${REGISTRY}/${PROJECT_NAME}/${APP}:${TAG} ${REGISTRY}/${PROJECT_NAME}/${APP}:${ENV_TAG}

push-image:
	docker push ${REGISTRY}/${PROJECT_NAME}/${APP}:${TAG}
	docker push ${REGISTRY}/${PROJECT_NAME}/${APP}:${ENV_TAG}

swag_init:
	swag init -g api/main.go -o api/docs

.PHONY: proto

migration-up:
	migrate -path ./migrations -database 'postgres://rakhimjon:147ajt369@0.0.0.0:5432/m_go_post_service?sslmode=disable' up

run:
	go run cmd/main.go
