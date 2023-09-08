package main

import (
  "fmt"
  "log"
  "net/http"

  "github.com/go-fed/activitypub"
  "github.com/unmsmfisi-socialapplication/social_app/internal/multipublication/application"
)

func main() {
  // Load the multipublication configuration.
  config, err := config.LoadMultipublicationConfig()
  if err != nil {
    log.Fatal(err)
  }

  // Create a new ActivityPub client.
  client := activitypub.NewClient(nil)

  // Create a new multipublication application.
  app := application.NewMultipublicationApplication(config, client)

  // Register the multipublication handler.
  http.HandleFunc("/multipublication", app.Handle)

  // Start the server.
  log.Fatal(http.ListenAndServe(":8