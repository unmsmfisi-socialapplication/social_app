// file of use

package usecase

import (
  "context"
  "errors"
  "github.com/go-fed/activitypub"
  "github.com/unmsmfisi-socialapplication/social_app/backend/internal/multipublication/application"
)

// MultipublicationApplication is the application for multipublication.
type MultipublicationApplication struct {
  config *config.MultipublicationConfig
  client activitypub.Client
}

// NewMultipublicationApplication creates a new MultipublicationApplication.
func NewMultipublicationApplication(config *config.MultipublicationConfig, client activitypub.Client) *MultipublicationApplication {
  return &MultipublicationApplication{
    config: config,
    client: client,
  }
}

// Publish publishes a publication on the configured social networks.
func (app *MultipublicationApplication) Publish(ctx context.Context, text string) error {
  // Get the authorization tokens for the configured social networks.
  tokens := []string{}
  for _, network := range app.config.Networks {
    token, err := app.client.GetToken(network)
    if err != nil {
      return err
    }

    tokens = append(tokens, token)
  }

  // Publish the publication on each social network.
  for _, token := range tokens {
    // Create a new publication.
    publication := activitypub.NewActivityPubObject(activitypub.ObjectTypeActivity)
    publication.Id = activitypub.NewID()
    publication.Actor = activitypub.NewActor(app.client.Me())
    publication.Object = activitypub.NewLink(publication.Id)
    publication.Published = time.Now()
    publication.Content = activitypub.NewString(text)

    // Publish the publication.
    err := app.client.Post(token, publication)
    if err != nil {
      return err
    }
  }

  return nil
}
