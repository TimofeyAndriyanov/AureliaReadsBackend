package main

import (
	"database/sql"
	"log"
	"log/slog"
	"os"

	"github.com/TimofeyAndriyanov/AureliaReadsBackend/api"
	"github.com/TimofeyAndriyanov/AureliaReadsBackend/api/routers"
	"github.com/TimofeyAndriyanov/AureliaReadsBackend/data/database"
	articlesRepositoryImpl "github.com/TimofeyAndriyanov/AureliaReadsBackend/data/repository/articles"
	userRepositoryImpl "github.com/TimofeyAndriyanov/AureliaReadsBackend/data/repository/user"
	argon2IdServiceImpl "github.com/TimofeyAndriyanov/AureliaReadsBackend/data/services/argon2"
	sha512ServiceImpl "github.com/TimofeyAndriyanov/AureliaReadsBackend/data/services/sha512"
	articleUseCases "github.com/TimofeyAndriyanov/AureliaReadsBackend/domain/usecases/article"
	authUseCases "github.com/TimofeyAndriyanov/AureliaReadsBackend/domain/usecases/auth"
)

func main() {
	host := os.Getenv("HOST")
	port := os.Getenv("PORT")

	postgresHost := os.Getenv("POSTGRES_HOST")
	postgresPort := os.Getenv("POSTGRES_PORT")
	postgresUser := os.Getenv("POSTGRES_USER")
	postgresPassword := os.Getenv("POSTGRES_PASSWORD")
	postgresDb := os.Getenv("POSTGRES_DB")

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	db, dbErr := database.NewDatabase(
		postgresHost,
		postgresPort,
		postgresUser,
		postgresPassword,
		postgresDb,
		logger,
	)

	if dbErr != nil {
		log.Fatal(dbErr)
		return
	}

	defer func(db *sql.DB) {
		if closeErr := db.Close(); closeErr != nil {
			log.Fatal(closeErr)
			return
		}
	}(db)

	sha512Service := sha512ServiceImpl.NewSHA512HashService()
	argon2IdHashService := argon2IdServiceImpl.NewArgon2IdHashService()

	userRepository := userRepositoryImpl.NewUserRepository(db)
	sessionRepository := userRepositoryImpl.NewSessionRepository(db)

	articlesRepo := articlesRepositoryImpl.NewArticlesRepository(db)

	signUpUseCase := authUseCases.NewSignUpUseCase(
		userRepository,
		sessionRepository,
		argon2IdHashService,
		sha512Service,
	)

	signInUseCase := authUseCases.NewSignInUseCase(
		userRepository,
		sessionRepository,
		argon2IdHashService,
		sha512Service,
	)

	createArticleUseCase := articleUseCases.NewCreateArticleUseCase(articlesRepo)
	getAllArticlesUseCase := articleUseCases.NewGetAllArticlesUseCase(articlesRepo)
	findArticleByIdUseCase := articleUseCases.NewFindArticleByIdUseCase(articlesRepo)
	deleteArticleUseCase := articleUseCases.NewDeleteArticleUseCase(articlesRepo)

	authHandler := routers.NewAuthHandler(
		signInUseCase,
		signUpUseCase,
		logger,
	)

	articleHandler := routers.NewArticleHandler(
		createArticleUseCase,
		getAllArticlesUseCase,
		findArticleByIdUseCase,
		deleteArticleUseCase,
		logger,
	)

	api.NewServer(
		host,
		port,
		authHandler,
		articleHandler,
	)
}
