// SPDX-FileCopyrightText: 2020 Brett Smith <xbcsmith@gmail.com>
// SPDX-License-Identifier: Apache-2.0

package convert

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/xbcsmith/jsonify/jsonify"
)

// convertCmd represents the convert command.
var convertCmd = &cobra.Command{ // nolint:gochecknoglobals
	Use:   "convert",
	Short: "convert JSON into YAML or YAML into JSON",
	Long:  `Convert JSON into YAML or YAML into JSON`,
	PreRunE: func(cmd *cobra.Command, args []string) error {
		err := viper.BindPFlags(cmd.Flags())
		if err != nil {
			return err
		}
		return nil
	},
	RunE: convertRunCmd,
}

// ConvertCmd export for using in main jsonify.
var ConvertCmd = convertRunCmd // nolint:gochecknoglobals

// convertRunCmd is the main command for jsonify.
func convertRunCmd(cmd *cobra.Command, args []string) error {
	noindent := viper.GetBool("noindent")

	if len(args) == 0 {
		raw, err := ioutil.ReadAll(os.Stdin)
		if err != nil {
			return errors.Wrap(err, "error reading file")
		}
		output, err := jsonify.Converter(raw, noindent)
		if err != nil {
			return errors.Wrap(err, "error converting input")
		}
		fmt.Printf("%s", output)
		return nil
	}

	for _, filepath := range args {
		raw, err := ioutil.ReadFile(filepath)
		if err != nil {
			return errors.Wrap(err, "error reading file")
		}
		output, err := jsonify.Converter(raw, noindent)
		if err != nil {
			return errors.Wrap(err, "error converting input")
		}
		fmt.Printf("%s", output)
	}

	return nil
}

// NewConvertCmd returns a new convert command.
func NewConvertCmd() *cobra.Command {
	convertCmd.Flags().Bool("noindent", false, "skip format json")
	return convertCmd
}
