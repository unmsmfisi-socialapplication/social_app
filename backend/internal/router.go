package internal

import (
	// Importa tus paquetes necesarios aquí

	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	"github.com/unmsmfisi-socialapplication/social_app/internal/comment"
	"github.com/unmsmfisi-socialapplication/social_app/internal/login/application"
	"github.com/unmsmfisi-socialapplication/social_app/internal/login/infrastructure"

	interest_topics_application "github.com/unmsmfisi-socialapplication/social_app/internal/interest_topics/application"
	interest_topics_repository "github.com/unmsmfisi-socialapplication/social_app/internal/interest_topics/infraestructure/repository"
	interest_topics_infraestructure "github.com/unmsmfisi-socialapplication/social_app/internal/interest_topics/infraestructure"

	"github.com/unmsmfisi-socialapplication/social_app/pkg/database"
)

func Router() http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)

	corsMiddleware := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	})

	r.Use(corsMiddleware.Handler)

	err := database.InitDatabase()
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}

	dbInstance := database.GetDB()

	commentRouter := comment.CommentModuleRouter(dbInstance)
	r.Mount("/comments", commentRouter)

	dbRepo := infrastructure.NewUserDBRepository(dbInstance)
	loginUseCase := application.NewLoginUseCase(dbRepo)
	loginHandler := infrastructure.NewLoginHandler(loginUseCase)

	//interesTopics
	dbinterestTopics := interest_topics_repository.NewUserInterestsDBRepository(dbInstance)
	selectTopicUseCase := interest_topics_application.NewInterestTopicsUseCase(dbinterestTopics)
	selecTopicHandler := interest_topics_infraestructure.NewSelectTopicHandler(selectTopicUseCase)

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
	r.Post("/login", loginHandler.HandleLogin)
	r.Post("/interestTopics", selecTopicHandler.HandleSelectTopic)
	return r
}
