.PHONY: run build start-bin clean help docker-build docker-run docker-stop

include .env
export

.DEFAULT_GOAL := help

APP_NAME=aurelia-reads
IMAGE_NAME=$(APP_NAME)-backend-image

run:
	@go run main.go

build:
	@go build -o $(APP_NAME) main.go

start-bin:
	@./$(APP_NAME)

clean:
	@-rm -rf $(APP_NAME)

docker-up:
	@sudo docker compose up -d

docker-down:
	@sudo docker compose down

help:
	@echo "Доступные цели:"
	@echo "  run         - Собрать и запустить"
	@echo "  build       - Собрать и сохранить"
	@echo "  start-bin   - Запустить бинарник"
	@echo "  clean       - Удалить собранный бинарник"
	@echo "  docker-up   - Поднять все контейнеры через docker-compose"
	@echo "  docker-down - Остановить контейнеры"
	@echo "  help        - Все доступные цели"