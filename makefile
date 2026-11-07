.PHONY: run build clean help docker-up docker-down migrate-up migrate-down migrate-up1 migrate-down1 migrate-create migrate-version

include .env
export

.DEFAULT_GOAL := help

APP_NAME=aurelia-reads-backend
IMAGE_NAME=$(APP_NAME)-image

POSTGRES_URL=postgres://${POSTGRES_USER}:${POSTGRES_PASSWORD}@${POSTGRES_HOST}:${POSTGRES_PORT}/${POSTGRES_DB}?sslmode=disable

run:
	@go run cmd/main.go

build:
	@go build -o bin/$(APP_NAME) cmd/main.go

clean:
	@-rm -rf bin/$(APP_NAME)

docker-up:
	@sudo docker compose up -d

docker-down:
	@sudo docker compose down

migrate-create:
	@if [ -z "$(NAME)" ]; then \
		echo "Usage: make migrate-create NAME=name_new_version_database"; \
		exit 1; \
	fi

	@migrate create -ext sql -dir migrations -seq $(NAME)

migrate-version:
	@migrate -path migrations -database "$(POSTGRES_URL)" version

migrate-up:
	@migrate -path migrations -database "$(POSTGRES_URL)" up

migrate-down:
	@migrate -path migrations -database "$(POSTGRES_URL)" down

migrate-up1:
	@migrate -path migrations -database "$(POSTGRES_URL)" up 1

migrate-down1:
	@migrate -path migrations -database "$(POSTGRES_URL)" down 1

help:
	@echo "Доступные цели:"
	@echo "  run             - Собрать и запустить"
	@echo "  build           - Собрать и сохранить"
	@echo "  clean           - Удалить собранный бинарник"
	@echo "  docker-up       - Поднять все контейнеры через docker-compose"
	@echo "  docker-down     - Остановить контейнеры"
	@echo "  migrate-create  - Создать новую версию базы данных"
	@echo "  migrate-version - Посмотреть текущую версию базы данных"
	@echo "  migrate-up      - Мигрировать базу данных до последней версии"
	@echo "  migrate-down    - Откатить базу данных до нуля"
	@echo "  migrate-up1     - Мигрировать базу данных на следующию версию"
	@echo "  migrate-down1   - Откатить базу данных к предыдущей версии"
	@echo "  help            - Все доступные цели"