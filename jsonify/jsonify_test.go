// SPDX-FileCopyrightText: 2020 Brett Smith <xbcsmith@gmail.com>
// SPDX-License-Identifier: Apache-2.0

package jsonify

import (
	"encoding/json"
	"testing"

	yaml "gopkg.in/yaml.v3"
	"gotest.tools/assert"
	is "gotest.tools/assert/cmp"
)

type Tests struct {
	a []byte
	b []byte
	c map[string]interface{}
	d interface{}
	e []byte
}

const ystr = `
bar:
  - buz
  - cuz
  - duz
baz:
    caz: fuz
flag: true
foo: show_value_of_foo
fuzzy:
    complicated-it-is:
        could_be-but:
            not-really_possible:
                until_it_is: true
yyy:
  - one
  - 2
  - true
  - "4"
  - key: value
  - - 1
    - "2"
    - things:
        - complicated: true
          couldbe: maybe
          notreally: false
zzz-zzz:
    buz:
      - 1
      - 2
      - 3
`

const jstr = `{
  "bar": [
    "buz",
    "cuz",
    "duz"
  ],
  "baz": {
    "caz": "fuz"
  },
  "flag": true,
  "foo": "show_value_of_foo",
  "fuzzy": {
    "complicated-it-is": {
      "could_be-but": {
        "not-really_possible": {
          "until_it_is": true
        }
      }
    }
  },
  "yyy": [
    "one",
    2,
    true,
    "4",
    {
      "key": "value"
    },
    [
      1,
      "2",
      {
        "things": [
          {
            "complicated": true,
            "couldbe": "maybe",
            "notreally": false
          }
        ]
      }
    ]
  ],
  "zzz-zzz": {
    "buz": [
      1,
      2,
      3
    ]
  }
}
`

func NewTests() *Tests {
	var j map[string]interface{}
	if err := json.Unmarshal([]byte(jstr), &j); err != nil {
		panic(err)
	}
	var y map[string]interface{}
	if err := yaml.Unmarshal([]byte(ystr), &y); err != nil {
		panic(err)
	}
	tests := &Tests{
		a: []byte(jstr),
		b: []byte(ystr),
		c: j,
		d: y,
		e: []byte(`[{"brackets" : "json"}]`),
	}
	return tests
}

func TestPathFinderJson(t *testing.T) {
	tests := NewTests()
	path := `$.baz.caz`
	expected := `fuz`
	actual, err := PathFinder(tests.a, path)
	assert.Assert(t, is.Nil(err))
	assert.Equal(t, expected, actual)
}

func TestPathFinderYaml(t *testing.T) {
	tests := NewTests()
	path := `$.baz.caz`
	expected := `fuz`
	actual, err := PathFinder(tests.b, path)
	assert.Assert(t, is.Nil(err))
	assert.Equal(t, expected, actual)
}

func TestPathFinderJsonMissing(t *testing.T) {
	tests := NewTests()
	path := `$.foo.bar`
	actual, err := PathFinder(tests.a, path)
	assert.Assert(t, is.Nil(actual))
	assert.ErrorContains(t, err, "object is not map")
}

func TestPathFinderYamlMissing(t *testing.T) {
	tests := NewTests()
	path := `$.foo.bar`
	actual, err := PathFinder(tests.b, path)
	assert.Assert(t, is.Nil(actual))
	assert.ErrorContains(t, err, "object is not map")
}

func TestConverterJson(t *testing.T) {
	tests := NewTests()
	expected := `- key:`
	actual, err := Converter(tests.a, false)
	assert.Assert(t, is.Nil(err))
	assert.Assert(t, !IsJSON(actual))
	assert.Assert(t, is.Contains(string(actual), expected))
}

func TestConverterAlsoJson(t *testing.T) {
	tests := NewTests()
	expected := `brackets:`
	actual, err := Converter(tests.e, false)
	assert.Assert(t, is.Nil(err))
	assert.Assert(t, !IsJSON(actual))
	assert.Assert(t, is.Contains(string(actual), expected))
}

func TestConverterYaml(t *testing.T) {
	tests := NewTests()
	expected := `"bar": [`
	actual, err := Converter(tests.b, false)
	assert.Assert(t, is.Nil(err))
	assert.Assert(t, IsJSON(actual))
	assert.Assert(t, is.Contains(string(actual), expected))
}

func TestConverterYamlNoIndent(t *testing.T) {
	tests := NewTests()
	expected := `bar":[`
	actual, err := Converter(tests.b, true)
	assert.Assert(t, is.Nil(err))
	assert.Assert(t, IsJSON(actual))
	assert.Assert(t, is.Contains(string(actual), expected))
}
