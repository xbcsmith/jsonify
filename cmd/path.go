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
	"io/ioutil"
	"os"

	yaml "gopkg.in/yaml.v3"

	"github.com/oliveagle/jsonpath"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func pathfinder(raw []byte, path string) (interface{}, error) {
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
		fmt.Printf("path not found: %v\n", err)
		return nil, err
	}
	return res, nil
}

// pathCmd represents the path command
var pathCmd = &cobra.Command{
	Use:   "path",
	Short: "path json path syntax support",
	Long:  `path json path syntax support for JSON or YAML`,
	PreRunE: func(cmd *cobra.Command, args []string) error {
		err := viper.BindPFlags(cmd.Flags())
		if err != nil {
			return err
		}
		return nil
	},
	Run: pathRunCmd,
}

func pathRunCmd(cmd *cobra.Command, args []string) {
	path := viper.GetString("jsonpath")
	if len(args) > 0 {
		for _, filepath := range args {
			raw, err := ioutil.ReadFile(filepath)
			if err != nil {
				fmt.Println(err)
				os.Exit(-1)
			}
			results, err := pathfinder(raw, path)
			if err != nil {
				fmt.Println(err)
				os.Exit(-1)
			}
			output, err := json.Marshal(results)
			if err != nil {
				fmt.Printf("Marshal error: %v\n", err)
				os.Exit(-1)
			}
			fmt.Printf("%s\n", string(output))
		}
	} else {
		raw, err := ioutil.ReadAll(os.Stdin)
		if err != nil {
			fmt.Println(err)
			os.Exit(-1)
		}
		results, err := pathfinder(raw, path)
		if err != nil {
			fmt.Println(err)
			os.Exit(-1)
		}
		output, err := json.Marshal(results)
		if err != nil {
			fmt.Printf("Marshal error: %v\n", err)
			os.Exit(-1)
		}
		fmt.Printf("%s\n", string(output))
	}
	os.Exit(0)
}

func init() {
	RootCmd.AddCommand(pathCmd)
	pathCmd.Flags().StringP("jsonpath", "p", "", "JSON path")
}
