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
	"fmt"
	"html/template"
	"io/ioutil"
	"os"
	"reflect"
	"strings"

	"github.com/spf13/cobra"
	yaml "gopkg.in/yaml.v2"
)

const tmpl = `   {{.Key}}  {{.Type}}
`

func inspector(raw []byte) (string, error) {
	results := `type Foo struct {
`
	end := `
}
`
	var output map[string]interface{}
	isjson := IsJSON(raw)
	if isjson != true {
		err := yaml.Unmarshal([]byte(raw), &output)
		if err != nil {
			return results + end, err
		}
	} else {
		err := json.Unmarshal([]byte(raw), &output)
		if err != nil {
			return results + end, err
		}
	}

	for k, v := range output {
		builder := &strings.Builder{}
		key := fmt.Sprintf("%s", strings.Title(k))
		value := fmt.Sprintf("%s", reflect.TypeOf(v))
		data := map[string]interface{}{"Key": key, "Type": value}
		t := template.Must(template.New("keyval").Parse(tmpl))
		if err := t.Execute(builder, data); err != nil {
			return results, err
		}
		s := builder.String()
		results = results + s
	}
	return results + end, nil
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

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// keyCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// keyCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}
