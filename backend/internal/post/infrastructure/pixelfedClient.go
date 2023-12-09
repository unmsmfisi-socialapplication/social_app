package infrastructure

import (
    "log"
    "github.com/unmsmfisi-socialapplication/social_app/internal/post/domain"
)

type pixelfedClient struct {
}

func NewPixelfedClient() domain.PixelfedAPI {
    return &pixelfedClient{}
}

func (p *pixelfedClient) MultipostPixelfeed(post domain.PostCreate) error {
      var err error = nil

    if err != nil {
        log.Printf("Ocurrió un error al publicar en Pixelfed: %v", err)
        return err
    }
    
    return nil
}

