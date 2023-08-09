package json_test

import (
	"testing"

	"github.com/debeando/go-common/json"
	"github.com/stretchr/testify/assert"
)

type TestCase struct {
	Foo string `json:"f"`
	Bar int    `json:"b"`
	Baz bool   `json:"B"`
}

func TestGet(t *testing.T) {
	s, e := json.Escape(`{"foo":"bar"}`)
	assert.Equal(t, s, `{\"foo\":\"bar\"}`)
	assert.NoError(t, e)
}

func TestStructToJSON(t *testing.T) {
	z := &TestCase{Foo: "foo", Bar: 1, Baz: false}

	s, e := json.StructToJSON(z)
	assert.Equal(t, s, `{"f":"foo","b":1,"B":false}`)
	assert.NoError(t, e)
}

func TestStringToStruct(t *testing.T) {
	z := &TestCase{}
	e := json.StringToStruct(`{"f":"foo","b":1,"B":false}`, &z)
	assert.NoError(t, e)
	s, e := json.StructToJSON(z)
	assert.NoError(t, e)
	assert.Equal(t, s, `{"f":"foo","b":1,"B":false}`)
}
