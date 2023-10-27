package router

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/unmsmfisi-socialapplication/social_app/internal/comment"
	email "github.com/unmsmfisi-socialapplication/social_app/internal/email_sender"
	"github.com/unmsmfisi-socialapplication/social_app/internal/login/application"
	"github.com/unmsmfisi-socialapplication/social_app/internal/login/infrastructure"
	"github.com/unmsmfisi-socialapplication/social_app/internal/profile"
	"github.com/unmsmfisi-socialapplication/social_app/internal/post"

	registerapplication "github.com/unmsmfisi-socialapplication/social_app/internal/register/application"
	registerinfrastructure "github.com/unmsmfisi-socialapplication/social_app/internal/register/infrastructure"
	"github.com/unmsmfisi-socialapplication/social_app/pkg/database"
)

func Router() http.Handler {

	err := database.InitDatabase()
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}

	dbInstance := database.GetDB()

	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(configCorsMiddleware())

	commentRouter := comment.CommentModuleRouter(dbInstance)

	postRoutes := post.PostModuleRouter(dbInstance)

    profileRouter := profile.ProfileModuleRouter(dbInstance)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("{\"hello\": \"world\"}"))
	})

	r.Get("/slow", func(w http.ResponseWriter, r *http.Request) {
		// Simulates some hard work.
		//
		// We want this handler to complete successfully during a shutdown signal,
		// so consider the work here as some background routine to fetch a long running
		// search query to find as many results as possible, but, instead we cut it short
		// and respond with what we have so far. How a shutdown is handled is entirely
		// up to the developer, as some code blocks are preemptible, and others are not.
		time.Sleep(5 * time.Second)

		w.Write([]byte(fmt.Sprintf("{\"response\": \"all done slow\"}")))
	})


	// Login
	loginRepo := infrastructure.NewUserDBRepository(dbInstance)
	loginUseCase := application.NewLoginUseCase(loginRepo)
	loginHandler := infrastructure.NewLoginHandler(loginUseCase)
	r.Post("/login", loginHandler.HandleLogin)

	// Register
	registerRepo := registerinfrastructure.NewUserRepository(dbInstance)
	registerUseCase := registerapplication.NewRegistrationUseCase(registerRepo)
	registerHandler := registerinfrastructure.NewRegisterUserHandler(registerUseCase)
	r.Post("/register", registerHandler.RegisterUser)

	r.Mount("/comments", commentRouter)

	r.Mount("/post", postRoutes)

	r.Mount("/profile", profileRouter)

	//Email-sender

	emailRouter := email.EmailModuleRouter()
	r.Mount("/email", emailRouter)
	return r
}
