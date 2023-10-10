package email

import (
	"github.com/go-chi/chi"
	"github.com/unmsmfisi-socialapplication/social_app/internal/email_sender/application"
	"github.com/unmsmfisi-socialapplication/social_app/internal/email_sender/infrastructure"
)

func EmailModuleRouter() *chi.Mux {
	r := chi.NewRouter()
	auth := infrastructure.LoginAuth("contact.appsocial@gmail.com", "hyus qffk zoow hzzj")

	realEmailSender := &infrastructure.RealEmailSender{}
	emailSenderRepository := infrastructure.NewEmailSenderRepository(auth, realEmailSender)
	emailSenderUseCase := application.NewEmailSenderUseCase(emailSenderRepository)
	emailSenderHandler := infrastructure.NewEmailSenderHandler(emailSenderUseCase)

	r.Post("/send", emailSenderHandler.HandleSendEmail)

	return r
}
