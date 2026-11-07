package main

func main() {
	//port := os.Getenv("PORT")

	//postgresHost := os.Getenv("POSTGRES_HOST")
	//postgresPort := os.Getenv("POSTGRES_PORT")
	//postgresUser := os.Getenv("POSTGRES_USER")
	//postgresPassword := os.Getenv("POSTGRES_PASSWORD")
	//postgresDb := os.Getenv("POSTGRES_DB")

	//logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	//db, dbErr := database.NewDatabase(
	//	postgresHost,
	//	postgresPort,
	//	postgresUser,
	//	postgresPassword,
	//	postgresDb,
	//	logger,
	//)

	//if dbErr != nil {
	//	log.Fatal(dbErr)
	//	return
	//}

	//defer func(db *sql.DB) {
	//	if closeErr := db.Close(); closeErr != nil {
	//		log.Fatal(closeErr)
	//		return
	//	}
	//}(db)

	//sha512Service := sha512ServiceImpl.NewSHA512HashService()
	//argon2IdHashService := argon2IdServiceImpl.NewArgon2IdHashService()

	//userRepository := userRepositoryImpl.NewUserRepository(db)
	//sessionRepository := userRepositoryImpl.NewSessionRepository(db)

	//articlesRepo := articlesRepositoryImpl.NewArticlesRepository(db)

	//signUpUseCase := authUseCases.NewSignUpUseCase(
	//	userRepository,
	//	sessionRepository,
	//	argon2IdHashService,
	//	sha512Service,
	//)

	//signInUseCase := authUseCases.NewSignInUseCase(
	//	userRepository,
	//	sessionRepository,
	//	argon2IdHashService,
	//	sha512Service,
	//)

	//rifySessionTokenUseCase := authUseCases.NewVerifySessionTokenUseCase(
	//	sha512Service,
	//	sessionRepository,
	//

	//createArticleUseCase := articleUseCases.NewCreateArticleUseCase(articlesRepo)
	//getAllArticlesUseCase := articleUseCases.NewGetAllArticlesUseCase(articlesRepo)
	//findArticleByIdUseCase := articleUseCases.NewFindArticleByIdUseCase(articlesRepo)
	//deleteArticleUseCase := articleUseCases.NewDeleteArticleUseCase(articlesRepo)

	//authHandler := routers.NewAuthHandler(
	//	signInUseCase,
	//	signUpUseCase,
	//	logger,
	//)

	//articleHandler := routers.NewArticleHandler(
	//	createArticleUseCase,
	//	getAllArticlesUseCase,
	//	findArticleByIdUseCase,
	//	deleteArticleUseCase,
	//	logger,
	//)

	//authMiddlewareHandler := middleware.NewAuthMiddlewareHandler(verifySessionTokenUseCase)

	//api.NewServer(
	//	port,
	//	authHandler,
	//	articleHandler,
	//	authMiddlewareHandler,
	//)
}
