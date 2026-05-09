package main

import (
	"AureliaReadsBackend/api"
	"AureliaReadsBackend/data/repository"
	"AureliaReadsBackend/data/services"
	"AureliaReadsBackend/domain/usecases"
	"log"
	"os"
)

func main() {
	jwtSecret := os.Getenv("JWT_SECRET")

	jwtService := services.NewJwtService(jwtSecret)
	hashService := services.NewHashService()

	userRepository := repository.NewUserRepository()
	articlesRepository := repository.NewArticlesRepository()

	signUpUseCase := usecases.NewSignUpUseCase(
		userRepository,
		hashService,
		jwtService,
	)

	signInUseCase := usecases.NewSignInUseCase(
		userRepository,
		hashService,
		jwtService,
	)

	r := api.MainRoute(
		signUpUseCase,
		signInUseCase,
		articlesRepository,
		jwtService,
	)

	if err := r.Run(); err != nil {
		log.Fatalf("failed to run server: %v", err)
	}
}
