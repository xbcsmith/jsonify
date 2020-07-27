// SPDX-FileCopyrightText: 2020 Brett Smith <xbcsmith@gmail.com>
// SPDX-License-Identifier: Apache-2.0

package jsonify

import (
	"encoding/json"

	"github.com/pkg/errors"
	yaml "gopkg.in/yaml.v3"
)

func json2yaml(raw []byte) ([]byte, error) {
	var output interface{}
	if err := json.Unmarshal(raw, &output); err != nil {
		return nil, err
	}
	content, err := yaml.Marshal(output)
	if err != nil {
		return nil, errors.Wrap(err, "failed to convert to yaml")
	}
	return content, nil
}

func yaml2json(raw []byte, noindent bool) ([]byte, error) {
	var output interface{}
	if err := yaml.Unmarshal(raw, &output); err != nil {
		return nil, err
	}
	if noindent {
		content, err := json.Marshal(output)
		if err != nil {
			return nil, errors.Wrap(err, "failed to convert to yaml")
		}
		return content, nil
	}
	content, err := json.MarshalIndent(output, "", "  ")
	if err != nil {
		return nil, errors.Wrap(err, "failed to convert to yaml")
	}
	return content, nil
}

// Converter takes a raw byte array and returns a byte array of YAML or JSON.
func Converter(raw []byte, noindent bool) ([]byte, error) {
	isjson := IsJSON(raw)
	if !isjson {
		output, err := yaml2json(raw, noindent)
		if err != nil {
			return nil, errors.Wrap(err, "decode data")
		}
		return output, nil
	}
	output, err := json2yaml(raw)
	if err != nil {
		return nil, errors.Wrap(err, "decode data")
	}
	return output, nil
}
