// Copyright © 2018 Chris Hudson <chris@xetius.com>
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
	"fmt"
	"log"
	"os"

	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

var rootCmd = &cobra.Command{
	Use:   "github",
	Short: "A command line application for managing you organisations github",
	Long: `github CLI is a utility for managing your GitHub organisation.

Once configured, you are able to manage various aspects of your GitHub organisation.  
You can do the following things:

  - List repositories
  - Create a new repository from a template
  - Apply configuration to all repositories in your organisation (handy for mass changes)`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringP("organisation", "o", "", "name of your organisation")
	rootCmd.PersistentFlags().StringP("authtoken", "a", "", "authorisation token to access organisation")
	rootCmd.PersistentFlags().StringP("endpoint", "e", "", "URL of the API endpoint for your GitHub instance")
	rootCmd.PersistentFlags().StringP("useragent", "ua", "GitHub CLI Agent", "User-Agent to use for requests to your GitHub instance")

	err := viper.BindPFlag("organisation", rootCmd.PersistentFlags().Lookup("organisation")); if err != nil { log.Fatal(err) }
	err = viper.BindPFlag("authtoken", rootCmd.PersistentFlags().Lookup("authtoken")); if err != nil { log.Fatal(err) }
	err = viper.BindPFlag("endpoint", rootCmd.PersistentFlags().Lookup("endpoint")); if err != nil { log.Fatal(err) }
	err = viper.BindPFlag("useragent", rootCmd.PersistentFlags().Lookup("useragent")); if err != nil { log.Fatal(err) }
}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		viper.SetEnvPrefix("GHCLI")

		viper.AddConfigPath(home)
		viper.SetConfigName(".github")
	}

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
