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
	"github.com/octokit/go-octokit/octokit"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"log"
)

var repoListCmd = &cobra.Command{
	Use:   "list",
	Short: "lists all of the repositories for your organisation",
	Long: ``,

	Run: func(cmd *cobra.Command, args []string) {
		client := octokit.NewClientWith(viper.GetString("endpoint"), viper.GetString("useragent"), octokit.TokenAuth{AccessToken: viper.GetString("authtoken")}, nil)
		repos, result := client.Repositories().All(&octokit.OrgRepositoriesURL, octokit.M{"org": viper.GetString("organisation")})
		if result.HasError() {
			log.Fatal(result.Error())
		}

		for _, repo := range repos {
			println(repo.Name)
		}
	},
}

func init() {
	repoCmd.AddCommand(repoListCmd)
}
