package router

import (
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/unmsmfisi-socialapplication/social_app/internal/chat/domain"
	"github.com/unmsmfisi-socialapplication/social_app/internal/chat/infraestructure"
	"github.com/unmsmfisi-socialapplication/social_app/internal/comment"
	email "github.com/unmsmfisi-socialapplication/social_app/internal/email_sender"
	"github.com/unmsmfisi-socialapplication/social_app/internal/login/application"
	"github.com/unmsmfisi-socialapplication/social_app/internal/login/infrastructure"
	"github.com/unmsmfisi-socialapplication/social_app/internal/post"
	"github.com/unmsmfisi-socialapplication/social_app/internal/profile"
	"github.com/unmsmfisi-socialapplication/social_app/internal/register"

	interest_topics "github.com/unmsmfisi-socialapplication/social_app/internal/interest_topics"
	"github.com/unmsmfisi-socialapplication/social_app/pkg/database"

	auth "github.com/unmsmfisi-socialapplication/social_app/internal/auth"
	chat "github.com/unmsmfisi-socialapplication/social_app/internal/chat"

	auth_application "github.com/unmsmfisi-socialapplication/social_app/internal/auth/application"
	auth_infrastructure "github.com/unmsmfisi-socialapplication/social_app/internal/auth/infrastructure"
	follow "github.com/unmsmfisi-socialapplication/social_app/internal/follow"
)

func Router() http.Handler {
	err := database.InitDatabase()
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}

	dbInstance := database.GetDB()

	authRepo := auth_infrastructure.NewAuthSessionDBRepository(dbInstance)
	authUseCase := auth_application.NewAuthUseCase(authRepo)
	tokenMiddleware := auth.AuthModuleRouter(authUseCase)

	freeRoutes := chi.NewRouter()

	freeRoutes.Use(middleware.RequestID)
	freeRoutes.Use(middleware.Logger)
	freeRoutes.Use(configCorsMiddleware())
	//subrouter for protected routes
	protectedRoutes := chi.NewRouter()

	freeRoutes.Mount("/", protectedRoutes)

	// Apply the TokenMiddleware to all routes that are mounted on protectedRoutes subrouter\
	protectedRoutes.Use(tokenMiddleware.Middlewares()...)

	/* IMPORTANT:
	Use freeRoutes to your endpoint if you consider your component is not for a registered user

	Otherwhise, use protectedRoutes when you mount your endpoint.

	*/
	protectedRoutes.Get("/", func(w http.ResponseWriter, freeRoutes *http.Request) {
		w.Write([]byte("{\"hello\": \"world\"}"))
	})

	//AUTH
	authRouter := auth.AuthModuleRouter(authUseCase)
	protectedRoutes.Mount("/auth", authRouter)

	commentRouter := comment.CommentModuleRouter(dbInstance)

	postRoutes := post.PostModuleRouter(dbInstance)

	profileRouter := profile.ProfileModuleRouter(dbInstance)

	// Login
	loginRepo := infrastructure.NewUserDBRepository(dbInstance)
	loginUseCase := application.NewLoginUseCase(loginRepo)
	loginHandler := infrastructure.NewLoginHandler(loginUseCase)
	freeRoutes.Post("/login", loginHandler.HandleLogin)

	//Chat
	hub := domain.NewHub()
	chatHandler := infraestructure.NewHandler(hub)

	go hub.RunChatManager()

	chatRouter := chat.ChatModuleRouter(chatHandler)
	protectedRoutes.Mount("/chat", chatRouter)

	// Register
	registerModule := register.RegisterModule(dbInstance)
	freeRoutes.Mount("/register", registerModule)

	protectedRoutes.Mount("/comments", commentRouter)

	protectedRoutes.Mount("/post", postRoutes)

	protectedRoutes.Mount("/profile", profileRouter)

	//Email-sender
	emailRouter := email.EmailModuleRouter()
	protectedRoutes.Mount("/email", emailRouter)

	//interestTopics
	interestTopicsRouter := interest_topics.InterestTopicsModuleRouter(dbInstance)
	protectedRoutes.Mount("/interestTopics", interestTopicsRouter)

	// Follow Profile
	followRouter := follow.FollowModuleRouter(dbInstance)
	protectedRoutes.Mount("/follow_profile", followRouter)
	return freeRoutes
}
