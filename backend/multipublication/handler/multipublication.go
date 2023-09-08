// file of controller

package handler

import (
  "context"
  "fmt"
  "net/http"

  "github.com/go-fed/activitypub"
  "github.com/unmsmfisi-socialapplication/social_app/backend/internal/multipublication/application"
)

// MultipublicationHandler handles requests for multipublication.
type MultipublicationHandler struct {
  app *application.MultipublicationApplication
}

// NewMultipublicationHandler creates a new MultipublicationHandler.
func NewMultipublicationHandler(app *application.MultipublicationApplication) *MultipublicationHandler {
  return &MultipublicationHandler{app}
}

// Handle handles a request for multipublication.
func (h *MultipublicationHandler) Handle(w http.ResponseWriter, r *http.Request) {
  // Get the publication data.
  data := struct {
    Text string `json:"text"`
  }{}

  err := json.NewDecoder(r.Body).Decode(&data)
  if err != nil {
    fmt.Println(err)
    return
  }

  // Publish the publication on the configured social networks.
  h.app.Publish(context.Background(), data.Text)

  // Return a successful response.
  w.WriteHeader(http.StatusOK)
}
