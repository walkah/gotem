/*
Copyright © 2024 James Walker <walkah@walkah.net>
*/
package cmd

import (
	"fmt"
	"os"

	git "github.com/go-git/go-git/v5"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"walkah.dev/walkah/gotem/internal/util"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Track the current path",
	Long:  `This command adds the current path to the list of git clones to track.`,
	Run: func(cmd *cobra.Command, args []string) {
		if cwd, err := os.Getwd(); err == nil {
			relativePath, err := util.GetRelativePath(cwd)
			if err != nil {
				panic(err)
			}

			repo, err := git.PlainOpen(cwd)
			if err != nil {
				panic(err)
			}
			remote, err := repo.Remote("origin")
			if err != nil {
				panic(err)
			}

			viper.Set(relativePath, map[string]string{"remote": remote.Config().URLs[0], "path": relativePath})
			viper.WriteConfig()
			fmt.Println(fmt.Sprintf("✅ Added %s!", relativePath))
		}
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
