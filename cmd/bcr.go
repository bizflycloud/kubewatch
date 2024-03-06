/*
Copyright 2016 Skippbox, Ltd.
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
	"github.com/bitnami-labs/kubewatch/config"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// bcr
var bcrConfigCmd = &cobra.Command{
	Use:   "bcr",
	Short: "specific bcr configuration",
	Long:  `specific bcr configuration`,
	Run: func(cmd *cobra.Command, args []string) {
		conf, err := config.New()
		if err != nil {
			logrus.Fatal(err)
		}

		name, err := cmd.Flags().GetString("name")
		if err == nil {
			if len(name) > 0 {
				conf.Handler.Bcr.Name = name
			}
		}
		host, err := cmd.Flags().GetString("host")
		if err == nil {
			if len(host) > 0 {
				conf.Handler.Bcr.Host = host
			}
		}
		appCredId, err := cmd.Flags().GetString("appCredId")
		if err == nil {
			if len(appCredId) > 0 {
				conf.Handler.Bcr.AppCredId = appCredId
			}
		}
		appCredSecret, err := cmd.Flags().GetString("appCredSecret")
		if err == nil {
			if len(appCredId) > 0 {
				conf.Handler.Bcr.AppCredSecret = appCredSecret
			}
		}
		region, err := cmd.Flags().GetString("region")
		if err == nil {
			if len(appCredId) > 0 {
				conf.Handler.Bcr.Region = region
			}
		}

		if err = conf.Write(); err != nil {
			logrus.Fatal(err)
		}
	},
}

func init() {
	bcrConfigCmd.Flags().StringP("name", "", "", "bcr name")
	bcrConfigCmd.Flags().StringP("host", "", "", "Bizfly host url")
	bcrConfigCmd.Flags().StringP("appCredId", "", "", "bizfly app credential id")
	bcrConfigCmd.Flags().StringP("appCredSecret", "", "", "bizfly app credential secret")
	bcrConfigCmd.Flags().StringP("region", "", "", "region")
}
