// SPDX-FileCopyrightText: 2020 Brett Smith <xbcsmith@gmail.com>
// SPDX-License-Identifier: Apache-2.0

package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/xbcsmith/jsonify/cmd/convert"
	"github.com/xbcsmith/jsonify/cmd/path"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands.
var rootCmd = &cobra.Command{
	Use:   "jsonify",
	Short: "jsonify -- manipulate json and yaml from stdin",
	Long: `jsonify is a small executable used to manipulate json and yaml from the command line
        Example:
        echo '{"foo":"show_value_of_foo","bar":"buz"}' | jsonify
        jsonify convert foo.yml`,
	PreRunE: func(cmd *cobra.Command, args []string) error {
		err := viper.BindPFlags(cmd.Flags())
		if err != nil {
			return err
		}
		return nil
	},
	RunE: convert.ConvertCmd,
}

// Execute adds all child commands to the root command sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.jsonify.yaml)")
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	convertCmd := convert.NewConvertCmd()
	rootCmd.AddCommand(convertCmd)
	pathCmd := path.NewPathCmd()
	rootCmd.AddCommand(pathCmd)
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	}

	viper.SetConfigName(".jsonify")
	viper.AddConfigPath("$HOME")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		fmt.Printf("Using config file: %s\n", viper.ConfigFileUsed())
	}
}
