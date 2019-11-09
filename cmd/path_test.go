// Copyright © 2016 Brett Smith <bc.smith@sas.com>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

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
	actual, err := pathfinder(tests.a, path)
	assert.Assert(t, is.Nil(err))
	assert.Equal(t, expected, actual)
}

func TestPathFinderYaml(t *testing.T) {
	tests := NewTests()
	path := `$.baz.caz`
	expected := `fuz`
	actual, err := pathfinder(tests.b, path)
	assert.Assert(t, is.Nil(err))
	assert.Equal(t, expected, actual)
}

func TestPathFinderJsonMissing(t *testing.T) {
	tests := NewTests()
	path := `$.foo.bar`
	actual, err := pathfinder(tests.a, path)
	assert.Assert(t, is.Nil(actual))
	assert.ErrorContains(t, err, "object is not map")
}

func TestPathFinderYamlMissing(t *testing.T) {
	tests := NewTests()
	path := `$.foo.bar`
	actual, err := pathfinder(tests.b, path)
	assert.Assert(t, is.Nil(actual))
	assert.ErrorContains(t, err, "object is not map")
}
