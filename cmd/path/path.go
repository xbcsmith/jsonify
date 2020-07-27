// SPDX-FileCopyrightText: 2020 Brett Smith <xbcsmith@gmail.com>
// SPDX-License-Identifier: Apache-2.0

package path

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/xbcsmith/jsonify/jsonify"
)

// pathCmd represents the path command.
var pathCmd = &cobra.Command{ // nolint:gochecknoglobals
	Use:   "path",
	Short: "JSON path syntax support",
	Long:  `JSON path syntax support for JSON or YAML`,
	PreRunE: func(cmd *cobra.Command, args []string) error {
		err := viper.BindPFlags(cmd.Flags())
		if err != nil {
			return err
		}
		return nil
	},
	RunE: pathRunCmd,
}

func pathRunCmd(cmd *cobra.Command, args []string) error {
	path := viper.GetString("jsonpath")

	if len(args) == 0 {
		raw, err := ioutil.ReadAll(os.Stdin)
		if err != nil {
			return errors.Wrap(err, "error reading file")
		}
		results, err := jsonify.PathFinder(raw, path)
		if err != nil {
			return errors.Wrap(err, "error finding jsonpath")
		}
		output, err := json.Marshal(results)
		if err != nil {
			return errors.Wrap(err, "error marshalling output")
		}
		fmt.Printf("%s\n", string(output))
		return nil
	}

	for _, filepath := range args {
		raw, err := ioutil.ReadFile(filepath)
		if err != nil {
			return errors.Wrap(err, "error reading file")
		}
		results, err := jsonify.PathFinder(raw, path)
		if err != nil {
			return errors.Wrap(err, "error finding jsonpath")
		}
		output, err := json.Marshal(results)
		if err != nil {
			return errors.Wrap(err, "error marshalling output")
		}
		fmt.Printf("%s\n", string(output))
	}

	return nil
}

// NewPathCmd returns a new path command.
func NewPathCmd() *cobra.Command {
	pathCmd.Flags().StringP("jsonpath", "p", "", "JSON path")
	return pathCmd
}
