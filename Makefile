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

up:
	@make down
	@echo "------ UP ------"
	@$(SERVICES) up -d -V
	@echo "------ SUCCEED ------"

down:
	@echo "------ DOWN ------"
	@$(SERVICES) down
	@echo "------ SUCCEED ------"

app-sh:
	@echo "------ YOU'RE INSIDE $(PROJECT_PREFIX)-app-container ------"
	@docker exec -it app-$(PROJECT_PREFIX)-container sh
