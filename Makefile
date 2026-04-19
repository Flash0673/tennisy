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
