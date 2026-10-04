package main

import (
	"log/slog"
	"os"

	"github.com/TimofeyAndriyanov/AureliaReadsBackend/data/database"
	articlesRepositoryImpl "github.com/TimofeyAndriyanov/AureliaReadsBackend/data/repository/articles"
	userRepositoryImpl "github.com/TimofeyAndriyanov/AureliaReadsBackend/data/repository/user"
	argon2IdServiceImpl "github.com/TimofeyAndriyanov/AureliaReadsBackend/data/services/argon2"
	sha512ServiceImpl "github.com/TimofeyAndriyanov/AureliaReadsBackend/data/services/sha512"
	"github.com/TimofeyAndriyanov/AureliaReadsBackend/domain/usecases"
)

func main() {
	_ = os.Getenv("HOST")
	_ = os.Getenv("PORT")

	postgresHost := os.Getenv("POSTGRES_HOST")
	postgresPort := os.Getenv("POSTGRES_PORT")
	postgresUser := os.Getenv("POSTGRES_USER")
	postgresPassword := os.Getenv("POSTGRES_PASSWORD")
	postgresDb := os.Getenv("POSTGRES_DB")

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	_, _ = database.NewDatabase(
		postgresHost,
		postgresPort,
		postgresUser,
		postgresPassword,
		postgresDb,
		logger,
	)

	_ = sha512ServiceImpl.NewSHA512HashService()
	argon2IdHashService := argon2IdServiceImpl.NewArgon2IdHashService()

	userRepository := userRepositoryImpl.NewUserRepository()
	_ = articlesRepositoryImpl.NewArticlesRepository()

	_ = usecases.NewSignUpUseCase(
		userRepository,
		argon2IdHashService,
	)

	_ = usecases.NewSignInUseCase(
		userRepository,
		argon2IdHashService,
	)
}
