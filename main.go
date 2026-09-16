package main

import (
	"errors"
	"fmt"
	"net/http"
	"os"

	articlesRepositoryImpl "github.com/TimofeyAndriyanov/AureliaReadsBackend/data/repository/articles"
	userRepositoryImpl "github.com/TimofeyAndriyanov/AureliaReadsBackend/data/repository/user"
	argon2IdServiceImpl "github.com/TimofeyAndriyanov/AureliaReadsBackend/data/services/argon2"
	sha512ServiceImpl "github.com/TimofeyAndriyanov/AureliaReadsBackend/data/services/sha512"
	"github.com/TimofeyAndriyanov/AureliaReadsBackend/domain/usecases"
)

func main() {
	host := os.Getenv("HOST")

	_ = os.Getenv("POSTGRES_HOST")
	_ = os.Getenv("POSTGRES_USER")
	_ = os.Getenv("POSTGRES_USER")
	_ = os.Getenv("POSTGRES_PASSWORD")
	_ = os.Getenv("POSTGRES_DB")

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

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World"))
	})

	err := http.ListenAndServe(host, nil)

	if errors.Is(err, http.ErrServerClosed) {
		fmt.Println("server stop")
	} else {
		fmt.Println(err)
	}

}
