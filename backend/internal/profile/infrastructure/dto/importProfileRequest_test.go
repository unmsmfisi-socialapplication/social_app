package dto

import (
	"bytes"
	"encoding/json"
	"io"
	"testing"
)

func NewProfileRequestTest(t *testing.T) {
	t.Log("NewProfileRequestTest")

	data := []byte(`
	{
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
			"profileImage": "https://appsocial.com/sofia/profileImage.jpg",
			"coverImage": "https://appsocial.com/sofia/coverImage.jpg",
			"endpoints": {
				"sharedInbox": "https://appsocial.com/inbox"
			},
			"importedFrom": "https://mastodon.ejemplo.com/@sofiR"
		},
		"to": [
			"https://appsocial.com/usuarios/"
		],
		"cc": "https://appsocial.com/seguidores/sofia"
	}
	`)

	bodyRequest := bytes.NewReader(data)
	readCloser := io.NopCloser(bodyRequest)

	expectedRequest := ImportProfileRequest{
		Context: json.RawMessage{},
		Type:    "Profile",
		Actor:   "https://appsocial.com/sofia/",
		Name:    "Perfil de Sofia",
		Object: ObjectInput{
			Id:   "https://appsocial.com/sofia/id1234",
			Type: "Person",
			Name: "Sofia Rodriguez",
		},
	}

	importProfileRequest, err := NewImportProfileRequest(&readCloser)
	if err != nil {
		t.Error("Error to create profile request")
	}

	if expectedRequest.Actor != importProfileRequest.Actor {
		t.Error("Actor is not the expected")
	}

	if expectedRequest.Name != importProfileRequest.Name {
		t.Error("Name is not the expected")
	}

	if expectedRequest.Object.Id != importProfileRequest.Object.Id {
		t.Error("Object Id is not the expected")
	}

	if expectedRequest.Object.Type != importProfileRequest.Object.Type {
		t.Error("Object Type is not the expected")
	}

	if expectedRequest.Object.Name != importProfileRequest.Object.Name {
		t.Error("Object Name is not the expected")
	}

	if expectedRequest.To[0] != importProfileRequest.To[0] {
		t.Error("To is not the expected")
	}

	if expectedRequest.CC != importProfileRequest.CC {
		t.Error("Cc is not the expected")
	}
}
