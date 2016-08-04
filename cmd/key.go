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
	//"encoding/json"
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"github.com/spf13/cobra"
	//"io/ioutil"
	//"os"
	"strings"
)

// keyCmd represents the key command
var keyCmd = &cobra.Command{
	Use:   "key",
	Short: "Return the value of specified key",
	Long:  `Return the value of specified key`,
	Run:   getKey,
}

func getKey(cmd *cobra.Command, args []string) {

	fmt.Println("PrintArgs: " + strings.Join(args, " "))

	spew.Dump(args)
	/*
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

		for k, v := range m {
			if _, ok := args[k]; ok {
				fmt.Println(v)
			}
		}

		os.Exit(0)
	*/
}

func init() {
	RootCmd.AddCommand(keyCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// keyCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// keyCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}
