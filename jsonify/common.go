// SPDX-FileCopyrightText: 2020 Brett Smith <xbcsmith@gmail.com>
// SPDX-License-Identifier: Apache-2.0

package jsonify

import (
	"bytes"
	"unicode"
)

// IsJSON try to guess if file is JSON or YAML.
func IsJSON(buf []byte) bool {
	trim := bytes.TrimLeftFunc(buf, unicode.IsSpace)
	if bytes.HasPrefix(trim, []byte("{")) {
		return true
	}
	if bytes.HasPrefix(trim, []byte("[")) {
		return true
	}
	return false
}
