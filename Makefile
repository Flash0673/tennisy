### CODE GEN

LOCAL_BIN:=$(CURDIR)/bin

.PRONY: bin-deps
bin-deps:
	GOBIN=$(LOCAL_BIN) go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	GOBIN=$(LOCAL_BIN) go install -mod=mod google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	GOBIN=$(LOCAL_BIN) go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@latest
	GOBIN=$(LOCAL_BIN) go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@latest
	GOBIN=$(LOCAL_BIN) go install github.com/xo/dbtpl@latest

.PRONY: generate
generate:
	mkdir -p pb/
	protoc -I=vendor.protogen --proto_path=api \
    	--go_out=pb \
    	--go_opt paths=source_relative \
    	--plugin=protoc-gen-go=bin/protoc-gen-go \
    	--go-grpc_out=pb \
    	--plugin=protoc-gen-go-grpc=bin/protoc-gen-go-grpc \
    	--go-grpc_opt paths=source_relative \
    	--grpc-gateway_out=pb \
    	--grpc-gateway_opt paths=source_relative \
    	--grpc-gateway_opt generate_unbound_methods=true \
    	--plugin=protoc-gen-grpc-gateway=bin/protoc-gen-grpc-gateway \
    	--openapiv2_out=pb --openapiv2_opt logtostderr=true \
    	--plugin=protoc-gen-openapiv2=bin/protoc-gen-openapiv2 \
    	api/auth/v1/auth.proto

.PRONY: vendor-proto
vendor-proto:
		@if [ ! -d vendor.protogen/google ]; then \
			git clone https://github.com/googleapis/googleapis vendor.protogen/googleapis &&\
			mkdir -p  vendor.protogen/google/ &&\
			mv vendor.protogen/googleapis/google/api vendor.protogen/google &&\
			rm -rf vendor.protogen/googleapis ;\
		fi


### DB

DB_HOST=127.0.0.1
DB_NAME=tennisly
DB_USER=postgres
DB_PASS=password
DB_PORT=5432
MIGRATION_FOLDER=./tools/migrations

.PRONY: infra ## поднимает инфрастуктуру для проекта
infra:
	docker-compose -f ./.infra/build/package/docker-compose.yaml up -d --force-recreate --wait

.PRONY: infra-stop ## останавливает контейнеры
infra-stop:
		docker-compose -f ./.infra/build/package/docker-compose.yaml down

# If the first argument is "run"...
ifeq (migration,$(firstword $(MAKECMDGOALS)))
  # use the rest as arguments for "run"
  MIGRATION_ARGS := $(wordlist 2,$(words $(MAKECMDGOALS)),$(MAKECMDGOALS))
  # ...and turn them into do-nothing targets
  $(eval $(MIGRATION_ARGS):;@:)
endif
.PHONY: migration ## проверка миграций на ошибки
migration:
	goose create $(MIGRATION_ARGS) sql --dir ${MIGRATION_FOLDER}

.PHONY: migration-status ## проверка статуса миграций
migration-status:
	goose postgres "user=${DB_USER} password=${DB_PASS} dbname=${DB_NAME} host=${DB_HOST} port=${DB_PORT} sslmode=disable" status -dir ${MIGRATION_FOLDER}

.PHONY: migrations-up ## накатка миграций
migrations-up:
	goose postgres "user=${DB_USER} password=${DB_PASS} dbname=${DB_NAME} host=${DB_HOST} port=${DB_PORT} sslmode=disable" up -dir ${MIGRATION_FOLDER}

.PHONY: migrations-down ## откатка миграций на 1 назад
migrations-down:
	goose postgres "user=${DB_USER} password=${DB_PASS} dbname=${DB_NAME} host=${DB_HOST} port=${DB_PORT} sslmode=disable" down -dir ${MIGRATION_FOLDER}


.PHONY: migrations-reset ## откатка ВСЕХ миграций
migrations-reset:
	goose postgres "user=${DB_USER} password=${DB_PASS} dbname=${DB_NAME} host=${DB_HOST} port=${DB_PORT} sslmode=disable" down-to 0 -dir ${MIGRATION_FOLDER}

XO_OUTPUT_PATH=./internal/xo
XO_TEMPLATE_PATH=./tools/xo_templates

.PHONY: dump-templates ## дамп темплейтов
dump-templates:
	./bin/dbtpl dump $(XO_TEMPLATE_PATH)

.PHONY: xo ## генерация dto базы данных
xo: migrations-up
	rm -r $(XO_OUTPUT_PATH)
	mkdir -p $(XO_OUTPUT_PATH)
	./bin/dbtpl schema -o $(XO_OUTPUT_PATH) --src $(XO_TEMPLATE_PATH) --schema public \
 	"pgsql://$(DB_USER):$(DB_PASS)@localhost:$(DB_PORT)/$(DB_NAME)?sslmode=disable"

	rm $(XO_OUTPUT_PATH)/dbtpl.dbtpl.go
	rm $(XO_OUTPUT_PATH)/goosedbversion.dbtpl.go
