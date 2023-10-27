package dto

import (
	"bytes"
	"io"
	"testing"

	"github.com/go-ap/activitypub"
)

func NewProfileRequestTest(t *testing.T) {
	t.Log("NewProfileRequestTest")

	data := []byte(`{
      "@context": [
        "https://www.w3.org/ns/activitystreams",
        {
          "@language": "es"
        }
      ],
      "type": "Profile",
      "actor": "https://appsocial.com/sofia/",
      "name": "Perfil de Sofia",
      "object": {
        "id": "https://appsocial.com/sofia/id1234",
        "type": "Person",
        "name": "Sofia Rodriguez",
        "preferredUsername": "SofiR",
        "summary": "Amante de la naturaleza y entusiasta de la tecnología.",
            "icon": "https://kenzoishii.example.com/image/165987aklre4",
        "endpoints": {
          "sharedInbox": "https://appsocial.com/inbox"
        },
        "importedFrom": "https://mastodon.ejemplo.com/@sofiR"
      },
      "to": [
        "https://appsocial.com/usuarios/"
      ],
      "cc": "https://appsocial.com/seguidores/sofia",
        "icon": "https://appsocial.com/image/sofia"
    }`)

	bodyRequest := bytes.NewReader(data)
	readCloser := io.NopCloser(bodyRequest)

	// Create a ioReadCloser object
	expectedPerson := activitypub.Actor{
		ID:                "https://example.com/profile/1",
		Type:              "Person",
		Name:              activitypub.DefaultNaturalLanguageValue("John Doe"),
		PreferredUsername: activitypub.DefaultNaturalLanguageValue("John"),
		Summary:           activitypub.DefaultNaturalLanguageValue("John's personal profile"),
		Inbox:             activitypub.IRI("https://example.com/inbox"),
		Outbox:            activitypub.IRI("https://example.com/outbox"),
		Followers:         activitypub.IRI("https://example.com/followers"),
		Following:         activitypub.IRI("https://example.com/following"),
		Liked:             activitypub.IRI("https://example.com/liked"),
	}

	expectedRequest := ImportProfileRequest{
		Person: &expectedPerson,
	}

	importProfileRequest, err := NewImportProfileRequest(&readCloser)
	if err != nil {
		t.Error("Error to create profile request")
	}

	if expectedRequest.Person != importProfileRequest.Person {
		t.Error("Error invalid Person")
	}
}
