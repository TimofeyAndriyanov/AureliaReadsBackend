package main

import (
	"database/sql"
	"log/slog"
	"os"

	"github.com/TimofeyAndriyanov/AureliaReadsBackend/data/database"
	articlesRepositoryImpl "github.com/TimofeyAndriyanov/AureliaReadsBackend/data/repository/articles"
	userRepositoryImpl "github.com/TimofeyAndriyanov/AureliaReadsBackend/data/repository/user"
	argon2IdServiceImpl "github.com/TimofeyAndriyanov/AureliaReadsBackend/data/services/argon2"
	sha512ServiceImpl "github.com/TimofeyAndriyanov/AureliaReadsBackend/data/services/sha512"
	authUseCases "github.com/TimofeyAndriyanov/AureliaReadsBackend/domain/usecases/auth"
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

	db, _ := database.NewDatabase(
		postgresHost,
		postgresPort,
		postgresUser,
		postgresPassword,
		postgresDb,
		logger,
	)

	defer func(db *sql.DB) {
		if db.Close() != nil {
			return
		}
	}(db)

	_ = sha512ServiceImpl.NewSHA512HashService()
	argon2IdHashService := argon2IdServiceImpl.NewArgon2IdHashService()

	userRepository := userRepositoryImpl.NewUserRepository(db)
	_ = articlesRepositoryImpl.NewArticlesRepository(db)

	_ = authUseCases.NewSignUpUseCase(
		userRepository,
		argon2IdHashService,
	)

	_ = authUseCases.NewSignInUseCase(
		userRepository,
		argon2IdHashService,
	)
}
