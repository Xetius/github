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
	"github.com/spf13/viper"
	"log"

	"github.com/spf13/cobra"
)

var repoCreateCmd = &cobra.Command{
	Use:   "create [flags] [repo name]",
	Short: "Creates a new repository",
	Long: `
Creates a new repository from a template
`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Creating repository %s\n", args[0])
	},
}

func init() {
	repoCmd.AddCommand(repoCreateCmd)

	repoCreateCmd.PersistentFlags().StringP("template", "t", "", "create repository template file name")

	err := viper.BindPFlag("template", repoCreateCmd.PersistentFlags().Lookup("template"));  if err != nil { log.Fatal(err) }
}
