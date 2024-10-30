# Создать и запустить контейнеры со сборкой 
# DEFAULT_CONFIG_PATH = ./configs/default_config.yaml
DEFAULT_CONFIG_PATH = ./configs/default_config.yaml

up:
	docker compose up -d

down:
	docker compose down

restart:
	docker compose restart

dev:
	go run ./cmd/main/main.go \
		--config_path=$(DEFAULT_CONFIG_PATH)
