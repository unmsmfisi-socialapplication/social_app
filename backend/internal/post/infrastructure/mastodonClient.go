package infrastructure

import (
    "log"
    "github.com/unmsmfisi-socialapplication/social_app/internal/post/domain"
)

type mastodonClient struct {
}

func NewMastodonClient() domain.MastodonAPI {
    return &mastodonClient{
    }
}

func (m *mastodonClient) MultipostMastodon(post domain.PostCreate) error {
   
    var err error = nil

    if err != nil {
        log.Printf("Ocurrió un error al publicar en Mastodon: %v", err)
        return err
    }

    return nil
}

