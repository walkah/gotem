/*
Copyright © 2024 James Walker <walkah@walkah.net>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"walkah.dev/walkah/gotem/internal/git"
)

// upCmd represents the up command
var upCmd = &cobra.Command{
	Use:   "up",
	Short: "Update yer git",
	Long:  `Pulls the latest in all repos`,
	Run: func(cmd *cobra.Command, args []string) {
		for _, project := range viper.AllSettings() {
			path := project.(map[string]interface{})["path"].(string)
			err := git.PullLatest(path)
			if err != nil {
				fmt.Println(err)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(upCmd)
}
