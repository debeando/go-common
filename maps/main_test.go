package maps_test

import (
	"reflect"
	"testing"

	"github.com/debeando/go-common/maps"
)

func TestKeys(t *testing.T) {
	result := maps.Keys([]map[string]string{{"foo": "a", "bar": "1"}, {"foo": "b", "bar": "2"}})
	expected := []string{"bar", "foo"}

	if !reflect.DeepEqual(result, expected) {
		t.Error("Expected: []string{\"foo\", \"bar\"}")
	}
}

func TestIn(t *testing.T) {
	expected := make(map[string]string)
	expected["test"] = "test"

	if !maps.In("test", expected) {
		t.Error("Expected: true")
	}

	if maps.In("foo", expected) {
		t.Error("Expected: false")
	}
}
