// +build !integration

package linode

import (
	"encoding/json"
	"testing"

	"github.com/alexsacr/linode/_third_party/testify/assert"
	"github.com/alexsacr/linode/_third_party/testify/require"
)

func TestMarshalArgsErrors(t *testing.T) {
	_, err := marshallArgs("foo")
	assert.Error(t, err)

	noTag := struct {
		foo string
	}{
		"bar",
	}

	_, err = marshallArgs(noTag)
	assert.Error(t, err)

	badIntEncode := struct {
		Foo *string `args:"foo,int"`
	}{
		String("bar"),
	}

	_, err = marshallArgs(badIntEncode)
	assert.Error(t, err)
}

func TestUnmarshalSingleErrs(t *testing.T) {
	err := unmarshalSingle(json.RawMessage{}, "", "foo")
	assert.Error(t, err)

	var foo string
	err = unmarshalSingle(json.RawMessage([]byte(`{"bar":"baz"}`)), "foo", &foo)
	assert.Error(t, err)

	var outInt int
	var outStr string
	var outBool bool
	var outSlice []struct{}

	err = unmarshalSingle(json.RawMessage([]byte(`{"foo": "bar"}`)), "foo", &outInt)
	assert.Error(t, err)
	assert.Empty(t, outInt)
	err = unmarshalSingle(json.RawMessage([]byte(`{"foo": 1}`)), "foo", &outStr)
	assert.Error(t, err)
	assert.Empty(t, outStr)
	err = unmarshalSingle(json.RawMessage([]byte(`{"foo": "bar"}`)), "foo", &outBool)
	assert.Error(t, err)
	assert.Empty(t, outBool)
	err = unmarshalSingle(json.RawMessage([]byte(`{"foo": "bar"}`)), "foo", &outSlice)
	assert.Error(t, err)
	assert.Empty(t, outSlice)
}

func TestUnmarshalSingleBool(t *testing.T) {
	// doesn't currently get exercised in normal usage
	var out bool
	err := unmarshalSingle(json.RawMessage([]byte(`{"foo": true}`)), "foo", &out)
	require.NoError(t, err)
	assert.True(t, out)
}

func TestUnmarshalMultiMapErrs(t *testing.T) {
	err := unmarshalMultiMap(json.RawMessage{}, "foo")
	assert.Error(t, err)

	var badOut *string
	err = unmarshalMultiMap(json.RawMessage([]byte(`[{"foo": "bar"}]`)), &badOut)
	assert.Error(t, err)
}
