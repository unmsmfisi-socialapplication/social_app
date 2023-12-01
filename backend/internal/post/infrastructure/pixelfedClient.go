package infrastructure

import (
    "github.com/unmsmfisi-socialapplication/social_app/internal/post/domain"
)

type pixelfedClient struct {
}

func NewPixelfedClient(/* parámetros de configuración */) domain.PixelfedAPI {
    return &pixelfedClient{
    }
}

func (p *pixelfedClient) MultipostPixelfeed(post domain.PostCreate) error {

    return nil
}
