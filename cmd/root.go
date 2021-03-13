/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"github.com/pedrorsantana/greysec/cmd/greysec"
	"github.com/spf13/cobra"
)

var cfgKey string
var cfgInterface string
var cfgCacheSize int

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "greysec",
	Short: "GreySec is a firewall connected to the base of greynoise.io that identifies and blocks requests from attacking machines.",
	Long: `GreySec is a firewall application created to use greynoise.io data to identify
 package communication with an attacker and block the communication.

Greysec is able to block automated attacks and also block automatic scanners
 like shodan.io and zoomeye.org.

Usage Example:
  greysec --key "XPTO" [--interface eth0] [--cacheSize 120]

Feel free to contribute in Github: https://github.com/pedrorsantana/greysec/
`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		greysec.Run(cfgKey, cfgInterface, cfgCacheSize)
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	rootCmd.PersistentFlags().StringVar(&cfgKey, "key", "", "Authentication greynoise.io key.")
	rootCmd.MarkPersistentFlagRequired("key")

	rootCmd.PersistentFlags().StringVar(&cfgInterface, "interface", "", "The interface defined to listen.")

	rootCmd.PersistentFlags().IntVar(&cfgCacheSize, "cacheSize", 120, "The size of the cache, can help with bandwidth problems.")
}
