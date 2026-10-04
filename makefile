include .env
export

IMAGE_NAME := aurelia-reads-backend-image

service-run:
	@go run main.go

docker-build:
	@sudo docker build . -t $(IMAGE_NAME)

docker-run:
	@sudo docker run -p ${PORT}:${PORT} $(IMAGE_NAME)