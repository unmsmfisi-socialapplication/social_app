package helpers

import (
	"net/url"
	"reflect"
	"testing"

	"github.com/unmsmfisi-socialapplication/social_app/internal/post/domain"
)

func TestParsePaginationParams(t *testing.T) {
	validQuery := url.Values{
		"page":  {"2"},
		"limit": {"10"},
	}

	expectedParams := domain.PostPaginationParams{
		Page:  2,
		Limit: 10,
	}

	params, err := ParsePaginationParams(validQuery)

	if err != nil {
		t.Errorf("Se esperaba que ParsePaginationParams no devolviera un error, pero se obtuvo un error: %v", err)
	}

	if !reflect.DeepEqual(params, expectedParams) {
		t.Errorf("Los parámetros obtenidos no son los esperados. Esperado: %+v, Obtenido: %+v", expectedParams, params)
	}

	missingValuesQuery := url.Values{}

	params, err = ParsePaginationParams(missingValuesQuery)

	if err != nil {
		t.Errorf("Se esperaba que ParsePaginationParams no devolviera un error con valores faltantes, pero se obtuvo un error: %v", err)
	}

	if !reflect.DeepEqual(params, domain.PostPaginationParams{}) {
		t.Errorf("Los parámetros obtenidos no son los esperados con valores faltantes. Esperado: %+v, Obtenido: %+v", domain.PostPaginationParams{}, params)
	}

	invalidValuesQuery := url.Values{
		"page":  {"invalid"},
		"limit": {"invalid"},
	}

	_, err = ParsePaginationParams(invalidValuesQuery)

	if err == nil {
		t.Error("Se esperaba un error al proporcionar valores no numéricos, pero no se obtuvo un error")
	}
}
