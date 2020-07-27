// SPDX-FileCopyrightText: 2020 Brett Smith <xbcsmith@gmail.com>
// SPDX-License-Identifier: Apache-2.0

package jsonify

import (
	"encoding/json"

	yaml "gopkg.in/yaml.v3"

	"github.com/oliveagle/jsonpath"
	"github.com/pkg/errors"
)

// PathFinder takes a jsonpath query and returns an interface.
func PathFinder(raw []byte, path string) (interface{}, error) {
	var output interface{}
	isjson := IsJSON(raw)
	if !isjson {
		err := yaml.Unmarshal(raw, &output)
		if err != nil {
			return nil, err
		}
	} else {
		err := json.Unmarshal(raw, &output)
		if err != nil {
			return nil, err
		}
	}
	res, err := jsonpath.JsonPathLookup(output, path)
	if err != nil {
		return nil, errors.Wrap(err, "path not found")
	}
	return res, nil
}
