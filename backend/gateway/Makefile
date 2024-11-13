# Создать и запустить контейнеры со сборкой 
DEFAULT_CONFIG_PATH = ./configs/default_config.yaml
ROUTES_PATH = ./configs/routes.yaml

up:
	docker compose up -d

down:
	docker compose down

restart:
	docker compose restart

restart-go:
	docker compose restart golang

dev:
	go run ./cmd/main/main.go \
		--config_path=$(DEFAULT_CONFIG_PATH) \
		--routes_path=$(ROUTES_PATH)

# Применить миграции к базе данных
local-migrate:
	go run ./cmd/migrator/migrator.go --config=./config/config_dev.json
migrate: start
	docker run --rm --network auth_database -v ./:/code -w /code golang:latest go run ./cmd/migrator/migrator.go --config=./config/config_docker_dev.json

#
init-test-venv:
	cd tests \
	&& python3 -m venv venv \
	&& . venv/bin/activate \
	&& pip install -r requirements.txt


# Запустить тестирование проекта. Перед запуском обязательно применить команду make init_db
test: up
	.  tests/venv/bin/activate \
	&& pytest -s

# Сгенерировать описание API по комментариям в коде
swagger:
	docker run --rm -v ./:/code ghcr.io/swaggo/swag:latest init -g ./cmd/main/main.go -o ./api
