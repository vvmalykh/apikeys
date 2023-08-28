include .env

SERVICES=docker-compose -f ./docker-compose.yaml

run:
	@echo "------ RUN ------"
	@docker run --rm -t -d -it ${PROJECT_PREFIX}
	@echo "------ SUCCEED ------"

make re-build:
	@echo "------ BUILD ------"
	@$(SERVICES) build app
	@$(SERVICES) up -d --no-deps app

build:
	@make down
	@echo "------ BUILD ------"
	@$(SERVICES) build --no-cache
	@make up
	@make migrate

up:
	@make down
	@echo "------ UP ------"
	@$(SERVICES) up -d -V
	@echo "------ SUCCEED ------"

down:
	@echo "------ DOWN ------"
	@$(SERVICES) down
	@echo "------ SUCCEED ------"

migrate:
	@echo "------ Running Migrations ------"
	@docker-compose run --rm app sh -c 'migrate -path=/db/migrations -database "postgres://username:password@db:5432/app_db?sslmode=disable" up'
	@echo "------ Migrations Complete ------"

migrate-down:
	@echo "------ Running Migrations ------"
	@docker-compose run --rm app sh -c 'migrate -path=/db/migrations -database "postgres://username:password@db:5432/app_db?sslmode=disable" down'
	@echo "------ Migrations Complete ------"

create-migration:
	@echo "------ Creating new migration ------"
	@docker-compose run --rm app sh -c 'migrate create -ext sql -dir /db/migrations -seq $(name)'
	@echo "------ New migration created ------"


app-sh:
	@echo "------ YOU'RE INSIDE $(PROJECT_PREFIX)-app-container ------"
	@docker exec -it app-$(PROJECT_PREFIX)-container sh
