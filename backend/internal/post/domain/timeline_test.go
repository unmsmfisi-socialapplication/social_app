package domain

import (
	"reflect"
	"testing"
)

func TestNewQueryResult(t *testing.T) {
	timeline := []TimelineRes{}
	result := NewQueryResult(timeline)

	if !reflect.DeepEqual(result.Results, timeline) {
		t.Errorf("Expected results: %v, results obtained: %v", timeline, result.Results)
	}
}
