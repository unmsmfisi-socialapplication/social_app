package domain

import (
	"time"
)

type Post struct {
	Id            int64
	UserId        int64
	Title         string
	Description   string
	HasMultimedia bool
	Public        bool
	Multimedia    string
	InsertionDate time.Time
	UpdateDate    time.Time
}

type MastodonContent struct {
    Text       string `json:"text"`       // Text of the toot in Mastodon
    Visibility string `json:"visibility"` // Visibility of the toot (public, private, unlisted, etc.)
    // Other Mastodon-specific fields
}

type PixelfedContent struct {
    Caption   string   `json:"caption"`  // Caption in Pixelfed
    ImageURLs []string `json:"imageUrls"` // URLs of images in Pixelfed
    Tags      []string `json:"tags"`      // Tags in Pixelfed
    // Other Pixelfed-specific fields
}

type CreatePost struct {
    UserId        int64  `json:"userId"`
    Title         string `json:"title" db:"title" validate:"max=100"`
    Description   string `json:"description" db:"description" validate:"max=1000"`
    HasMultimedia bool   `json:"hasMultimedia"`
    Public        bool   `json:"public"`
    Multimedia    string `json:"multimedia" db:"multimedia" validate:"max=1000"`

    Mastodon *MastodonContent `json:"mastodon,omitempty"`
    Pixelfed *PixelfedContent `json:"pixelfed,omitempty"`
}

