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
	"testing"

	"gotest.tools/assert"
	is "gotest.tools/assert/cmp"
)

func TestConverterJson(t *testing.T) {
	tests := NewTests()
	expected := `- key:`
	actual, err := converter(tests.a)
	assert.Assert(t, is.Nil(err))
	assert.Assert(t, !IsJSON(actual))
	assert.Assert(t, is.Contains(string(actual), expected))
}

func TestConverterYaml(t *testing.T) {
	tests := NewTests()
	expected := `"bar": [`
	actual, err := converter(tests.b)
	assert.Assert(t, is.Nil(err))
	assert.Assert(t, IsJSON(actual))
	assert.Assert(t, is.Contains(string(actual), expected))
}
