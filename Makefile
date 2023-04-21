ifeq ($(POSTGRES_SETUP_TEST),)
	POSTGRES_SETUP_TEST := user=test password=test dbname=test host=localhost port=5432 sslmode=disable
endif

PKG_PATH=$(CURDIR)/pkg
MIGRATION_FOLDER=$(PKG_PATH)/db/pg/migrations

.PHONY: migration-create
migration-create:
	goose -dir "$(MIGRATION_FOLDER)" create "$(name)" sql

.PHONY: test-migration-up
test-migration-up:
	goose -dir "$(MIGRATION_FOLDER)" postgres "$(POSTGRES_SETUP_TEST)" up

.PHONY: test-migration-down
test-migration-down:
	goose -dir "$(MIGRATION_FOLDER)" postgres "$(POSTGRES_SETUP_TEST)" down


.PHONY: compose-up
compose-up:
	docker-compose build
	docker-compose up -d postgres
	docker-compose up -d cache

.PHONY: compose-rm
compose-rm:
	docker-compose down

.PHONY: test
test:
	docker-compose build
	docker-compose up -d postgres
	docker-compose up -d cache
	goose -dir "$(MIGRATION_FOLDER)" postgres "$(POSTGRES_SETUP_TEST)" up
	go test -v -tags '!integration' ./...
	go test -v -tags=integration ./test
	goose -dir "$(MIGRATION_FOLDER)" postgres "$(POSTGRES_SETUP_TEST)" down
	docker-compose down



