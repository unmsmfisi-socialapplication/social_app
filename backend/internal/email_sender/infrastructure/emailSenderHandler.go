package infrastructure

import (
	"encoding/json"
	"io"
	"net/http"
	"regexp"

	"github.com/unmsmfisi-socialapplication/social_app/pkg/utils"

	"github.com/unmsmfisi-socialapplication/social_app/internal/email_sender/application"
	"github.com/unmsmfisi-socialapplication/social_app/internal/email_sender/domain"
)

const (
	status_err     = "ERROR"
	status_success = "OK"
	err_Payload    = "Invalid request payload"
	err_email      = "Wrong email format"
	err_fields     = "Empty fields"
	err_internal   = "Internal server error"
	succ_sent      = "Email sent successfully"
)

type EmailSenderHandler struct {
	useCase application.EmailSenderUseCaseInterface
}

func NewEmailSenderHandler(useCase application.EmailSenderUseCaseInterface) *EmailSenderHandler {
	return &EmailSenderHandler{
		useCase: useCase,
	}
}

func (esh *EmailSenderHandler) HandleSendEmail(w http.ResponseWriter, r *http.Request) {
	bodyBytes, err := io.ReadAll(r.Body)
	if err != nil {
		utils.SendJSONResponse(w, http.StatusBadRequest, status_err, err_Payload)
		return
	}

	var emailPayload struct {
		To      string `json:"to"`
		Subject string `json:"subject"`
		Body    string `json:"body"`
	}
	err = json.Unmarshal(bodyBytes, &emailPayload)
	if err != nil {
		utils.SendJSONResponse(w, http.StatusBadRequest, status_err, err_Payload)
		return
	}

	var emailRegex = `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	if !regexp.MustCompile(emailRegex).MatchString(emailPayload.To) {
		utils.SendJSONResponse(w, http.StatusBadRequest, status_err, err_email)
		return
	}

	var raw map[string]interface{}
	err = json.Unmarshal(bodyBytes, &raw)
	if err != nil {
		utils.SendJSONResponse(w, http.StatusBadRequest, status_err, err_Payload)
		return
	}

	if len(raw) != 3 || raw["To"] == nil || raw["Subject"] == nil || raw["Body"] == nil {
		utils.SendJSONResponse(w, http.StatusBadRequest, status_err, err_Payload)
		return
	}
	if raw["To"] == "" || raw["Subject"] == "" || raw["Body"] == "" {
		utils.SendJSONResponse(w, http.StatusBadRequest, status_err, err_fields)
		return
	}
	email, err := domain.NewEmail(emailPayload.To, emailPayload.Subject, emailPayload.Body)
	if err != nil {
		utils.SendJSONResponse(w, http.StatusInternalServerError, status_err, err_internal)
		return
	}

	err = esh.useCase.SendEmail(email)
	if err != nil {
		utils.SendJSONResponse(w, http.StatusInternalServerError, status_err, err_internal)
		return
	}

	utils.SendJSONResponse(w, http.StatusOK, status_success, succ_sent)
}
