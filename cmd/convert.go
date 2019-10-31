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

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	yaml "gopkg.in/yaml.v3"
)

// IsJSON try to guess if file is json or yaml
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

func json2yaml(raw []byte) ([]byte, error) {
	var output interface{}
	if err := json.Unmarshal([]byte(raw), &output); err != nil {
		return nil, err
	}

	content, err := yaml.Marshal(output)
	if err != nil {
		return nil, fmt.Errorf("failed to convert to yaml : %v", err)
	}
	return content, nil
}

func yaml2json(raw []byte, noindent bool) ([]byte, error) {
	// ms := yaml.MapSlice{}
	var output interface{}
	if err := yaml.Unmarshal(raw, &output); err != nil {
		return nil, err
	}

	if noindent {
		content, err := json.Marshal(output)
		if err != nil {
			return nil, fmt.Errorf("failed to convert to yaml : %v", err)
		}
		return content, nil
	}

	content, err := json.MarshalIndent(output, "", "  ")
	if err != nil {
		return nil, fmt.Errorf("failed to convert to yaml : %v", err)
	}
	return content, nil

}

func converter(raw []byte, noindent bool) ([]byte, error) {
	isjson := IsJSON(raw)
	if isjson != true {
		output, err := yaml2json(raw, noindent)
		if err != nil {
			fmt.Printf("decode data: %v", err)
			return nil, err
		}
		return output, nil
	}
	output, err := json2yaml(raw)
	if err != nil {
		fmt.Printf("decode data: %v", err)
		return nil, err
	}
	return output, nil
}

// convertCmd represents the convert command
var convertCmd = &cobra.Command{
	Use:   "convert",
	Short: "convert json into yaml and yaml into json",
	Long:  `Convert JSON into YAML and YAML into JSON`,
	PreRunE: func(cmd *cobra.Command, args []string) error {
		err := viper.BindPFlags(cmd.Flags())
		if err != nil {
			return err
		}
		return nil
	},
	Run: convertRunCmd,
}

func convertRunCmd(cmd *cobra.Command, args []string) {
	noindent := viper.GetBool("noindent")
	if len(args) > 0 {
		for _, filepath := range args {
			raw, err := ioutil.ReadFile(filepath)
			if err != nil {
				fmt.Println(err)
				os.Exit(-1)
			}
			output, err := converter(raw, noindent)
			if err != nil {
				fmt.Println(err)
				os.Exit(-1)
			}
			fmt.Printf("%s", output)
		}
	} else {
		raw, err := ioutil.ReadAll(os.Stdin)
		if err != nil {
			fmt.Println(err)
			os.Exit(-1)
		}
		output, err := converter(raw, noindent)
		if err != nil {
			fmt.Println(err)
			os.Exit(-1)
		}
		fmt.Printf("%s", output)
	}
	os.Exit(0)

}

func init() {
	convertCmd.Flags().Bool("noindent", false, "skip format json")
	RootCmd.AddCommand(convertCmd)
}
