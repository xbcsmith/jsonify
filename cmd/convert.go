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
	"fmt"
	"io/ioutil"
	"os"

	"encoding/json"
	"unicode"

	"github.com/icza/dyno"
	"github.com/spf13/cobra"
	yaml "gopkg.in/yaml.v2"
)

// IsJSON try to guess if file is json or yaml
func IsJSON(buf []byte) bool {
	var prefix = []byte("{")
	trim := bytes.TrimLeftFunc(buf, unicode.IsSpace)
	return bytes.HasPrefix(trim, prefix)
}

func json2yaml(d []byte) ([]byte, error) {
	ms := yaml.MapSlice{}
	err := yaml.Unmarshal(d, &ms)
	if err != nil {
		return nil, err
	}
	m2 := dyno.ConvertMapI2MapS(ms)
	return yaml.Marshal(m2)
}

func yaml2json(d []byte) ([]byte, error) {
	ms := yaml.MapSlice{}
	err := yaml.Unmarshal(d, &ms)
	if err != nil {
		return nil, err
	}
	m2 := dyno.ConvertMapI2MapS(ms)
	return json.MarshalIndent(m2, "", "  ")
}

// convertCmd represents the convert command
var convertCmd = &cobra.Command{
	Use:   "convert",
	Short: "convert json into yaml and yaml into json",
	Long:  `Convert JSON into YAML and YAML into JSON`,
	Run:   convertRunCmd,
}

func convertRunCmd(cmd *cobra.Command, args []string) {
	raw, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	isjson := IsJSON(raw)
	if isjson != true {
		output, err := yaml2json(raw)
		if err != nil {
			fmt.Printf("decode data: %v", err)
			os.Exit(-1)
		}
		fmt.Printf("%s", output)
	} else {
		output, err := json2yaml(raw)
		if err != nil {
			fmt.Printf("decode data: %v", err)
			os.Exit(-1)
		}
		fmt.Printf("%s", output)
	}
	os.Exit(0)
}

func init() {
	RootCmd.AddCommand(convertCmd)
}
