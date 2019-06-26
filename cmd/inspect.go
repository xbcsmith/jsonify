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
	"reflect"

	"github.com/spf13/cobra"
)

// inspectCmd represents the inspect command
var inspectCmd = &cobra.Command{
	Use:   "inspect",
	Short: "inspect json types",
	Long:  `Display types and kinds for all elements in json`,
	Run:   inspectAll,
}

func inspectAll(cmd *cobra.Command, args []string) {

	raw, err := ioutil.ReadAll(os.Stdin)

	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

	var data interface{}

	err = json.Unmarshal(raw, &data)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

	m := data.(map[string]interface{})

	fmt.Println("INSPECT MODULE")

	for k, v := range m {
		fmt.Printf("key:%v  value:%v  kind:%s  type:%s\n\n", k, v, reflect.TypeOf(v).Kind(), reflect.TypeOf(v))
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
