include .env
export

service-run:
	@go run main.go

docker-build:
	@sudo docker build . -t aurelia-reads-backend-image

docker-run:
	@sudo docker run aurelia-reads-backend-image