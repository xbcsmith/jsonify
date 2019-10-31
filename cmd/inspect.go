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
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"go/parser"
	"go/printer"
	"go/token"
	yaml "gopkg.in/yaml.v3"
	"html/template"
	"io/ioutil"
	"os"
	"reflect"
	"strings"
)

const tmpl = `   {{.Key}}  {{.Type}}`

const begin = `
package main

// Foo struct generated
type Foo struct {
`
const stmpl = `  {{.Key}} struct {
`

const end = `}

`
const pad = `   `

const ret = `
`
const etmpl = `json:"{{.Key}}" yaml:"{{.Key}}"`

const (
	indentAdj   = 0
	tabWidth    = 8
	printerMode = printer.UseSpaces | printer.TabIndent
)

var (
	parserMode parser.Mode
)

func initParserMode() {
	parserMode = parser.ParseComments
	if true {
		parserMode |= parser.AllErrors
	}
}

func maketmpl(data map[string]interface{}, tmpl string) (string, error) {
	builder := &strings.Builder{}
	t := template.Must(template.New("new").Parse(tmpl))
	if err := t.Execute(builder, data); err != nil {
		return ``, err
	}
	s := builder.String()
	return s, nil
}

func keyHandler(key string, low bool) string {
	if low {
		new := strings.ToLower(key)
		return new
	}
	r := strings.NewReplacer("-", "_")
	t := r.Replace(key)
	new := ``
	for _, wrd := range strings.Split(t, "_") {
		new = new + strings.Title(wrd)
	}

	return new
}

func tostruct(input map[string]interface{}) (string, error) {
	results := ``
	for k, v := range input {
		// TODO look for _ or - and remove them while Capitalizing next letter
		key := fmt.Sprintf("%s", keyHandler(k, false))
		value := fmt.Sprintf("%s", reflect.TypeOf(v))
		data := map[string]interface{}{"Key": key, "Type": value}
		s, err := maketmpl(data, tmpl)
		if err != nil {
			return results, err
		}
		edata := map[string]interface{}{"Key": keyHandler(k, true)}
		e, err := maketmpl(edata, etmpl)
		if err != nil {
			return results, err
		}
		results = results + s + pad + "`" + e + "`" + ret
	}
	return results, nil
}

// isSpace is copied from go/src/cmd/gofmt/internal.go
// isSpace reports whether the byte is a space character.
// isSpace defines a space as being among the following bytes: ' ', '\t', '\n' and '\r'.
func isSpace(b byte) bool {
	return b == ' ' || b == '\t' || b == '\n' || b == '\r'
}

// format is derived from go/src/cmd/gofmt/internal.go
func format(src []byte) ([]byte, error) {
	initParserMode()
	fset := token.NewFileSet()
	empty := []byte("")
	file, err := parser.ParseFile(fset, "foo.go", src, parserMode)
	if err != nil {
		return empty, err
	}
	i, j := 0, 0
	for j < len(src) && isSpace(src[j]) {
		if src[j] == '\n' {
			i = j + 1
		}
		j++
	}
	var res []byte
	res = append(res, src[:i]...)
	indent := 0
	hasSpace := false
	for _, b := range src[i:j] {
		switch b {
		case ' ':
			hasSpace = true
		case '\t':
			indent++
		}
	}
	if indent == 0 && hasSpace {
		indent = 1
	}
	for i := 0; i < indent; i++ {
		res = append(res, '\t')
	}
	cfg := printer.Config{
		Mode:     0,
		Tabwidth: tabWidth,
		Indent:   0,
	}
	cfg.Indent = indent + indentAdj
	var buf bytes.Buffer
	if err := cfg.Fprint(&buf, fset, file); err != nil {
		return empty, err
	}
	sourceAdj := func(src []byte, indent int) []byte {
		return bytes.TrimSpace(src)
	}
	out := sourceAdj(buf.Bytes(), cfg.Indent)
	if len(out) == 0 {
		return src, nil
	}
	res = append(res, out...)
	i = len(src)
	for i > 0 && isSpace(src[i-1]) {
		i--
	}

	return append(res, src[i:]...), nil

}

func inspector(raw []byte) (string, error) {
	var output map[string]interface{}
	isjson := IsJSON(raw)
	if isjson != true {
		err := yaml.Unmarshal([]byte(raw), &output)
		if err != nil {
			return begin + end, err
		}
	} else {
		err := json.Unmarshal([]byte(raw), &output)
		if err != nil {
			return begin + end, err
		}
	}
	results, err := tostruct(output)
	if err != nil {
		return begin + end, err
	}
	for k, v := range output {
		T, ok := v.(map[string]interface{})
		if ok {
			d := map[string]interface{}{"Key": keyHandler(k, false)}
			b, _ := maketmpl(d, stmpl)
			s, _ := tostruct(T)
			results = results + b + pad + s + pad + end
		}

	}
	code := begin + results + end
	final, err := format([]byte(code))
	if err != nil {
		return "", err
	}
	return string(final), nil
}

// inspectCmd represents the inspect command
var inspectCmd = &cobra.Command{
	Use:   "inspect",
	Short: "inspect json types",
	Long:  `Display types and kinds for all elements in json`,
	Run:   inspectRunCmd,
}

func inspectRunCmd(cmd *cobra.Command, args []string) {

	if len(args) > 0 {
		for _, filepath := range args {
			raw, err := ioutil.ReadFile(filepath)
			if err != nil {
				fmt.Println(err)
				os.Exit(-1)
			}
			output, err := inspector(raw)
			if err != nil {
				fmt.Println(err)
				os.Exit(-1)
			}
			fmt.Print(output)
		}
	} else {
		raw, err := ioutil.ReadAll(os.Stdin)
		if err != nil {
			fmt.Println(err)
			os.Exit(-1)
		}
		output, err := inspector(raw)
		if err != nil {
			fmt.Println(err)
			os.Exit(-1)
		}
		fmt.Print(output)
	}
	os.Exit(0)
}

func init() {
	RootCmd.AddCommand(inspectCmd)
}
